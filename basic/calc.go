package main

import "fmt"

var Pv int

func init() {
	fmt.Println("init() of calc")
}

func Pow2() int {
	return Pv * Pv
}

func AddSub(x, y int) (int, int) {
	return x + y, x - y
}
