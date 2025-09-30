package main

import "fmt"

func main() {
	fmt.Println("O valor total da soma é:", sum(2, 2))
	// fmt.Println("O valor total da subtração é:", sub(2, 2))
	// fmt.Println("O valor total da multiplicação é:", times(2, 2))
	// fmt.Println("O valor total da soma com dobro é:", sumX(2, 2))
}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func times(a, b int) int {
	return a * b
}

func sumX(a, b int) int {
	return a + b + b
}
