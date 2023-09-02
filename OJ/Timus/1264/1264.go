package main

import "fmt"

func main() {
	var operand1, operand2 int
	fmt.Scan(&operand1, &operand2)
	operand2 += 1
	fmt.Printf("%d\n", operand1*operand2)
}
