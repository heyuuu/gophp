package core

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
)

func ICOPY_LIST(src *u_char, dst *u_char, last *u_char) {
	for {
		*((*int)(dst)) = *((*int)(src))
		src += ISIZE
		dst += ISIZE
		if src >= last {
			break
		}
	}
}
func ICOPY_ELT(src *u_char, dst *u_char, i int) {
	for {
		*((*int)(dst)) = *((*int)(src))
		src += ISIZE
		dst += ISIZE
		if !(b.AssignOp(&i, "-=", ISIZE)) {
			break
		}
	}
}
func CCOPY_LIST(src *int, dst *u_char, last *u_char) {
	for {
		*src++
		b.PostInc(&(*dst)) = (*src) - 1
		if src >= last {
			break
		}
	}
}
func CCOPY_ELT(src *int, dst *u_char, i int) {
	for {
		*src++
		b.PostInc(&(*dst)) = (*src) - 1
		if !(b.AssignOp(&i, "-=", 1)) {
			break
		}
	}
}
func EVAL(p *u_char) **u_char {
	return (**u_char)((*u_char)(0 + ((*u_char)(p+PSIZE-1-(*u_char)(0)) & ^(PSIZE - 1))))
}
func PhpMergesort(base any, nmemb int, size int, cmp func(any, any) int) int {
	var i int
	var sense int
	var big int
	var iflag int
	var f1 *u_char
	var f2 *u_char
	var t *u_char
	var b *u_char
	var tp2 *u_char
	var q *u_char
	var l1 *u_char
	var l2 *u_char
	var list2 *u_char
	var list1 *u_char
	var p2 *u_char
	var p *u_char
	var last *u_char
	var p1 **u_char
	if size < PSIZE/2 {
		errno = EINVAL
		return -1
	}
	if nmemb == 0 {
		return 0
	}

	/*
	 * XXX
	 * Stupid subtraction for the Cray.
	 */

	iflag = 0
	if size%ISIZE == 0 && (*byte)(base-(*byte)(0))%ISIZE == 0 {
		iflag = 1
	}
	if b.Assign(&list2, zend.Malloc(nmemb*size+PSIZE)) == nil {
		return -1
	}
	list1 = base
	Setup(list1, list2, nmemb, size, cmp)
	last = list2 + nmemb*size
	big = 0
	i = big
	for (*EVAL)(list2) != last {
		l2 = list1
		p1 = EVAL(list1)
		p2 = list2
		tp2 = p2
		for ; p2 != last; p1 = EVAL(l2) {
			p2 = (*EVAL)(p2)
			f1 = l2
			l1 = list1 + (p2 - list2)
			f2 = l1
			if p2 != last {
				p2 = (*EVAL)(p2)
			}
			l2 = list1 + (p2 - list2)
			for f1 < l1 && f2 < l2 {
				if (*cmp)(f1, f2) <= 0 {
					q = f2
					b = f1
					t = l1
					sense = -1
				} else {
					q = f1
					b = f2
					t = l2
					sense = 0
				}
				if big == 0 {
					for b.AssignOp(&b, "+=", size) < t && cmp(q, b) > sense {
						if b.PreInc(&i) == 6 {
							big = 1
							goto EXPONENTIAL
						}
					}
				} else {
				EXPONENTIAL:
					for i = size; ; i <<= 1 {
						if b.Assign(&p, b+i) >= t {
							if b.Assign(&p, t-size) > b && (*cmp)(q, p) <= sense {
								t = p
							} else {
								b = p
							}
							break
						} else if (*cmp)(q, p) <= sense {
							t = p
							if i == size {
								big = 0
							}
							goto FASTCASE
						} else {
							b = p
						}
					}
					for t > b+size {
						i = ((t - b) / size >> 1) * size
						if (*cmp)(q, b.Assign(&p, b+i)) <= sense {
							t = p
						} else {
							b = p
						}
					}
					goto COPY
				FASTCASE:
					for i > size {
						if (*cmp)(q, b.Assign(&p, b+b.AssignOp(&i, ">>=", 1))) <= sense {
							t = p
						} else {
							b = p
						}
					}
				COPY:
					b = t
				}
				i = size
				if q == f1 {
					if iflag != 0 {
						ICOPY_LIST(f2, tp2, b)
						ICOPY_ELT(f1, tp2, i)
					} else {
						CCOPY_LIST(f2, tp2, b)
						CCOPY_ELT(f1, tp2, i)
					}
				} else {
					if iflag != 0 {
						ICOPY_LIST(f1, tp2, b)
						ICOPY_ELT(f2, tp2, i)
					} else {
						CCOPY_LIST(f1, tp2, b)
						CCOPY_ELT(f2, tp2, i)
					}
				}
			}
			if f2 < l2 {
				if iflag != 0 {
					ICOPY_LIST(f2, tp2, l2)
				} else {
					CCOPY_LIST(f2, tp2, l2)
				}
			} else if f1 < l1 {
				if iflag != 0 {
					ICOPY_LIST(f1, tp2, l1)
				} else {
					CCOPY_LIST(f1, tp2, l1)
				}
			}
			*p1 = l2
		}
		tp2 = list1
		list1 = list2
		list2 = tp2
		last = list2 + nmemb*size
	}
	if base == list2 {
		memmove(list2, list1, nmemb*size)
		list2 = list1
	}
	zend.Free(list2)
	return 0
}
func Swap(a *u_char, b *u_char) {
	s = b
	i = size
	for {
		tmp = *a
		b.PostInc(&(*a)) = *s
		b.PostInc(&(*s)) = tmp
		if !(b.PreDec(&i)) {
			break
		}
	}
	a -= size
}
func Reverse(bot *u_char, top int) {
	s = top
	for {
		i = size
		for {
			tmp = *bot
			b.PostInc(&(*bot)) = *s
			b.PostInc(&(*s)) = tmp
			if !(b.PreDec(&i)) {
				break
			}
		}
		s -= size2
		if bot >= s {
			break
		}
	}
}
func Setup(list1 *u_char, list2 *u_char, n int, size int, cmp func(any, any) int) {
	var i int
	var length int
	var size2 int
	var sense int
	var f1 *u_char
	var f2 *u_char
	var s *u_char
	var l2 *u_char
	var last *u_char
	var p2 *u_char
	var tmp u_char
	size2 = size * 2
	if n <= 5 {
		Insertionsort(list1, n, size, cmp)
		(*EVAL)(list2) = (*u_char)(list2 + n*size)
		return
	}

	/*
	 * Avoid running pointers out of bounds; limit n to evens
	 * for simplicity.
	 */

	i = 4 + (n & 1)
	Insertionsort(list1+(n-i)*size, i, size, cmp)
	last = list1 + size*(n-i)
	(*EVAL)(list2 + (last - list1)) = list2 + n*size
	p2 = list2
	f1 = list1
	sense = cmp(f1, f1+size) > 0
	for ; f1 < last; sense = !sense {
		length = 2

		/* Find pairs with same sense. */

		for f2 = f1 + size2; f2 < last; f2 += size2 {
			if cmp(f2, f2+size) > 0 != sense {
				break
			}
			length += 2
		}
		if length < THRESHOLD {
			for {
				(*EVAL)(p2) = f1 + size2 - list1 + list2
				p2 = (*EVAL)(p2)
				if sense > 0 {
					Swap(f1, f1+size)
				}
				if b.AssignOp(&f1, "+=", size2) >= f2 {
					break
				}
			}
		} else {
			l2 = f2
			for f2 = f1 + size2; f2 < l2; f2 += size2 {
				if cmp(f2-size, f2) > 0 != sense {
					(*EVAL)(p2) = f2 - list1 + list2
					p2 = (*EVAL)(p2)
					if sense > 0 {
						Reverse(f1, f2-size)
					}
					f1 = f2
				}
			}
			if sense > 0 {
				Reverse(f1, f2-size)
			}
			f1 = f2
			if f2 < last || cmp(f2-size, f2) > 0 {
				(*EVAL)(p2) = f2 - list1 + list2
				p2 = (*EVAL)(p2)
			} else {
				(*EVAL)(p2) = list2 + n*size
				p2 = (*EVAL)(p2)
			}
		}
	}
}
func Insertionsort(a *u_char, n int, size int, cmp func(any, any) int) {
	var ai *u_char
	var s *u_char
	var t *u_char
	var u *u_char
	var tmp u_char
	var i int
	for ai = a + size; b.PreDec(&n) >= 1; ai += size {
		for t = ai; t > a; t -= size {
			u = t - size
			if cmp(u, t) <= 0 {
				break
			}
			Swap(u, t)
		}
	}
}
