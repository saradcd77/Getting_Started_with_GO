package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter num: ")
	text = reader.ReadString('\n')
	var result boolean
	result = fizzBuzz(5)
}

func fizzBuzz(num int) boolean {
	if num%5 == 0 && num%3 == 0 {
		fmt.Println("Fizz Buzz!")
	} else if num%5 == 0 {
		fmt.Println("Fizz")
	} else if num%3 == 0 {
		fmt.Println("Buzz!")
	} else {
		fmt.Println("None")
	}
}
