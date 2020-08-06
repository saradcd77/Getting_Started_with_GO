package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a num: ")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	fizzBuzz(input)
}

func fizzBuzz(input int64) {
	if input%5 == 0 && input%3 == 0 {
		fmt.Println("Fizz Buzz!")
	} else if input%5 == 0 {
		fmt.Println("Fizz")
	} else if input%3 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("None")
	}
}
