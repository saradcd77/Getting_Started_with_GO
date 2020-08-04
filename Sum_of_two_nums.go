package main

import "fmt"

func main() {
	var result int
	result = add(5, 10)
	fmt.Println("The result is:", result)

}

func add(a int, b int) int {
	var c int
	c = a + b
	return c
}
