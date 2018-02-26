package main

import "testing"

func TestGetColor(t *testing.T) {
	got := get_color("pure-white")
	expected := "#ffffff"

	if got != expected {
		t.Errorf("get_color rendered incorrectly:\n  expected: %s\ngot:    %s\n", expected, got)
	}
}
