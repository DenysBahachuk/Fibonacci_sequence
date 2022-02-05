package main

import (
	"fmt"
	"os"
	"strconv"
)

func fibNumber(n int) int {

	if n < 2 {
		return n
	}

	return fibNumber(n-1) + fibNumber(n-2)
}

func main() {

	//Enter the number of fibonacci sequence numbers you want to display
	input := os.Args
	if len(input) != 2 {
		fmt.Println("Enter one number!")
		return
	}

	number, err := strconv.Atoi(input[1])
	if err != nil {
		fmt.Println("Enter a number!\nError:", err)
	} else if number <= 0 {
		fmt.Println("Enter one positive integer number!")
	}

	for i := 0; i < number; i++ {
		fmt.Printf("%d ", fibNumber(i))
	}

}
