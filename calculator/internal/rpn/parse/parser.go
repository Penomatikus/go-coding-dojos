package parse

import (
	"errors"
	"fmt"

	"github.com/Penomatikus/calculator/internal/rpn/stack"
	"github.com/Penomatikus/calculator/internal/rpn/token"
)

var (
	ErrInvalidEquation = errors.New("invalid equation")
	ErrDivisionByZero  = errors.New("division by zero not allowed")

	allowedSeparators         = []rune{' ', '\n', '\t', '\r', '\f', '\b'}
	allowedNumericChars       = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}
	lowerPrecedenceOperators  = []rune{token.Addition, token.Substaction}
	higherPrecedenceOperators = []rune{token.Multiplication, token.Division}
	allowedRpnOperators       = append(lowerPrecedenceOperators, higherPrecedenceOperators...)
)

// convenience type
type equationRune rune

func (eq equationRune) oneOf(chars []rune) bool {
	for _, c := range chars {
		if rune(eq) == c {
			return true
		}
	}
	return false
}

func (eq equationRune) is(r rune) bool {
	return rune(eq) == r
}

func errorAtIndex(err error, i int, msg string) error {
	return fmt.Errorf("%w: Error at index %d: %s", err, i, msg)
}

// toPostfix is reading equation one byte at a time
func ToPostfix(equation string) ([]string, error) {
	notation := make([]string, 0)
	operatorStack := stack.New[equationRune]()

	equationRunes := []equationRune(equation)
	if equationRunes[0].oneOf(append(allowedRpnOperators, token.ClosingParenthesis)) {
		return nil, errorAtIndex(ErrInvalidEquation, 0, fmt.Sprintf("Equation does not start with a number or '(': %c", equationRunes[0]))
	}

	operand := make([]rune, 0)
	var err error
	for i, er := range equationRunes {
		if err = divZeroCheck(equationRunes, er, i); err != nil {
			return nil, err
		}

		switch {
		case er.oneOf(allowedNumericChars):
			operand = append(operand, rune(er))

		case er.is(token.OpeningParenthesis):
			operatorStack.Push(er)

		case er.is(token.ClosingParenthesis):
			pushOperand(&operand, &notation)
			err = processClosingParenthesis(&notation, operatorStack)

		case er.oneOf(allowedRpnOperators):
			processRpnOperator(&notation, operatorStack, er)

		case er.oneOf(allowedSeparators):
			pushOperand(&operand, &notation)

		default:
			return nil, errorAtIndex(ErrInvalidEquation, i, fmt.Sprintf("'%c' is not an allowed character.", er))
		}

		if err != nil {
			return nil, errorAtIndex(ErrInvalidEquation, i, err.Error())
		}
	}

	// push the remaining operand and merge the operator stack to the rpn stack in LIFO
	pushOperand(&operand, &notation)
	for i := operatorStack.Len(); i >= 1; i-- {
		o := operatorStack.Pop()
		if o.is(token.OpeningParenthesis) {
			return nil, errorAtIndex(ErrInvalidEquation, -1, "missing closing parenthesis")
		}
		notation = append(notation, string(o))
	}

	return notation, nil
}

// pushOperand pushes operand to rpnStack and resets operand.
func pushOperand(operand *[]rune, notation *[]string) {
	if len(*operand) > 0 {
		*notation = append(*notation, string(*operand))
		*operand = (*operand)[:0]
	}
}

// divZeroCheck returns an error if current is 0 and a divisor
func divZeroCheck(equationRunes []equationRune, current equationRune, currentIndex int) error {
	if !current.is('0') {
		return nil
	}

	lookupIndex := currentIndex - 2
	if lookupIndex > 0 && equationRunes[lookupIndex].is(token.Division) {
		return errorAtIndex(ErrDivisionByZero, currentIndex, "Divison by 0 not allowed")
	}

	return nil
}

// processClosingParenthesis pushes all poped items from operatorStack to rpnStack until "(".
// Return an error if, the operator stack has no operators left but no opening parenthesis was found.
func processClosingParenthesis(notation *[]string, operatorStack *stack.Stack[equationRune]) error {
	for !operatorStack.Empty() {
		o := operatorStack.Pop()
		if !o.is(token.OpeningParenthesis) {
			*notation = append(*notation, string(o))
			continue
		}
		return nil
	}
	return errors.New("missing opening parenthesis")
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
	lastOperator, ok := operatorStack.Peek()

	// No precedence check if stack is empty or top stack operator is an opening parenthesis
	if !ok || lastOperator.is(token.OpeningParenthesis) {
		operatorStack.Push(newOperator)
		return
	}

	precedenceBefore := precedenceOf(lastOperator)
	precedenceNext := precedenceOf(newOperator)

	// Rule 1
	if precedenceBefore < precedenceNext {
		operatorStack.Push(newOperator)
		return
	}

	// Rule 2
	if precedenceBefore == precedenceNext {
		*notation = append(*notation, string(operatorStack.Pop()))
		operatorStack.Push(newOperator)
	}

	// Rule 3
	if precedenceBefore > precedenceNext {
		for i := operatorStack.Len(); i >= 1; i-- {
			*notation = append(*notation, string(operatorStack.Pop()))
		}
		operatorStack.Push(newOperator)
		return
	}
}
