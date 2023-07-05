package main

import (
	"io"
	"log"
	"os"
)

func CopyFile(reader io.Reader, writer io.Writer) error {
	buff := make([]byte, 128)

	for {
		m, err := reader.Read(buff)
		if err == io.EOF || m == 0 {
			break
		}

		if err != nil {
			return err
		}

		if _, err := writer.Write(buff[:m]); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	cfile, err := os.Open("source.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer cfile.Close()

	pfile, err := os.Create("target.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer pfile.Close()

	if err := CopyFile(cfile, pfile); err != nil {
		log.Fatal(err)
	}
}
