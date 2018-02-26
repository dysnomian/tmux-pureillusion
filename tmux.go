package main

import (
	"os/exec"

	"github.com/buger/jsonparser"
	"github.com/google/logger"
)

func tmux_set(item string, value string) {
	k := string(item)
	v := string(value)

	logger.Info("tmux set -g %s %s\n", k, v)
	cmd := exec.Command("tmux", "set", "-g", k, v)

	stdoutStderr, err := cmd.CombinedOutput()
	logger.Info("%s", stdoutStderr)

	if err != nil {
		logger.Fatal("Failed writing tmux setting: { %s: %s }. %s\n", k, v, err)
	}
}

func tmux_setw(item string, value string) {
	k := string(item)
	v := string(value)

	logger.Info("tmux setw -g %s %s\n", k, v)
	cmd := exec.Command("tmux", "setw", "-g", k, v)

	stdoutStderr, err := cmd.CombinedOutput()
	logger.Info("%s", stdoutStderr)

	if err != nil {
		logger.Fatal("Failed writing tmux setting: { %s: %s }. %s\n", k, v, err)
	}
}

func tmux_set_color(key string, value string) {
	v := colors[string(value)]
	tmux_set(key, v)
}

func update_colors(background string) {
	filename := colors_filename(background)
	term_colors := read_json(filename)

	jsonparser.ObjectEach(term_colors, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		tmux_set_color(string(key), string(value))
		return nil
	})
}

func update_settings(filename string) {
	term_colors := read_json(filename)

	jsonparser.ObjectEach(term_colors, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		tmux_set(string(key), string(value))
		return nil
	})
}
