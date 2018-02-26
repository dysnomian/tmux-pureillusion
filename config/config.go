package config

import (
	"io/ioutil"

	"github.com/google/logger"
)

func read_json(filename string) []byte {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Error("Couldn't read json file %s: %s", filename, err)
	}

	return buffer
}
