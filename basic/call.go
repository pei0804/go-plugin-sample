package main

import (
	"fmt"
	"os"
	"plugin"
)

func main() {
	p, err := plugin.Open("calc.so")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(-1)
	}

	pv, err := p.Lookup("Pv")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(-1)
	}

	pow2, err := p.Lookup("Pow2")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(-1)
	}

	*pv.(*int) = 3
	p2 := pow2.(func() int)()
	fmt.Printf("pow2: %d * %d = %d\n", *pv.(*int), *pv.(*int), p2)

	addSub, err := p.Lookup("AddSub")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(-1)
	}

	x, y := 5, 2
	add, sub := addSub.(func(int, int) (int, int))(x, y)
	fmt.Printf("add: %d + %d = %d\n", x, y, add)
	fmt.Printf("sub: %d - %d = %d\n", x, y, sub)
}
