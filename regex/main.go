package main

import (
	"log"
	"os"
	"regexp"
)

const (
	EMAIL_REGEX = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

func main() {
	// get args
	args := os.Args[1:]

	// check args
	if len(args) < 1 {
		log.Fatal("Please provide an email address")
	}

	// check email
	email := args[0]
	re := regexp.MustCompile(EMAIL_REGEX)

	if !re.MatchString(email) {
		log.Fatal("Invalid email address")
	}
	log.Println("Email address is valid")
}
