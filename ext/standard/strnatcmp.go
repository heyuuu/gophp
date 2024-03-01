package standard

import (
	"cmp"
	"github.com/heyuuu/gophp/kits/ascii"
)

func compareRight(a, b string) (result int, leftA string, leftB string) {
	/* The longest run of digits wins.  That aside, the greatest
	   value wins, but we can't know that it will until we've scanned
	   both numbers to know that they have the same magnitude, so we
	   remember it in BIAS. */
	bias := 0
	for i := 0; ; i++ {
		eof1 := i == len(a) || !ascii.IsDigit(a[i])
		eof2 := i == len(b) || !ascii.IsDigit(b[i])
		if eof1 && eof2 {
			return bias, a[i:], b[i:]
		} else if eof1 {
			return -1, a[i:], b[i:]
		} else if eof2 {
			return 1, a[i:], b[i:]
		}
		if bias == 0 {
			bias = cmp.Compare(a[i], b[i])
		}
	}
}

func compareLeft(a, b string) (result int, leftA string, leftB string) {
	for i := 0; ; i++ {
		eof1 := i == len(a) || !ascii.IsDigit(a[i])
		eof2 := i == len(b) || !ascii.IsDigit(b[i])

		if eof1 && eof2 {
			return 0, a[i:], b[i:]
		} else if eof1 {
			return -1, a[i:], b[i:]
		} else if eof2 {
			return 1, a[i:], b[i:]
		}

		result = cmp.Compare(a[i], b[i])
		if result != 0 {
			return result, a[i:], b[i:]
		}
	}
}

func Strnatcmp(a string, b string, foldCase bool) int {
	if a == "" || b == "" {
		return cmp.Compare(len(a), len(b))
	}

	/* skip over leading zeros */
	for len(a) >= 2 && a[0] == '0' && ascii.IsDigit(a[1]) {
		a = a[1:]
	}
	for len(b) >= 2 && b[0] == '0' && ascii.IsDigit(b[1]) {
		b = b[1:]
	}

	var result int
	for a != "" && b != "" {
		if ascii.IsSpace(a[0]) || ascii.IsSpace(b[0]) {
			/* Skip consecutive whitespace */
			for a != "" && ascii.IsSpace(a[0]) {
				a = a[1:]
			}
			for b != "" && ascii.IsSpace(b[0]) {
				b = b[1:]
			}
		} else if ascii.IsDigit(a[0]) && ascii.IsDigit(b[0]) {
			/* process run of digits */
			if a[0] == '0' || b[0] == '0' {
				result, a, b = compareLeft(a, b)
			} else {
				result, a, b = compareRight(a, b)
			}
			if result != 0 {
				return result
			}
		} else {
			/* process run of char */
			if foldCase {
				result = cmp.Compare(ascii.ToLower(a[0]), ascii.ToLower(b[0]))
			} else {
				result = cmp.Compare(a[0], b[0])
			}
			if result != 0 {
				return result
			}
			a, b = a[1:], b[1:]
		}
	}

	return cmp.Compare(len(a), len(b))
}
