package main

import (
	"strings"
)

func open_segment_right(next Segment) string {
	status_bar := Segment{fg: "pure-white", bg: "soft-gray", content: "", bold: false}
	div := Divider{direction: "right", black: true}

	return Transition{status_bar, next, div}.Render()
}

func open_segment_left(next Segment) string {
	div := Divider{direction: "left", black: true}
	status_bar := Segment{fg: "pure-white", bg: "soft-gray", content: "", bold: false}

	return Transition{status_bar, next, div}.Render()
}

func close_segment_left(prev Segment) string {
	status_bar := Segment{fg: "pure-white", bg: "soft-gray", content: "", bold: false}
	div := Divider{direction: "left", black: true}

	return Transition{prev, status_bar, div}.Render()
}

func close_segment_right(prev Segment) string {
	status_bar := Segment{fg: "pure-white", bg: "soft-gray", content: "", bold: false}
	div := Divider{direction: "right", black: true}

	return Transition{prev, status_bar, div}.Render()
}

func status_left() string {
	div := Divider{direction: "right", black: true}

	seg1 := Segment{fg: "flipflap-gold", bg: "flipflap-yellow", content: "#S ", bold: true}
	seg2 := Segment{fg: "pure-pink", bg: "cocona-pink-deeper", content: "#(whoami) ", bold: true}
	seg3 := Segment{fg: "cocona-blue", bg: "soft-gray", content: "#I:#P", bold: true}

	var status_arr = []string{seg1.Render(), Transition{seg1, seg2, div}.Render(), seg2.Render(), Transition{seg2, seg3, div}.Render(), seg3.Render(), close_segment_right(seg3)}

	return strings.Join(status_arr, "")
}

func active_window() string {
	div := Divider{direction: "right", black: true}

	num := Segment{fg: "soft-gray", bg: "flipflap-yellow", content: "#I ", bold: true}
	accent := Segment{fg: "pure-white", bg: "flipflap-gold", content: "", bold: false}
	name := Segment{fg: "abyssal-black", bg: "pure-white", content: "#W ", bold: true}
	white := Segment{fg: "pure-white", bg: "cocona-pink", content: "", bold: false}

	num_accent := Transition{num, accent, div}
	accent_name := Transition{accent, name, div}
	name_white := Transition{name, white, div}

	var window_arr = []string{open_segment_right(num), num.Render(), num_accent.Render(), accent.Render(), accent_name.Render(), name.Render(), name_white.Render(), close_segment_right(white)}
	return strings.Join(window_arr, "")
}

func inactive_window() string {
	div := Divider{direction: "right", black: false}

	num := Segment{fg: "pure-white", bg: "flipflap-rose", content: "#I ", bold: true}
	name := Segment{fg: "pure-white", bg: "flipflap-rose", content: "#W ", bold: true}

	num_name := Transition{num, name, div}

	var window_arr = []string{open_segment_right(num), num.Render(), num_name.RenderWithFg(), name.Render(), close_segment_right(name)}
	return strings.Join(window_arr, "")
}

func last_window() string {
	var window_arr = []string{"default,fg=", colors["abyssal-black"]}
	return strings.Join(window_arr, "")
}

func status_right() string {
	// set -g status-right '#{cpu_bg_color} CPU: #{cpu_icon} #{cpu_percentage} | %a %h-%d %H:%M '
	div := Divider{direction: "left", black: true}

	cpu := Segment{fg: "soft-gray", bg: "#{cpu_bg_color}", content: "#{cpu_icon} #{cpu_percentage}", bold: true}
	battery := Segment{fg: "pure-white", bg: "#{battery_status_bg}", content: "#{battery_icon} #{battery_percentage}", bold: true}
	clock := Segment{fg: "cocona-blue", bg: "soft-gray", content: "%a %h-%d %H:%M ", bold: true}

	var status_arr = []string{open_segment_left(cpu), cpu.Render(), Transition{cpu, battery, div}.Render(), battery.Render(), Transition{battery, clock, div}.Render(), clock.Render(), close_segment_right(clock)}

	return strings.Join(status_arr, "")
}
