package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZvalAddRef(p *types2.Zval) {
	if p.IsRefcounted() {
		if p.IsReference() && p.GetRefcount() == 1 {
			types2.ZVAL_COPY(p, types2.Z_REFVAL_P(p))
		} else {
			// 			p.AddRefcount()
		}
	}
}
