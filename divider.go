package main

const (
	right_separator       string = string('')
	right_separator_black string = string('')
	left_separator        string = string('')
	left_separator_black  string = string('')
)

type Divider struct {
	direction string
	black     bool
}

func (divider Divider) Render() string {
	if divider.direction == "left" {
		if divider.black {
			return left_separator_black
		} else {
			return left_separator
		}
	} else {
		if divider.black {
			return right_separator_black
		} else {
			return right_separator
		}
	}
}
