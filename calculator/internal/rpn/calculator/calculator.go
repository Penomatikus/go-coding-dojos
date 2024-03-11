package calculator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Penomatikus/calculator/internal/rpn/parse"
	"github.com/Penomatikus/calculator/internal/rpn/stack"
	tkn "github.com/Penomatikus/calculator/internal/rpn/token"
)

func Calculate(equation string) (float64, error) {
	notation, err := parse.ToPostfix(equation)
	if err != nil {
		return 0, err
	}

	numStack := stack.New[float64]()
	for _, token := range notation {
		if num, ok := isNumber(token); ok {
			numStack.Push(num)
			continue
		}

		if !tkn.IsOperator(token) {
			panic(fmt.Sprintf("parser is not producing corrrect notation. %s is not an operator", token))
		}

		left := numStack.Pop()
		rigth := numStack.Pop()
		switch token {
		case string(tkn.Addition):
			numStack.Push(rigth + left)
		case string(tkn.Substaction):
			numStack.Push(rigth - left)
		case string(tkn.Division):
			numStack.Push(rigth / left)
		case string(tkn.Multiplication):
			numStack.Push(rigth * left)
		}
	}

	return numStack.Pop(), nil
}

func RPN(equation string) (string, error) {
	notation, err := parse.ToPostfix(equation)
	if err != nil {
		return "", err
	}
	return strings.Join(notation, " "), nil
}

func isNumber(s string) (float64, bool) {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f, true
	}
	return -1, false
}
