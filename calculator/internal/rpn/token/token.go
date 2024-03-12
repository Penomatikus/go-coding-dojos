// package token holds the allowed four basic arithmetic oprantors, a function to check if a string is one of that operatos and constans of opening and closing parenthesis.
package token

const (
	Addition           = '+'
	Substaction        = '-'
	Multiplication     = '*'
	Division           = '/'
	OpeningParenthesis = '('
	ClosingParenthesis = ')'
)

func IsOperator(s string) bool {
	if len(s) > 1 {
		return false
	}
	switch s[0] {
	case Addition, Substaction, Division, Multiplication:
		return true
	default:
		return false
	}
}
