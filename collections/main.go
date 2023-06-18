package main

import "fmt"

type Item interface{}

func Filter(s []Item, f func(Item) bool) Item {
	var p []Item // == nil
	for _, v := range s {
		if f(v) {
			p = append(p, v)
		}
	}
	return p
}

func Map(s []Item, f func(Item) interface{}) []interface{} {
	var p = make([]interface{}, len(s))
	for i, v := range s {
		p[i] = f(v)
	}
	return p
}

func Some(s []Item, f func(Item) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

func Every(s []Item, f func(Item) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func Find(s []Item, f func(Item) bool) Item {
	for _, v := range s {
		if f(v) {
			return v
		}
	}
	return nil
}

func main() {
	s := []Item{1, 2, 3, 4, 5}
	p := Filter(s, func(v Item) bool {
		return v.(int) > 3
	})
	fmt.Println(p)

	p2 := Map(s, func(v Item) interface{} {
		return v.(int) * 2
	})
	fmt.Println(p2)

	p3 := Some(s, func(v Item) bool {
		return v.(int) > 3
	})
	fmt.Println(p3)

	p4 := Every(s, func(v Item) bool {
		return v.(int) > 3
	})
	fmt.Println(p4)
}
