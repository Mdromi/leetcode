package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Enter the value of a: ")
	fmt.Scan(&a)

	fmt.Println("Enter the value of b: ")
	fmt.Scan(&b)
	fmt.Printf("%d\n", a+b)
}
