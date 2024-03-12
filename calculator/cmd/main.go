package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Penomatikus/calculator/internal/rpn/calculator"
)

var (
	rpn  = flag.Bool("n", false, "show the reserve polish notation of the equation")
	calc = flag.Bool("c", true, "calculate the equation")
	file = flag.String("f", "", "read equation from file")
)

func main() {
	flag.Parse()

	var equation string
	if len(*file) != 0 {
		data, err := os.ReadFile(*file)
		handleError(err)
		equation = string(data)
	} else {
		fmt.Scanln(&equation)
	}

	if *calc {
		cs, err := calculator.Calculate(equation)
		handleError(err)
		fmt.Printf("Solution: \n%f\n", cs)
	}

	if *rpn {
		rr, err := calculator.RPN(equation)
		handleError(err)
		fmt.Printf("RPN: \n%s\n", rr)
	}

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
