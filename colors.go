package main

import (
	"strings"

	"github.com/buger/jsonparser"
)

var colors map[string]string = color_map(relativePath("colors.json"))

func colors_filename(background string) string {
	var filename []string

	filename = append(filename, "tmux-colors-")
	filename = append(filename, background)
	filename = append(filename, ".json")

	return strings.Join(filename, "")
}

func get_color(key string) string {
	return colors[key]
}

func color_map(filename string) map[string]string {
	json := read_json(filename)

	var c map[string]string
	c = make(map[string]string)

	jsonparser.ObjectEach(json, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {

		c[string(key)] = string(value)
		return nil
	})

	return c
}
