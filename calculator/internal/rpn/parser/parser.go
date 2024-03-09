package parser

import (
	"fmt"
	"strings"

	"github.com/Penomatikus/calculator/internal/rpn/stack"
)

var (
	allowedSeperators         = []rune{' ', '\n', '\t', '\r', '\f', '\b'}
	allowedNumericChars       = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}
	lowerPrecedenceOperators  = []rune{'+', '-'}
	higherPrecedenceOperators = []rune{'*', '/'}
	allowedRpnOperators       = append(lowerPrecedenceOperators, higherPrecedenceOperators...)
	openingParenthesis        = '('
	closingParenthesis        = ')'
)

type equationRune rune

func (eq equationRune) oneOf(runes []rune) bool {
	for _, r := range runes {
		if rune(eq) == r {
			return true
		}
	}
	return false
}

func (eq equationRune) is(r rune) bool {
	return rune(eq) == r
}

// toPostfix is reading equation one byte at a time
func parseEquationToPostfix(equation string) string {
	rpnStack := stack.New[string]()
	operatorStack := stack.New[equationRune]()

	term := make([]rune, 0)
	equationRunes := []equationRune(equation)
	for i, er := range equationRunes {
		if er.oneOf(allowedNumericChars) {
			term = append(term, rune(er))
			continue
		}

		if er.is(openingParenthesis) {
			operatorStack.Push(er)
			continue
		}

		if er.is(closingParenthesis) {
			pushTerm(&term, rpnStack)
			processClosingParenthesis(rpnStack, operatorStack)
			continue
		}

		if er.oneOf(allowedRpnOperators) {
			processRpnOperator(rpnStack, operatorStack, er)
			continue
		}

		if er.oneOf(allowedSeperators) {
			pushTerm(&term, rpnStack)
			continue
		}

		if i != len(equationRunes) {
			fmt.Printf("Error at index %d: '%c' is not allowed", i, er)
		}
	}

	pushTerm(&term, rpnStack)

	for i := operatorStack.Len(); i >= 1; i-- {
		o, _ := operatorStack.Pop()
		rpnStack.Push(string(*o))
	}

	return strings.Join(rpnStack.PeekAll(), " ")
}

func pushTerm(term *[]rune, rpnStack *stack.Stack[string]) {
	if len(*term) > 0 {
		rpnStack.Push(string(*term))
		*term = (*term)[:0]
	}
}

func processClosingParenthesis(rpnStack *stack.Stack[string], operatorStack *stack.Stack[equationRune]) {
	for {
		o, ok := operatorStack.Pop()
		if ok && !o.is(openingParenthesis) {
			rpnStack.Push(string(*o))
			continue
		}
		break
	}
}

func precedenceOf(operator equationRune) int {
	if operator.oneOf(higherPrecedenceOperators) {
		return 1
	}
	if operator.oneOf(lowerPrecedenceOperators) {
		return 0
	}

	panic(fmt.Sprintf("You are using it wrong. '%c' not part of the rpn allowed operators.", operator))
}

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
