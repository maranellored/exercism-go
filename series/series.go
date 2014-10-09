package slice

func Frist(n int, s string) string {
	return s[:n]
}

func All(n int, s string) []string {
	var validNumbers []string

	for i, _ := range s {
		if (i + n) > len(s) {
			break
		}
		validNumbers = append(validNumbers, s[i:i+n])
	}
	return validNumbers
}

func First(n int, s string) (first string, ok bool) {

	if n > len(s) {
		return "", false
	}
	return Frist(n, s), true
}
