package main

import (
	"fmt"
	"reflect"
)

type SexType int

const (
	Male SexType = iota
	Female
)

type Person struct {
	Name string
	Age  int
	Sex  SexType
}

func main() {
	tung := Person{
		Name: "Tung",
		Age:  18,
		Sex:  Male,
	}

	fmt.Println(reflect.TypeOf(tung))
	fmt.Println(reflect.ValueOf(tung))
	fmt.Println(reflect.ValueOf(tung).Kind())

	r := reflect.ValueOf(tung)

	for i := 0; i < r.NumField(); i++ {
		fmt.Println(r.Field(i))
	}
}
