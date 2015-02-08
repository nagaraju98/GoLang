package main

import "fmt"

func swap(x, y string) (string, int) {
	return x + "-" + y, 10
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}