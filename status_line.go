package main

import (
	"strings"
)

func open_segment_left(next Segment) string {
	div := Divider{direction: "left", black: true}
	status_bar := Segment{fg: "pure-white", bg: "soft-gray", content: "", bold: false}

	return Transition{status_bar, next, div}.RenderSwapped()
}

func open_segment_right(next Segment) string {
	status_bar := Segment{fg: "pure-white", bg: "soft-gray", content: "", bold: false}
	div := Divider{direction: "right", black: true}

	return Transition{status_bar, next, div}.Render()
}

func close_segment_left(prev Segment) string {
	status_bar := Segment{fg: "soft-gray", bg: "soft-gray", content: "", bold: false}
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

	screen := Segment{fg: "soft-gray", bg: "flipflap-yellow", content: "#S ", bold: true}
	accent := Segment{fg: "pure-white", bg: "flipflap-gold", content: "", bold: false}
	name := Segment{fg: "abyssal-black", bg: "pure-white", content: "#(whoami) ", bold: true}
	papi := Segment{fg: "pure-white", bg: "papika-blue", content: "", bold: false}
	line := Segment{fg: "pure-white", bg: "abyssal-black", content: "#I:#P", bold: true}

	var status_arr = []string{screen.Render(), Transition{screen, accent, div}.Render(), accent.Render(), Transition{accent, name, div}.Render(), name.Render(), Transition{name, papi, div}.Render(), papi.Render(), Transition{papi, line, div}.Render(), close_segment_right(line)}

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
	div := Divider{direction: "left", black: true}

	cpu_icon := Segment{fg: "abyssal-black", bg: "flipflap-yellow", content: "cpu "}
	cpu_percent := Segment{fg: "abyssal-black", bg: "flipflap-gold", content: "#{cpu_percentage}", bold: true}
	battery := Segment{fg: "pure-white", bg: "yayaka-green-deeper", content: "#{battery_icon} #{battery_percentage}", bold: true}

	clock := Segment{fg: "soft-gray", bg: "pure-light-gray", content: "%a %h-%d %H:%M ", bold: true}

	var status_arr = []string{open_segment_left(cpu_icon), cpu_icon.Render(), Transition{cpu_icon, cpu_percent, div}.RenderSwapped(), cpu_percent.Render(), Transition{cpu_percent, battery, div}.RenderSwapped(), battery.Render(), Transition{battery, clock, div}.RenderSwapped(), clock.Render()}

	return strings.Join(status_arr, "")
}
