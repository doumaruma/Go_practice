package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}
func hello() {
	fmt.Println("Hello World")
}
func add2(x int) func(int) int {
	return func (y int) int {
		return y + x
	}
}	//関数ファクトリの実装

func closure1() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main(){
	x, y := 1, 3
	fmt.Printf("%d + %d = %d\n", x, y, add(x, y))
	hello()
	func1 :=  add2(5)
	fmt.Printf("%d\n", func1(3))

	clo := closure1()
	fmt.Printf("%d\n", clo())
	fmt.Printf("%d\n", clo())
	fmt.Printf("%d\n", clo())
	fmt.Printf("%d\n", clo())
}