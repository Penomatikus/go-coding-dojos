package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("%s NaN", os.Args[1]))
	}

	spiral := make([][]int, length)
	for i := range spiral {
		spiral[i] = make([]int, length)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i == 0 {
				spiral[i][j] = 1
			}
			if j == length-1 {
				spiral[i][j] = 1
			}
			if i == length-1 {
				spiral[i][j] = 1
			}
			if j == 0 && i > 1 {
				spiral[i][j] = 1
			}
			// if j == length%2 {
			// 	spiral[i][j] = 2
			// }
			// if j == length-i {
			// 	spiral[i][j] = 2
			// }
		}
		fmt.Println(spiral[i])
	}

}
