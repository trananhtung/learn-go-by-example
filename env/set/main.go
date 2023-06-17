package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("PORT", "8080")
	os.Setenv("ENV", "development")

	// get env

	port := os.Getenv("PORT")
	env := os.Getenv("ENV")

	fmt.Println("port: ", port)
	fmt.Println("env: ", env)
}
