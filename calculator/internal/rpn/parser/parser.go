package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Penomatikus/calculator/internal/rpn/stack"
)

var (
	ErrInvalidEquation = errors.New("invalid equation")
	ErrDivisionByZero  = errors.New("division by zero not allowed")

	allowedSeparators         = " \n\t\r\f\b"
	allowedNumericChars       = "0123456789."
	lowerPrecedenceOperators  = "+-"
	higherPrecedenceOperators = "*/"
	allowedRpnOperators       = lowerPrecedenceOperators + higherPrecedenceOperators
	openingParenthesis        = '('
	closingParenthesis        = ')'
)

// convenience type
type equationRune rune

func (eq equationRune) oneOf(chars string) bool {
	return strings.ContainsRune(chars, rune(eq))
}

func (eq equationRune) is(r rune) bool {
	return rune(eq) == r
}

func errorAtIndex(err error, i int, msg string) error {
	return fmt.Errorf("%w: Error at index %d: %s", err, i, msg)
}

// toPostfix is reading equation one byte at a time
func parseEquationToPostfix(equation string) (string, error) {
	rpnStack := stack.New[string]()
	operatorStack := stack.New[equationRune]()

	equationRunes := []equationRune(equation)
	if equationRunes[0].oneOf(fmt.Sprintf("%s%c", allowedRpnOperators, closingParenthesis)) {
		return "", errorAtIndex(ErrInvalidEquation, 0, fmt.Sprintf("Equation does not start with a number or '(': %c", equationRunes[0]))
	}

	term := make([]rune, 0)
	var err error
	for i, er := range equationRunes {
		if err = divZeroCheck(&equationRunes, er, i); err != nil {
			return "", err
		}

		switch {
		case er.oneOf(allowedNumericChars):
			term = append(term, rune(er))
		case er.is(openingParenthesis):
			operatorStack.Push(er)
		case er.is(closingParenthesis):
			pushTerm(&term, rpnStack)
			err = processClosingParenthesis(rpnStack, operatorStack)
		case er.oneOf(allowedRpnOperators):
			processRpnOperator(rpnStack, operatorStack, er)
		case er.oneOf(allowedSeparators):
			pushTerm(&term, rpnStack)
		default:
			return "", errorAtIndex(ErrInvalidEquation, i, fmt.Sprintf("'%c' is not an allowed character.", er))
		}

		if err != nil {
			return "", errorAtIndex(ErrInvalidEquation, i, err.Error())
		}
	}

	// push the remaining term and merge the operator stack to the rpn stack in LIFO
	pushTerm(&term, rpnStack)
	for i := operatorStack.Len(); i >= 1; i-- {
		o, _ := operatorStack.Pop()
		if o.is(openingParenthesis) {
			return "", errorAtIndex(ErrInvalidEquation, -1, "missing closing parenthesis")
		}
		rpnStack.Push(string(*o))
	}

	return strings.Join(rpnStack.PeekAll(), " "), nil
}

// pushTerm pushes term to rpnStack and resets term.
func pushTerm(term *[]rune, rpnStack *stack.Stack[string]) {
	if len(*term) > 0 {
		rpnStack.Push(string(*term))
		*term = (*term)[:0]
	}
}

// divZeroCheck returns an error if current is 0 and a divisor
func divZeroCheck(equationRunes *[]equationRune, current equationRune, currentIndex int) error {
	if !current.is('0') {
		return nil
	}

	lookupIndex := currentIndex - 2
	if lookupIndex > 0 && (*equationRunes)[lookupIndex].is('/') {
		return errorAtIndex(ErrDivisionByZero, currentIndex, "Divison by 0 not allowed")
	}

	return nil
}

// processClosingParenthesis pushes all poped items from operatorStack to rpnStack until "(".
// Return an error if, the operator stack has no operators left but no opening parenthesis was found.
func processClosingParenthesis(rpnStack *stack.Stack[string], operatorStack *stack.Stack[equationRune]) error {
	for {
		o, ok := operatorStack.Pop()
		if !ok {
			return errors.New("missing opening parenthesis")
		}
		if !o.is(openingParenthesis) {
			rpnStack.Push(string(*o))
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
func processRpnOperator(rpnStack *stack.Stack[string], operatorStack *stack.Stack[equationRune], newOperator equationRune) {
	lastOperator, _ := operatorStack.Peek()

	// No precedence check if stack is empty or top stack operator is an opening parenthesis
	if operatorStack.Empty() || lastOperator.is(openingParenthesis) {
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

	// Rule 2 || Rule 3
	if precedenceBefore > precedenceNext ||
		precedenceBefore == precedenceNext {
		o, _ := operatorStack.Pop()
		rpnStack.Push(string(*o))
		operatorStack.Push(newOperator)
		return
	}
}
