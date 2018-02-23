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

	seg_arr := []string{"#[fg=", colors[seg.fg], ",bg=", colors[seg.bg], bold, seg.content}

	return strings.Join(seg_arr, "")
}

func separator(direction string, black bool) string {
	var right_separator string = string('')
	var right_separator_black string = string('')
	var left_separator string = string('')
	var left_separator_black string = string('')

	if direction == "left" {
		if black {
			return left_separator_black
		} else {
			return left_separator
		}
	} else {
		if black {
			return right_separator_black
		} else {
			return right_separator
		}
	}
}

func transition(prev segment, next segment, direction string, black bool) string {
	var trans_arr = []string{"#[fg=", colors[prev.bg], ",bg=", colors[next.bg], ",nobold]", separator(direction, black)}

	logger.Info("Next bg: %s (%s)", next.bg, colors[next.bg])

	return strings.Join(trans_arr, "")
}

func open_segment_right(next segment) string {
	status_bar := segment{fg: "pure-white", bg: "soft-gray", content: "", bold: false}

	return transition(status_bar, next, "right", true)
}

func open_segment_left(next segment) string {
	status_bar := segment{fg: "pure-white", bg: "soft-gray", content: "", bold: false}

	return transition(next, status_bar, "left", true)
}
func close_segment_right(prev segment) string {
	status_bar := segment{fg: "pure-white", bg: "soft-gray", content: "", bold: false}

	return transition(prev, status_bar, "right", true)
}

func status_left() string {
	seg1 := segment{fg: "flipflap-gold", bg: "flipflap-yellow", content: "#S ", bold: true}
	seg2 := segment{fg: "pure-pink", bg: "cocona-pink-deeper", content: "#(whoami) ", bold: true}
	seg3 := segment{fg: "cocona-blue", bg: "soft-gray", content: "#I:#P", bold: true}

	var status_arr = []string{to_string(seg1), transition(seg1, seg2, "right", true), to_string(seg2), transition(seg2, seg3, "right", true), to_string(seg3), close_segment_right(seg3)}

	return strings.Join(status_arr, "")
}

func active_window() string {
	num := segment{fg: "soft-gray", bg: "flipflap-yellow", content: "#I ", bold: true}
	accent := segment{fg: "pure-white", bg: "flipflap-gold", content: "", bold: false}
	name := segment{fg: "abyssal-black", bg: "pure-white", content: "#W ", bold: true}
	white := segment{fg: "pure-white", bg: "cocona-pink", content: "", bold: false}

	var window_arr = []string{open_segment_left(num), to_string(num), transition(num, accent, "right", true), transition(accent, name, "right", true), to_string(name), transition(name, white, "right", true), close_segment_right(white)}
	return strings.Join(window_arr, "")
}

func inactive_window() string {
	window := segment{fg: "pure-white", bg: "flipflap-rose", content: "#I #W ", bold: true}

	var window_arr = []string{open_segment_right(window), to_string(window), close_segment_right(window)}
	return strings.Join(window_arr, "")
}
