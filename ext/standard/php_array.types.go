// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * ZendArrayGlobals
 */
type ZendArrayGlobals struct {
	multisort_func *zend.CompareFuncT
}

func (this ZendArrayGlobals) GetMultisortFunc() *zend.CompareFuncT       { return this.multisort_func }
func (this *ZendArrayGlobals) SetMultisortFunc(value *zend.CompareFuncT) { this.multisort_func = value }
