package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	fmt.Println(x)
	y := x[:2]                          // 1,2
	z := x[2:]                          // 3,4
	fmt.Println(cap(x), cap(y), cap(z)) // cap(y)=cap(x)-0, cap(z)=cap(x)-2
	yy := make([]int, 2)
	copy(yy, x[1:])
	fmt.Println(yy)
}
