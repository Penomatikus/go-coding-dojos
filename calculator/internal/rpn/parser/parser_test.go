package parser

import (
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
		if got := parseEquationToPostfix(tt.infix); got != tt.postfix {
			t.Fatalf("Infix: %s; Expected: %s Got: %s ", tt.infix, tt.postfix, got)
		}
	}
}

func Test_parseEquationToPostfix_Bad_Equation(t *testing.T) {
	type args struct {
		name, infix, postfix string
	}

	tests := []args{
		{name: "Rule 1", infix: "4 * (2 + a / 8)", postfix: "4 2 9 8 / + *"},
	}

	for _, tt := range tests {
		if got := parseEquationToPostfix(tt.infix); got != tt.postfix {
			t.Fatalf("Infix: %s; Expected: %s Got: %s ", tt.infix, tt.postfix, got)
		}
	}
}
