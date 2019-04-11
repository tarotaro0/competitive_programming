package main

import "fmt"

func main() {
	var a, b, c, d, e, k int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	fmt.Scan(&e)
	fmt.Scan(&k)

	if k >= (e - a) {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}
