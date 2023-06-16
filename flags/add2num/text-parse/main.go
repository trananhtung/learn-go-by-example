package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var text, convertType string
	flag.StringVar(&convertType, "type", "uppercase", "type of input")
	flag.StringVar(&text, "t", "", "text to parse")

	flag.Parse()

	switch convertType {
	case "uppercase":
		// convert text to uppercase
		uppercaseString := strings.ToUpper(text)
		fmt.Println(uppercaseString)

	case "lowercase":
		// convert text to lowercase
		lowercaseString := strings.ToLower(text)
		fmt.Println(lowercaseString)
	case "title":
		// convert text to title case
		titleString := strings.ToTitle(text)
		fmt.Println(titleString)
	}
}
