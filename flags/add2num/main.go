package main

import (
	"flag"
	"fmt"
)

func main() {
	a := flag.Int("a", 0, "first number")
	b := flag.Int("b", 0, "second number")
	flag.Parse()

	fmt.Printf("a + b = %d\n", *a+*b)
}
