package zend

import "sik/zend/types"

func GC_MAY_LEAK(ref *types.ZendRefcounted) bool    { return false }
func GcCheckPossibleRoot(ref *types.ZendRefcounted) {}
func GcProtect(protect types.ZendBool)              {}
func GcPossibleRoot(ref *types.ZendRefcounted)      {}
