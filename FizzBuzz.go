package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter a num: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSuffix(input, "\n'")
	strconv.ParseInt(input, 10, 32)
	fizzBuzz()
}

func fizzBuzz(input int) {
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

// func fizzBuzz(num int) {
// 	if num%5 == 0 && num%3 == 0 {
// 		fmt.Println("Fizz Buzz!")
// 	} else if num%5 == 0 {
// 		fmt.Println("Fizz")
// 	} else if num%3 == 0 {
// 		fmt.Println("Buzz")
// 	} else {
// 		fmt.Println("None")
// 	}
// }
