package main

import "testing"

func TestToString(t *testing.T) {
	segment := segment{fg: "pure-white", bg: "abyssal-black", content: "bleh!", bold: true}
	got := to_string(segment)
	expected := "#[fg=#FFFFFF,bg=#000000,bold] bleh! "

	if got != expected {
		t.Errorf("to_string rendered incorrectly:\n  expected: %s\ngot:    %s\n", expected, got)
	}
}
