package main

import (
	"fmt"
)

func boolVal() bool { return true }

func main() {
	x := true
	y := false
	z := boolVal()

	anded := x && y && z
	fmt.Print(anded)
	ored := x || y || z
	fmt.Print(ored)
	tst1 := !x || y && !z // !x || (y && !z)
	fmt.Print(tst1)
	tst2 := x && !(y || z)
	fmt.Print(tst2)
}
