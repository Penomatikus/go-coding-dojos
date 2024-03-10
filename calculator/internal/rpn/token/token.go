package token

const (
	Addition           = "+"
	Substaction        = "-"
	Multiplication     = "*"
	Division           = "/"
	OpeningParenthesis = "("
	ClosingParenthesis = ")"
)

func IsOperator(s string) bool {
	switch s {
	case Addition, Substaction, Division, Multiplication:
		return true
	default:
		return false
	}
}
