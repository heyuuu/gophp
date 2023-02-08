package zend

import b "sik/builtin"

var ZVAL_PTR_DTOR DtorFuncT = ZvalPtrDtor

type ZendRcDtorFuncT func(p *ZendRefcounted)

func __ZendRcDtorFuncTWrapper[T any](fun func(ptr *T)) ZendRcDtorFuncT {
	return func(p *ZendRefcounted) {
		fun(b.Cast[T](p))
	}
}

func __ZendRcDtorFuncTWrapper2[T any, R any](fun func(ptr *T) R) ZendRcDtorFuncT {
	return func(p *ZendRefcounted) {
		fun(b.Cast[T](p))
	}
}

var ZendRcDtorFuncMap = map[uint8]ZendRcDtorFuncT{
	IS_ARRAY:        __ZendRcDtorFuncTWrapper(ZendArrayDestroy),     // IS_ARRAY = 7
	IS_OBJECT:       __ZendRcDtorFuncTWrapper(ZendObjectsStoreDel),  // IS_OBJECT = 8
	IS_RESOURCE:     __ZendRcDtorFuncTWrapper2(ZendListFree),        // IS_RESOURCE = 9
	IS_REFERENCE:    __ZendRcDtorFuncTWrapper(ZendReferenceDestroy), // IS_REFERENCE = 10
	IS_CONSTANT_AST: __ZendRcDtorFuncTWrapper(ZendAstRefDestroy),    // IS_CONSTANT_AST = 11
}
