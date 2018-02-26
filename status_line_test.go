package main

import "testing"

func TestSegmentRender(t *testing.T) {
	segment := Segment{fg: "pure-white", bg: "abyssal-black", content: "bleh!", bold: true}
	got := segment.Render()
	expected := "#[fg=#ffffff,bg=#000000,bold] bleh!"

	if got != expected {
		t.Errorf("Segment.Render() rendered incorrectly:\n  expected: '%s'\n  got:      '%s'\n", expected, got)
	}
}

func TestDividerRender(t *testing.T) {
	div := Divider{direction: "left", black: true}
	got := div.Render()
	expected := string("")

	if got != expected {
		t.Errorf("Divider.Render() rendered incorrectly:\n  expected: '%s'\n  got:      '%s'\n", expected, got)
	}
}

func TestTransitionRender(t *testing.T) {
	prev := Segment{fg: "pure-white", bg: "abyssal-black", content: "", bold: true}
	next := Segment{fg: "pure-white", bg: "cocona-blue", content: "", bold: true}
	div := Divider{direction: "right", black: true}
	transition := Transition{prev: prev, next: next, divider: div}

	got := transition.Render()
	expected := "#[fg=#000000,bg=#366dd2,nobold]"

	if got != expected {
		t.Errorf("Transition.Render() rendered incorrectly:\n  expected: '%s'\n  got:      '%s'\n", expected, got)
	}
}
