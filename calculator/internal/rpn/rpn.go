package rpn

type term float32

type operator string

const (
	add operator = "+"
	sub operator = "-"
	div operator = "/"
	mul operator = "*"
)

var (
	// " " | "\n" | "\t" | "\r" | "\f" | "\b"
	allowedSeperators = []byte{0x20, 0x0A, 0x09, 0x0D, 0x0C, 0x08}
	// "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" | "."
	allowedNumericChars = []byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x2E}
	// "+" | "-" | "*" | "/" | "(" | ")"
	allowedNonRpnOperators = []byte{0x2B, 0x2D, 0x2A, 0x2F, 0x28, 0x29}
)

type equationByte byte

func (eq equationByte) oneOf(bytes []byte) (ok bool) {
	for _, b := range bytes {
		if byte(eq) == b {
			ok = true
			break
		}
	}
	return
}

// toPostfix is reading equation one byte at a time
func toPostfix(equation string) string {
	rpnStack := new(stack[string])
	// operatorStack := new(stack[string])

	equationBytes := []equationByte(equation)
	termOrOperatorBytes := new(stack[equationByte])
	// operatorStack := new(stack[equationByte])

	for i, eb := range equationBytes {
		if eb.oneOf(allowedNumericChars) || eb.oneOf(allowedNonRpnOperators) {
			termOrOperatorBytes.push(eb)
		}

		if eb.oneOf(allowedSeperators) {
			e, l := termOrOperatorBytes.popAll()

		}

		if eb.oneOf(allowedNonRpnOperators) {
			operatorStack.push(eb)
		}
	}
	// return ""
}

// toPostfix is reading equation one byte at a time  
// func toPostfix(equation string) string {
// 	rpnStack := new(stack[string])
// 	operatorStack := new(stack[string])

// 	equationBytes := []equationByte(equation)
// 	termOrOperatorBytes := new(stack[equationByte])
// 	for i := 0; i < len(equationBytes); i++ {
// 		eb := equationBytes[i]
// 		if eb.oneOf(allowedNumericChars) {

// 		}

// 		if equationBytes[i].oneOf(allowedSeperators) {
// 			// termOrOperatorBytes
// 		}
// 		if equationBytes[i].oneOf(allowedNonRpnOperators) {
// 			operatorStack.push()
// 		}
// 	}
// 	// return ""
// }

// 3 ÷ 4 • (9 + 3) + 3 - (1 - 4)

// [3, 4, ÷, 9, 3, + •, 3, +, 1, 4, -, - ]
// [-, (, -, ) ]

// 3 ÷ 4 • ((9 + 3) • (8 - 9)) + 3 - (1 - 4)

// [3, 4, ÷, 9, 3, +, 8, 9, -, •, •, 3, +, 1, 4, -, -]
// [-]

/* Regeln:
A) 1. Operator > 2. Operator: Den 1. in den Stack verschieben
B) 1. Operator < 2. Operator: Wir machen nichts
C) 1. Operator == 2. Operator: Den 1. in den Stack verschieben
D) Klammern sind Operatoren
	D.1) Es werden alle Operatoren, die nach ABC in einer KLammer gesammelt wurden, in den Stack verschoben)
*/

type stack[T any] struct {
	items []T
}

func (s *stack[T]) len() int {
	return len(s.items)
}

func (s *stack[T]) pop() (*T, bool) {
	if s.len() < 1 {
		return nil, false
	}

	i := s.items[s.len()-1]
	s.items = s.items[:s.len()-1]
	return &i, true
}

func (s *stack[T]) push(item T) {
	s.items = append(s.items, item)
}

// popAll returns the complete stack "as is" and its lenght and cleears s.
func (s *stack[T]) popAll() ([]T, int) {
	all := s.items
	s.items = s.items[:0]
	return all, len(all)
}
