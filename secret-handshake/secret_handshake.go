package secret

func Handshake(code int) []string {
	var binaryCodes = []struct {
		code    int
		message string
	}{
		{1, "wink"},
		{2, "double blink"},
		{4, "close your eyes"},
		{8, "jump"},
	}

	var strings []string

	if (code < 1) || (code > 31) {
		return strings
	}

	for _, binaryCode := range binaryCodes {
		if code&binaryCode.code == binaryCode.code {
			strings = append(strings, binaryCode.message)
		}
	}

	length := len(strings)
	if code&16 == 16 && length > 0 {
		for i := 0; i < length/2; i++ {
			strings[i], strings[length-1-i] = strings[length-1-i], strings[i]
		}
	}
	return strings
}
