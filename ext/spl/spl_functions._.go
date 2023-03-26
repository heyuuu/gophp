package spl

import (
	"sik/zend/types"
)

type CreateObjectFuncT func(class_type *types.ClassEntry) *types.ZendObject
