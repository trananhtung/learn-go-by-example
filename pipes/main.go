package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WordCount(f io.Reader) map[string]int {
	result := make(map[string]int)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result[scanner.Text()] += 1
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}

func main() {
	for key, value := range WordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", key, value)
	}

	// output pipe for next unix command
	os.Stdout.WriteString("next step\n")
}
