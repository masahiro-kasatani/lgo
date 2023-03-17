package main

import "fmt"

func funcX(s string) {
	fmt.Print(s)
}

func funcY(s string) func() {
	fmt.Print(s)
	return func() { fmt.Print("Y") }
}

func main() {
	s := "A"
	defer funcX(s)
	s = "B"
	defer funcX(s)
	defer funcY(s)
	defer funcY(s)()
	defer func() {
		s = "C"
		fmt.Print(s)
	}()
}
