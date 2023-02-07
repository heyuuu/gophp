// <<generate>>

package zend

func GC_G__() *ZendGcGlobals                         { return &GcGlobals }
func CG__() *ZendCompilerGlobals                     { return &CompilerGlobals }
func EG__() *ZendExecutorGlobals                     { return &ExecutorGlobals }
func INI_SCNG__() *ZendIniScannerGlobals             { return &IniScannerGlobals }
func LANG_SCNG__() *ZendPhpScannerGlobals            { return &LanguageScannerGlobals }
func FC__() *ZendFileContext                         { return CG__().GetFileContext() }
func _zendHashForeach(ht *HashTable) **Bucket        { return ht.foreachData() }
func _zendHashForeachReverse(ht *HashTable) **Bucket { return ht.foreachDataReserve() }
