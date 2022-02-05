package main

import (
	"fmt"
	"os"
	"strconv"
)

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

	var prev int64 = 1
	var curr int64 = 0
	var temp int64 = 0

	for i := 0; i < number; i++ {
		fmt.Printf("%d ", curr)
		temp = prev
		prev = curr
		curr = temp + curr
	}
}
