package main

import "fmt"

func checkType(a interface{}) {
	switch T := a.(type) {
	case int32:
		fmt.Println("int", T)
	case string:
		fmt.Println("string", T)
	default:
		fmt.Println("unknown", T)
	}
}

func main() {
	var a interface{} = "hello"
	checkType(a)
}
