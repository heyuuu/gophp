package standard

import (
	"github.com/heyuuu/gophp/kits/ascii"
)

func compareRight(a string, ap int, b string, bp int) (result int, newAp int, newBp int) {
	/* The longest run of digits wins.  That aside, the greatest
	   value wins, but we can't know that it will until we've scanned
	   both numbers to know that they have the same magnitude, so we
	   remember it in BIAS. */

	bias := 0
	for ; ; ap, bp = ap+1, bp+1 {
		if (ap == len(a) || !ascii.IsDigit(a[ap])) && (bp == len(b) || !ascii.IsDigit(b[bp])) {
			return bias, ap, bp
		} else if ap == len(a) || !ascii.IsDigit(a[ap]) {
			return -1, ap, bp
		} else if bp == len(b) || !ascii.IsDigit(b[bp]) {
			return 1, ap, bp
		} else if a[ap] < b[bp] {
			if bias == 0 {
				bias = -1
			}
		} else if a[ap] > b[bp] {
			if bias == 0 {
				bias = 1
			}
		}
	}
}
func compareLeft(a string, ap int, b string, bp int) (result int, newAp int, newBp int) {
	/* Compare two left-aligned numbers: the first to have a different value wins. */

	for ; ; ap, bp = ap+1, bp+1 {
		if (ap == len(a) || !ascii.IsDigit(a[ap])) && (bp == len(b) || !ascii.IsDigit(b[bp])) {
			return 0, ap, bp
		} else if ap == len(a) || !ascii.IsDigit(a[ap]) {
			return -1, ap, bp
		} else if bp == len(b) || !ascii.IsDigit(b[bp]) {
			return 1, ap, bp
		} else if a[ap] < b[bp] {
			return -1, ap, bp
		} else if a[ap] > b[bp] {
			return 1, ap, bp
		}
	}
}

func Strnatcmp(a string, b string, foldCase bool) int {
	if a == "" || b == "" {
		if a == b {
			return 0
		} else {
			if len(a) > len(b) {
				return 1
			} else {
				return -1
			}
		}
	}

	ap := 0 // for a
	bp := 0 // for b

	/* skip over leading zeros */
	for a[ap] == '0' && ap+1 < len(a) && ascii.IsDigit(a[ap+1]) {
		ap++
	}
	for b[bp] == '0' && bp+1 < len(b) && ascii.IsDigit(b[bp+1]) {
		bp++
	}

	for {
		/* Skip consecutive whitespace */
		for ascii.IsSpace(a[ap]) {
			ap++
		}
		for ascii.IsSpace(b[bp]) {
			bp++
		}

		/* process run of digits */
		if ascii.IsDigit(a[ap]) && ascii.IsDigit(b[bp]) {
			var result int
			if a[ap] == '0' || b[bp] == '0' {
				result, ap, bp = compareLeft(a, ap, b, bp)
			} else {
				result, ap, bp = compareRight(a, ap, b, bp)
			}
			if result != 0 {
				return result
			} else if ap == len(a) && bp == len(b) {
				/* End of the strings. Let caller sort them out. */
				return 0
			} else if ap == len(a) {
				return -1
			} else if bp == len(b) {
				return 1
			}
		}

		//
		ac := a[ap]
		bc := b[bp]
		if foldCase {
			ac = ascii.ToLower(ac)
			bc = ascii.ToLower(bc)
		}
		if ac < bc {
			return -1
		} else if ac > bc {
			return 1
		}
		ap++
		bp++
		if ap == len(a) && bp == len(b) {
			return 0
		} else if ap == len(a) {
			return -1
		} else if bp == len(b) {
			return 1
		}
	}
}
