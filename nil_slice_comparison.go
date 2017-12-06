package main

import "fmt"

func main() {
	var x []string
	y := make([]string, 0)
	z := make([]string, 0, 3)
	fmt.Println(x == nil)
	fmt.Println(y == nil)
	fmt.Println(z == nil)
}
