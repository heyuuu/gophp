// <<generate>>

package zend

func __GC_G() *ZendGcGlobals                         { return &GcGlobals }
func __CG() *ZendCompilerGlobals                     { return &CompilerGlobals }
func __EG() *ZendExecutorGlobals                     { return &ExecutorGlobals }
func __INI_SCNG() *ZendIniScannerGlobals             { return &IniScannerGlobals }
func __LANG_SCNG() *ZendPhpScannerGlobals            { return &LanguageScannerGlobals }
func _zendHashForeach(ht *HashTable) **Bucket        { return ht.foreachData() }
func _zendHashForeachReverse(ht *HashTable) **Bucket { return ht.foreachDataReserve() }
