package main

import (
	"strings"

	"github.com/google/logger"
)

type Transition struct {
	prev    Segment
	next    Segment
	divider Divider
}

func (tran Transition) Render() string {

	var trans_arr = []string{"#[fg=", colors[tran.prev.bg], ",bg=", colors[tran.next.bg], ",nobold]", tran.divider.Render()}

	logger.Info("Next bg: %s (%s)", tran.next.bg, colors[tran.next.bg])

	return strings.Join(trans_arr, "")
}

func (tran Transition) RenderWithFg() string {

	var trans_arr = []string{"#[fg=", colors[tran.prev.fg], ",bg=", colors[tran.next.bg], ",nobold]", tran.divider.Render()}

	logger.Info("Next bg: %s (%s)", tran.next.bg, colors[tran.next.bg])

	return strings.Join(trans_arr, "")
}

// For certain kinds of dividers that get weird on me
func (tran Transition) RenderSwapped() string {

	var trans_arr = []string{"#[fg=", colors[tran.next.bg], ",bg=", colors[tran.prev.bg], ",nobold]", tran.divider.Render()}

	logger.Info("Next bg: %s (%s)", tran.next.bg, colors[tran.next.bg])

	return strings.Join(trans_arr, "")
}
