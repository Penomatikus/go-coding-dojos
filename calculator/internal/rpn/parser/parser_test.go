package parser

import (
	"errors"
	"testing"
)

func Test_parseEquationToPostfix(t *testing.T) {
	type args struct {
		name, infix, postfix string
	}

	tests := []args{
		{name: "Rule 1", infix: "4 * (2 + 9 / 8)", postfix: "4 2 9 8 / + *"},
		{name: "Rule 2", infix: "2 * 3 + 12 / 4", postfix: "2 3 * 12 4 / +"},
		{name: "Rule 3", infix: "1 * 2 / 3", postfix: "1 2 * 3 /"},
		{name: "With .", infix: "3.412 / 4 * (9 + 3) + 3 - (1 - 4)", postfix: "3.412 4 / 9 3 + * 3 + 1 4 - -"},
		{name: "With nested parenthesis", infix: "3 / 4 * ((9 + 3) * (8 - 9)) + 3 - (1 - 4)", postfix: "3 4 / 9 3 + 8 9 - * * 3 + 1 4 - -"},
	}

	for _, tt := range tests {
		if got, _ := parseEquationToPostfix(tt.infix); got != tt.postfix {
			t.Fatalf("Infix: %s; Expected: %s Got: %s ", tt.infix, tt.postfix, got)
		}
	}
}

func Test_parseEquationToPostfix_Bad_Equation(t *testing.T) {
	type args struct {
		name, infix string
		err         error
	}

	tests := []args{
		{name: "Not allowed rune", infix: "4 * (2 + a / 8)", err: ErrInvalidEquation},
		{name: "Not allowed start", infix: "* (2 + a / 8)", err: ErrInvalidEquation},
		{name: "Not allowed start 2", infix: ") + (2 + a / 8)", err: ErrInvalidEquation},
		{name: "Division by zero", infix: "4 * (2 + 9 / 0)", err: ErrDivisionByZero},
	}

	for _, tt := range tests {
		got, err := parseEquationToPostfix(tt.infix)
		if got != "" {
			t.Fatal("Got a postfix notation, expected <empty>")
		}
		if err == nil {
			t.Fatal("Got no error on an empty postfix notation")

		}
		if !errors.Is(err, tt.err) {
			t.Fatal("Bad error")
		}
	}
}
