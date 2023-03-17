package main

import "fmt"

func funcX() int {
	x := 0
	defer func() { x = 1 }()
	return x
}

func funcY() (y int) {
	y = 0
	defer func() { y = 1 }()
	return y
}

func main() {
	fmt.Println("x:", funcX())
	fmt.Println("y:", funcY())
}
