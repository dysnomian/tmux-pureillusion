package main

import (
	"strings"

	"github.com/google/logger"
)

type segment struct {
	fg      string
	bg      string
	content string
	bold    bool
}

func to_string(seg segment) string {
	var bold string
	if seg.bold == true {
		bold = ",bold] "
	} else {
		bold = ",nobold] "
	}

	seg_arr := []string{"#[fg=", colors[seg.fg], ",bg=", colors[seg.bg], bold, seg.content, " "}

	return strings.Join(seg_arr, "")
}

func transition(prev segment, next segment, divider string) string {

	var trans_arr = []string{"#[fg=", colors[prev.bg], ",bg=", colors[next.bg], ",nobold]", divider}

	logger.Info("Next bg: %s (%s)", next.bg, colors[next.bg])

	return strings.Join(trans_arr, "")
}

func close_status_left(prev segment) string {
	status_bar := segment{fg: "pure-white", bg: "soft-gray", content: "", bold: false}

	return transition(prev, status_bar, string(''))
}

func status_left() string {
	divider := string('')

	seg1 := segment{fg: "flipflap-gold", bg: "flipflap-yellow", content: "#S", bold: true}
	seg2 := segment{fg: "pure-pink", bg: "cocona-pink-deeper", content: "#(whoami)", bold: true}
	seg3 := segment{fg: "cocona-blue", bg: "soft-gray", content: "#I:#P", bold: true}

	var status_arr = []string{to_string(seg1), transition(seg1, seg2, divider), to_string(seg2), transition(seg2, seg3, divider), to_string(seg3), close_status_left(seg3)}

	return strings.Join(status_arr, "")
}
