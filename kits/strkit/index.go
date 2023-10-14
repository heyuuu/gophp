package strkit

func IndexNewline(str string) (pos int, newlineLen int) {
	len_ := len(str)
	for i, c := range str {
		if c == '\r' {
			if i+1 < len_ && str[i+1] == '\n' {
				return i, 2
			} else {
				return i, 1
			}
		} else if c == '\n' {
			return i, 1
		}
	}
	return -1, 0
}

func IndexNewlineFrom(str string, from int) (pos int, newlineLen int) {
	len_ := len(str)
	for i := from; i < len_; i++ {
		c := str[i]
		if c == '\r' {
			if i+1 < len_ && str[i+1] == '\n' {
				return i, 2
			} else {
				return i, 1
			}
		} else if c == '\n' {
			return i, 1
		}
	}
	return -1, 0
}
