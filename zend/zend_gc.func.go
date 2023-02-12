// <<generate>>

package zend

func GC_MAY_LEAK(ref *ZendRefcounted) bool    { return false }
func GcCheckPossibleRoot(ref *ZendRefcounted) {}
func GcProtect(protect ZendBool)              {}
func GcPossibleRoot(ref *ZendRefcounted)      {}
