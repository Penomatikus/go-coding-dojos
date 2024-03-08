package parser

import (
	"fmt"
	"strings"
)

var (
	// " " | "\n" | "\t" | "\r" | "\f" | "\b"
	allowedSeperators = []byte{0x20, 0x0A, 0x09, 0x0D, 0x0C, 0x08}
	// "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" | "."
	allowedNumericChars = []byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x2E}
	// "+" | "-"
	lowerPrecedenceOperators = []byte{0x2B, 0x2D}
	// "*" | "/"
	higherPrecedenceOperators = []byte{0x2A, 0x2F}
	// "+" | "-" | "*" | "/"
	allowedRpnOperators = append(lowerPrecedenceOperators, higherPrecedenceOperators...)
	// "("
	openingParenthesis = byte(0x28)
	// ")"
	closingParenthesis = byte(0x29)
	// "("  | ")"
	allowedNonRpnOperators = []byte{openingParenthesis, closingParenthesis}
)

type equationByte byte

func (eq equationByte) oneOf(bytes []byte) bool {
	for _, b := range bytes {
		if byte(eq) == b {
			return true
		}
	}
	return false
}

func (eq equationByte) is(b byte) bool {
	return byte(eq) == b
}

// toPostfix is reading equation one byte at a time
func parseEquationToPostfix(equation string) string {
	rpnStack := new(stack[string])
	operatorStack := new(Stack[equationByte])

	termBytes := make([]byte, 0)
	equationBytes := []equationByte(equation)
	for i, eb := range equationBytes {
		if eb.oneOf(allowedNumericChars) {
			termBytes = append(termBytes, byte(eb))
			continue
		}

		if eb.oneOf(append(allowedRpnOperators, allowedNonRpnOperators...)) {
			if equationBytes[i-1].oneOf(allowedNumericChars) {
				rpnStack.push(string(termBytes))
				termBytes = termBytes[:0]
			}
			process(rpnStack, operatorStack, eb)
			termBytes = termBytes[:0]
			continue
		}

		if eb.oneOf(allowedSeperators) {
			if len(termBytes) > 0 {
				rpnStack.push(string(termBytes))
				termBytes = termBytes[:0]
			}
			continue
		}

		if i != len(equationBytes) {
			fmt.Printf("Error at index %d: %c is not allowed", i, eb)
		}
	}

	if len(termBytes) > 0 {
		rpnStack.push(string(termBytes))
	}

	for i := operatorStack.len(); i >= 1; i-- {
		o, _ := operatorStack.pop()
		rpnStack.push(string(*o))
	}

	return strings.Join(rpnStack.items, " ")
}

func precedenceOf(operator equationByte) int {
	if operator.oneOf(higherPrecedenceOperators) {
		return 1
	}
	if operator.oneOf(lowerPrecedenceOperators) {
		return 0
	}

	panic(fmt.Sprintf("You are using it wrong. '%c' not part of the rpn allowed operators.", operator))
}

func process(rpnStack *Stack[string], operatorStack *Stack[equationByte], newOperator equationByte) {
	// Stack empty, we can just push it
	if operatorStack.empty() {
		operatorStack.push(newOperator)
		return
	}

	lastOperator, _ := operatorStack.peek()

	// Opening parentheses
	if newOperator.is(openingParenthesis) || lastOperator.is(openingParenthesis) {
		operatorStack.push(newOperator)
		return
	}

	// Closing parentheses; Pop all until first appearance of "("
	if newOperator.is(closingParenthesis) {
		for {
			o, _ := operatorStack.pop()
			if !o.is(openingParenthesis) {
				rpnStack.push(string(*o))
				continue
			}
			break
		}
		return
	}

	precedenceBefore := precedenceOf(*lastOperator)
	precedenceNext := precedenceOf(newOperator)

	// Rule 1
	if precedenceBefore < precedenceNext {
		operatorStack.push(newOperator)
		return
	}

	// Rule 2 || Rule 3
	if precedenceBefore > precedenceNext ||
		precedenceBefore == precedenceNext {
		o, _ := operatorStack.pop()
		rpnStack.push(string(*o))
		operatorStack.push(newOperator)
		return
	}
}
