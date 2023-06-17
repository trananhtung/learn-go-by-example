package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type TOMLData struct {
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

func (t *TOMLData) ToTOMLConfig() error {
	buf := &bytes.Buffer{}

	encoder := toml.NewEncoder(buf)

	if err := encoder.Encode(t); err != nil {
		return err
	}

	// write to toml file
	os.WriteFile("config.toml", buf.Bytes(), 0644)

	return nil
}

func (t *TOMLData) FromTOMLConfig() error {
	// read from toml file
	file, err := os.ReadFile("config.toml")

	if err != nil {
		return err
	}

	err = toml.Unmarshal(file, &t)
	if err != nil {
		return err
	}

	// print pretty struct
	fmt.Printf("%+v\n", t)

	return nil
}

func main() {
	person := TOMLData{
		Name: "John Doe",
		Age:  42,
	}

	err := person.ToTOMLConfig()
	if err != nil {
		panic(err)
	}

	err = person.FromTOMLConfig()
	if err != nil {
		panic(err)
	}

}
