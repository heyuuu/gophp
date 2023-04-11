package standard

import "github.com/heyuuu/gophp/zend/types"

func CompareRight(a **byte, aend *byte, b **byte, bend *byte) int {
	var bias int = 0

	/* The longest run of digits wins.  That aside, the greatest
	   value wins, but we can't know that it will until we've scanned
	   both numbers to know that they have the same magnitude, so we
	   remember it in BIAS. */

	for {
		if ((*a) == aend || !(isdigit(int(uint8(*(*a)))))) && ((*b) == bend || !(isdigit(int(uint8(*(*b)))))) {
			return bias
		} else if (*a) == aend || !(isdigit(int(uint8(*(*a))))) {
			return -1
		} else if (*b) == bend || !(isdigit(int(uint8(*(*b))))) {
			return +1
		} else if (*(*a)) < (*(*b)) {
			if bias == 0 {
				bias = -1
			}
		} else if (*(*a)) > (*(*b)) {
			if bias == 0 {
				bias = +1
			}
		}
		*a++
		*b++
	}
	return 0
}
func CompareLeft(a **byte, aend *byte, b **byte, bend *byte) int {
	/* Compare two left-aligned numbers: the first to have a
	   different value wins. */

	for {
		if ((*a) == aend || !(isdigit(int(uint8(*(*a)))))) && ((*b) == bend || !(isdigit(int(uint8(*(*b)))))) {
			return 0
		} else if (*a) == aend || !(isdigit(int(uint8(*(*a))))) {
			return -1
		} else if (*b) == bend || !(isdigit(int(uint8(*(*b))))) {
			return +1
		} else if (*(*a)) < (*(*b)) {
			return -1
		} else if (*(*a)) > (*(*b)) {
			return +1
		}
		*a++
		*b++
	}
	return 0
}

func Strnatcmp(s1 string, s2 string, fold_case bool) int {
	return StrnatcmpEx(s1, s1, s2, s2, types.IntBool(fold_case))
}
func StrnatcmpEx(a *byte, a_len int, b *byte, b_len int, fold_case int) int {
	var ca uint8
	var cb uint8
	var ap *byte
	var bp byte
	var aend *byte = a + a_len
	var bend byte = b + b_len
	var fractional int
	var result int
	var leading short = 1
	if a_len == 0 || b_len == 0 {
		if a_len == b_len {
			return 0
		} else {
			if a_len > b_len {
				return 1
			} else {
				return -1
			}
		}
	}
	ap = a
	bp = b
	for true {
		ca = *ap
		cb = *bp

		/* skip over leading zeros */

		for leading && ca == '0' && ap+1 < aend && isdigit(int(uint8(*(ap + 1)))) {
			ca = *(b.PreInc(&ap))
		}
		for leading && cb == '0' && bp+1 < bend && isdigit(int(uint8(*(bp + 1)))) {
			cb = *(b.PreInc(&bp))
		}
		leading = 0

		/* Skip consecutive whitespace */

		for isspace(int(uint8(ca))) {
			ca = *(b.PreInc(&ap))
		}
		for isspace(int(uint8(cb))) {
			cb = *(b.PreInc(&bp))
		}

		/* process run of digits */

		if isdigit(int(uint8(ca))) && isdigit(int(uint8(cb))) {
			fractional = ca == '0' || cb == '0'
			if fractional != 0 {
				result = CompareLeft(&ap, aend, &bp, bend)
			} else {
				result = CompareRight(&ap, aend, &bp, bend)
			}
			if result != 0 {
				return result
			} else if ap == aend && bp == bend {

				/* End of the strings. Let caller sort them out. */

				return 0
			} else if ap == aend {
				return -1
			} else if bp == bend {
				return 1
			} else {

				/* Keep on comparing from the current point. */

				ca = *ap
				cb = *bp
			}
		}
		if fold_case != 0 {
			ca = toupper(int(uint8(ca)))
			cb = toupper(int(uint8(cb)))
		}
		if ca < cb {
			return -1
		} else if ca > cb {
			return +1
		}
		ap++
		bp++
		if ap >= aend && bp >= bend {

			/* The strings compare the same.  Perhaps the caller
			   will want to call strcmp to break the tie. */

			return 0
		} else if ap >= aend {
			return -1
		} else if bp >= bend {
			return 1
		}
	}
}
