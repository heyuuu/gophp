// <<generate>>

package zend

const ZvalDtorFunc = RcDtorFunc
const ZvalPtrDtorWrapper = ZvalPtrDtor
const ZvalInternalPtrDtorWrapper = ZvalInternalPtrDtor
const ZVAL_PTR_DTOR DtorFuncT = ZvalPtrDtor
const ZVAL_INTERNAL_PTR_DTOR = ZvalInternalPtrDtor
const ZendStringDestroy = _efree

type ZendRcDtorFuncT func(p *ZendRefcounted)

var ZendRcDtorFunc []ZendRcDtorFuncT = []ZendRcDtorFuncT{ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendStringDestroy), ZendRcDtorFuncT(ZendArrayDestroy), ZendRcDtorFuncT(ZendObjectsStoreDel), ZendRcDtorFuncT(ZendListFree), ZendRcDtorFuncT(ZendReferenceDestroy), ZendRcDtorFuncT(ZendAstRefDestroy)}
