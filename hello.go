package main

import "fmt"

var total int

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(mult(1, 2))
	fmt.Println(divide(1, 2))

}

func add(x int, y int) int {

	return x + y

}

func sub(x int, y int) int {

	return x - y
}

func mult(x, y int) int {

	return x * y

}

func divide(x, y int) float32 {
	return float32(x) / float32(y)

}
