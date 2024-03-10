package parse

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Penomatikus/calculator/internal/rpn/stack"
	"github.com/Penomatikus/calculator/internal/rpn/token"
)

var (
	ErrInvalidEquation = errors.New("invalid equation")
	ErrDivisionByZero  = errors.New("division by zero not allowed")

	allowedSeparators         = " \n\t\r\f\b"
	allowedNumericChars       = "0123456789."
	lowerPrecedenceOperators  = string(token.Addition + token.Substaction)
	higherPrecedenceOperators = string(token.Multiplication + token.Division)
	allowedRpnOperators       = string(lowerPrecedenceOperators + higherPrecedenceOperators)
)

// convenience type
type equationRune rune

func (eq equationRune) oneOf(chars string) bool {
	return strings.ContainsRune(chars, rune(eq))
}

func (eq equationRune) is(s string) bool {
	return string(eq) == s
}

func errorAtIndex(err error, i int, msg string) error {
	return fmt.Errorf("%w: Error at index %d: %s", err, i, msg)
}

// toPostfix is reading equation one byte at a time
func ToPostfix(equation string) ([]string, error) {
	notation := make([]string, 0)
	operatorStack := stack.New[equationRune]()

	equationRunes := []equationRune(equation)
	if equationRunes[0].oneOf(allowedRpnOperators + string(token.ClosingParenthesis)) {
		return nil, errorAtIndex(ErrInvalidEquation, 0, fmt.Sprintf("Equation does not start with a number or '(': %c", equationRunes[0]))
	}

	term := make([]rune, 0)
	var err error
	for i, er := range equationRunes {
		if err = divZeroCheck(&equationRunes, er, i); err != nil {
			return nil, err
		}

		switch {
		case er.oneOf(allowedNumericChars):
			term = append(term, rune(er))

		case er.is(token.OpeningParenthesis):
			operatorStack.Push(er)

		case er.is(token.ClosingParenthesis):
			pushTerm(&term, &notation)
			err = processClosingParenthesis(&notation, operatorStack)

		case er.oneOf(allowedRpnOperators):
			processRpnOperator(&notation, operatorStack, er)

		case er.oneOf(allowedSeparators):
			pushTerm(&term, &notation)

		default:
			return nil, errorAtIndex(ErrInvalidEquation, i, fmt.Sprintf("'%c' is not an allowed character.", er))
		}

		if err != nil {
			return nil, errorAtIndex(ErrInvalidEquation, i, err.Error())
		}
	}

	// push the remaining term and merge the operator stack to the rpn stack in LIFO
	pushTerm(&term, &notation)
	for i := operatorStack.Len(); i >= 1; i-- {
		o, _ := operatorStack.Pop()
		if o.is(token.OpeningParenthesis) {
			return nil, errorAtIndex(ErrInvalidEquation, -1, "missing closing parenthesis")
		}
		notation = append(notation, string(*o))
	}

	return notation, nil
}

// pushTerm pushes term to rpnStack and resets term.
func pushTerm(term *[]rune, notation *[]string) {
	if len(*term) > 0 {
		*notation = append(*notation, string(*term))
		*term = (*term)[:0]
	}
}

// divZeroCheck returns an error if current is 0 and a divisor
func divZeroCheck(equationRunes *[]equationRune, current equationRune, currentIndex int) error {
	if !current.is("0") {
		return nil
	}

	lookupIndex := currentIndex - 2
	if lookupIndex > 0 && (*equationRunes)[lookupIndex].is(token.Division) {
		return errorAtIndex(ErrDivisionByZero, currentIndex, "Divison by 0 not allowed")
	}

	return nil
}

// processClosingParenthesis pushes all poped items from operatorStack to rpnStack until "(".
// Return an error if, the operator stack has no operators left but no opening parenthesis was found.
func processClosingParenthesis(notation *[]string, operatorStack *stack.Stack[equationRune]) error {
	for {
		o, ok := operatorStack.Pop()
		if !ok {
			return errors.New("missing opening parenthesis")
		}
		if !o.is(token.OpeningParenthesis) {
			*notation = append(*notation, string(*o))
			continue
		}
		break
	}
	return nil
}

// precedenceOf returns the precedence of operator
//   - +, -: 0
//   - *, /: 1
func precedenceOf(operator equationRune) int {
	if operator.oneOf(higherPrecedenceOperators) {
		return 1
	}
	if operator.oneOf(lowerPrecedenceOperators) {
		return 0
	}

	panic(fmt.Sprintf("You are using it wrong. '%c' not part of the rpn allowed operators.", operator))
}

// processRpnOperator
func processRpnOperator(notation *[]string, operatorStack *stack.Stack[equationRune], newOperator equationRune) {
	lastOperator, _ := operatorStack.Peek()

	// No precedence check if stack is empty or top stack operator is an opening parenthesis
	if operatorStack.Empty() || lastOperator.is(token.OpeningParenthesis) {
		operatorStack.Push(newOperator)
		return
	}

	precedenceBefore := precedenceOf(*lastOperator)
	precedenceNext := precedenceOf(newOperator)

	// Rule 1
	if precedenceBefore < precedenceNext {
		operatorStack.Push(newOperator)
		return
	}

	// Rule 2
	if precedenceBefore == precedenceNext {
		o, _ := operatorStack.Pop()
		*notation = append(*notation, string(*o))
		operatorStack.Push(newOperator)
	}

	// Rule 3
	if precedenceBefore > precedenceNext {
		for i := operatorStack.Len(); i >= 1; i-- {
			o, _ := operatorStack.Pop()
			*notation = append(*notation, string(*o))
		}
		operatorStack.Push(newOperator)
		return
	}
}
