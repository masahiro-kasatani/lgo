package main

import "fmt"

func funcX() {
	fmt.Print("A")
}

func funcY(s string) {
	fmt.Print(s)
}

func funcZ(s string) func() {
	fmt.Print(s)
	return func() { fmt.Print("B") }
}

func main() {
	defer funcX()
	s := "C"
	defer funcY(s)
	s = "D"
	defer funcZ(s)
	defer funcZ(s)()
}
