package main

import "fmt"

func main() {
	const c1 string = "hoge"
	const (
		c2 = 1
		c3 = c1
	)
	fmt.Println(c1, c2, c3)

	const (
		c4 = iota
		c5 = iota
		c6
	)
	const (
		c7 = iota
		c8 = 0
		c9 = iota
		c10
	)
	fmt.Println(c4, c5, c6)
	fmt.Println(c7, c8, c9, c10)
}
