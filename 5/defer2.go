package main

import (
	"fmt"
	"log"
)

func funcX() {
	defer func() {
		err := recover()
		fmt.Println("funcX recover:", err)
	}()
	log.Panic("panic")
}

func funcY() {
	defer func() {
		err := recover()
		fmt.Println("funcY recover:", err)
	}()
	log.Fatal("fatal")
}

func main() {
	funcX()
	funcY()
}
