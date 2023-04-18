package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZendSort2(a any, b any, cmp types.CompareFuncT, swp types.SwapFuncT) {
	if cmp(a, b) > 0 {
		swp(a, b)
	}
}
func ZendSort3(a any, b any, c any, cmp types.CompareFuncT, swp types.SwapFuncT) {
	if cmp(a, b) <= 0 {
		if cmp(b, c) <= 0 {
			return
		}
		swp(b, c)
		if cmp(a, b) > 0 {
			swp(a, b)
		}
		return
	}
	if cmp(c, b) <= 0 {
		swp(a, c)
		return
	}
	swp(a, b)
	if cmp(b, c) > 0 {
		swp(b, c)
	}
}
func ZendSort4(
	a any,
	b any,
	c any,
	d any,
	cmp types.CompareFuncT,
	swp types.SwapFuncT,
) {
	ZendSort3(a, b, c, cmp, swp)
	if cmp(c, d) > 0 {
		swp(c, d)
		if cmp(b, c) > 0 {
			swp(b, c)
			if cmp(a, b) > 0 {
				swp(a, b)
			}
		}
	}
}
func ZendSort5(
	a any,
	b any,
	c any,
	d any,
	e any,
	cmp types.CompareFuncT,
	swp types.SwapFuncT,
) {
	ZendSort4(a, b, c, d, cmp, swp)
	if cmp(d, e) > 0 {
		swp(d, e)
		if cmp(c, d) > 0 {
			swp(c, d)
			if cmp(b, c) > 0 {
				swp(b, c)
				if cmp(a, b) > 0 {
					swp(a, b)
				}
			}
		}
	}
}
func ZendInsertSort(base any, nmemb int, siz int, cmp types.CompareFuncT, swp types.SwapFuncT) {
	switch nmemb {
	case 0:
		fallthrough
	case 1:

	case 2:
		ZendSort2(base, (*byte)(base+siz), cmp, swp)
	case 3:
		ZendSort3(base, (*byte)(base+siz), (*byte)(base+siz+siz), cmp, swp)
	case 4:
		var siz2 int = siz + siz
		ZendSort4(base, (*byte)(base+siz), (*byte)(base+siz2), (*byte)(base+siz+siz2), cmp, swp)
	case 5:
		var siz2 int = siz + siz
		ZendSort5(base, (*byte)(base+siz), (*byte)(base+siz2), (*byte)(base+siz+siz2), (*byte)(base+siz2+siz2), cmp, swp)
	default:
		var i *byte
		var j *byte
		var k *byte
		var start *byte = (*byte)(base)
		var end *byte = start + nmemb*siz
		var siz2 int = siz + siz
		var sentry *byte = start + 6*siz
		for i = start + siz; i < sentry; i += siz {
			j = i - siz
			if cmp(j, i) <= 0 {
				continue
			}
			for j != start {
				j -= siz
				if cmp(j, i) <= 0 {
					j += siz
					break
				}
			}
			for k = i; k > j; k -= siz {
				swp(k, k-siz)
			}
		}
		for i = sentry; i < end; i += siz {
			j = i - siz
			if cmp(j, i) <= 0 {
				continue
			}
			for {
				j -= siz2
				if cmp(j, i) <= 0 {
					j += siz
					if cmp(j, i) <= 0 {
						j += siz
					}
					break
				}
				if j == start {
					break
				}
				if j == start+siz {
					j -= siz
					if cmp(i, j) > 0 {
						j += siz
					}
					break
				}

			}
			for k = i; k > j; k -= siz {
				swp(k, k-siz)
			}
		}
	}
}
func ZendSort(base any, nmemb int, siz int, cmp types.CompareFuncT, swp types.SwapFuncT) {
	for true {
		if nmemb <= 16 {
			ZendInsertSort(base, nmemb, siz, cmp, swp)
			return
		} else {
			var i *byte
			var j *byte
			var start *byte = (*byte)(base)
			var end *byte = start + nmemb*siz
			var offset int = nmemb >> int64(1)
			var pivot *byte = start + offset*siz
			if nmemb>>int64(10) != 0 {
				var delta int = (offset >> int64(1)) * siz
				ZendSort5(start, start+delta, pivot, pivot+delta, end-siz, cmp, swp)
			} else {
				ZendSort3(start, pivot, end-siz, cmp, swp)
			}
			swp(start+siz, pivot)
			pivot = start + siz
			i = pivot + siz
			j = end - siz
			for true {
				for cmp(pivot, i) > 0 {
					i += siz
					if i == j {
						goto done
					}
				}
				j -= siz
				if j == i {
					goto done
				}
				for cmp(j, pivot) > 0 {
					j -= siz
					if j == i {
						goto done
					}
				}
				swp(i, j)
				i += siz
				if i == j {
					goto done
				}
			}
		done:
			swp(pivot, i-siz)
			if i-siz-start < end-i {
				ZendSort(start, (i-start)/siz-1, siz, cmp, swp)
				base = i
				nmemb = (end - i) / siz
			} else {
				ZendSort(i, (end-i)/siz, siz, cmp, swp)
				nmemb = (i-start)/siz - 1
			}
		}
	}
}
