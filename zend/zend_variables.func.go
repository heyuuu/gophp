package zend

import (
	"github.com/heyuuu/gophp/zend/types"
)

func ZvalAddRef(p *types.Zval) {
	if p.IsRefcounted() {
		if p.IsReference() && p.GetRefcount() == 1 {
			types.ZVAL_COPY(p, types.Z_REFVAL_P(p))
		} else {
			// 			p.AddRefcount()
		}
	}
}
