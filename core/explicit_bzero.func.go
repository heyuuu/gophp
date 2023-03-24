package core

func PhpExplicitBzero(dst any, siz int) {
	var i int = 0
	var buf *uint8 = (*volatile__unsigned__char)(dst)
	for ; i < siz; i++ {
		buf[i] = 0
	}
}
