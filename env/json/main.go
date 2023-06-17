package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func LoadConfig(path, envPrefix string, config interface{}) error {
	err := LoadFile(path, config)
	if err != nil {
		return errors.Wrap(err, "failed to load config file")
	}
	return nil
}

func LoadFile(path string, config interface{}) error {
	file, err := os.Open(path)

	if err != nil {
		return errors.Wrap(err, "failed to open config file")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return errors.Wrap(err, "failed to decode config file")
	}

	return nil

}

type Config struct {
	Version string `json:"version" required:"true"`
	Secret  string `json:"secret" required:"true"`
	Counter int    `json:"counter"`
}

func main() {
	var err error

	tmpFolder := os.TempDir()

	tempFile := tmpFolder + "/config.json"

	file, err := os.Create(tempFile)

	if err != nil {
		panic(err)
	}

	// close file
	defer file.Close()
	defer os.Remove(tempFile)

	configFile := `{
	"version": "1.0.0",
	"secret": "secret key",
	"counter": 50	
}`

	// write some text line-by-line to file
	if _, err = file.Write([]byte(configFile)); err != nil {
		panic(err)
	}

	var config Config

	err = LoadConfig(tempFile, "", &config)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Final Config: %#v\n", config)
}
