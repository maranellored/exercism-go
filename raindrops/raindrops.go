package raindrops

import (
	"bytes"
	"strconv"
)

func Convert(input int) string {
	var string_buffer bytes.Buffer

	if input%3 == 0 {
		string_buffer.WriteString("Pling")
	}

	if input%5 == 0 {
		string_buffer.WriteString("Plang")
	}

	if input%7 == 0 {
		string_buffer.WriteString("Plong")
	}

	if string_buffer.Len() > 0 {
		return string_buffer.String()
	}

	return strconv.Itoa(input)
}
