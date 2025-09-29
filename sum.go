package main

import "fmt"

func main() {
	fmt.Println("O valor total da soma é:", sum(2, 2))
}

func sum(a, b int) int {
	return a + b
}
