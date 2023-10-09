package php

import "io"

func vmEcho(w io.Writer, zv Val) {
	str := ZvalGetString(zv)
	if len(str) > 0 {
		io.WriteString(w, str)
	}
}
