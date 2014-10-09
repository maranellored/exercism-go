package leap

func IsLeapYear(year int) bool {
	if (year%4 == 0) && (year%100 != 0) {
		return true
	}

	return year%400 == 0
}
