package date

func daysInMouth(year int, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	default: // 2
		if isLeapYear(year) {
			return 29
		} else {
			return 28
		}
	}
}

func isLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	return year%4 == 0
}

func englishSuffix(m int) string {
	if m >= 10 && m <= 19 {
		return "th"
	}
	switch m % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

func cond[T any](c bool, trueValue T, falseValue T) T {
	if c {
		return trueValue
	}
	return falseValue
}
