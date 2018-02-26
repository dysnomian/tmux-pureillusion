package main

import "strings"

type Segment struct {
	fg      string
	bg      string
	content string
	bold    bool
}

func (seg Segment) Render() string {
	var bold string
	var fg string
	var bg string

	if seg.bold == true {
		bold = ",bold] "
	} else {
		bold = ",nobold] "
	}

	// If colors can't find the string, assume it's rendered dynamically
	if colors[seg.fg] == "" {
		fg = seg.fg
	} else {
		fg = colors[seg.fg]
	}

	if colors[seg.bg] == "" {
		bg = seg.bg
	} else {
		bg = colors[seg.bg]
	}

	seg_arr := []string{"#[fg=", fg, ",bg=", bg, bold, seg.content}

	return strings.Join(seg_arr, "")
}
