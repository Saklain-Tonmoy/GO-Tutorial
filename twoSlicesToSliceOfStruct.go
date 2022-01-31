package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	type MyCombo struct {
		number int
		boolean bool
	}

	a := []MyCombo{}

	for i := 0; i < len(q); i++ {
		a = append(a, MyCombo{q[i], r[i]})
	}

	fmt.Println(a)

}