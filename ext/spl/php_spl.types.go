// <<generate>>

package spl

/**
 * ZendSplGlobals
 */
type ZendSplGlobals struct {
	autoload_extensions *zend.ZendString
	autoload_functions  *zend.HashTable
	hash_mask_handle    intPtr
	hash_mask_handlers  intPtr
	hash_mask_init      int
	autoload_running    int
}

func (this ZendSplGlobals) GetAutoloadExtensions() *zend.ZendString { return this.autoload_extensions }
func (this *ZendSplGlobals) SetAutoloadExtensions(value *zend.ZendString) {
	this.autoload_extensions = value
}
func (this ZendSplGlobals) GetAutoloadFunctions() *zend.HashTable { return this.autoload_functions }
func (this *ZendSplGlobals) SetAutoloadFunctions(value *zend.HashTable) {
	this.autoload_functions = value
}
func (this ZendSplGlobals) GetHashMaskHandle() intPtr         { return this.hash_mask_handle }
func (this *ZendSplGlobals) SetHashMaskHandle(value intPtr)   { this.hash_mask_handle = value }
func (this ZendSplGlobals) GetHashMaskHandlers() intPtr       { return this.hash_mask_handlers }
func (this *ZendSplGlobals) SetHashMaskHandlers(value intPtr) { this.hash_mask_handlers = value }
func (this ZendSplGlobals) GetHashMaskInit() int              { return this.hash_mask_init }
func (this *ZendSplGlobals) SetHashMaskInit(value int)        { this.hash_mask_init = value }
func (this ZendSplGlobals) GetAutoloadRunning() int           { return this.autoload_running }
func (this *ZendSplGlobals) SetAutoloadRunning(value int)     { this.autoload_running = value }

/**
 * AutoloadFuncInfo
 */
type AutoloadFuncInfo struct {
	func_ptr *zend.ZendFunction
	obj      zend.Zval
	closure  zend.Zval
	ce       *zend.ZendClassEntry
}

func (this AutoloadFuncInfo) GetFuncPtr() *zend.ZendFunction       { return this.func_ptr }
func (this *AutoloadFuncInfo) SetFuncPtr(value *zend.ZendFunction) { this.func_ptr = value }
func (this AutoloadFuncInfo) GetObj() zend.Zval                    { return this.obj }
func (this *AutoloadFuncInfo) SetObj(value zend.Zval)              { this.obj = value }
func (this AutoloadFuncInfo) GetClosure() zend.Zval                { return this.closure }
func (this *AutoloadFuncInfo) SetClosure(value zend.Zval)          { this.closure = value }
func (this AutoloadFuncInfo) GetCe() *zend.ZendClassEntry          { return this.ce }
func (this *AutoloadFuncInfo) SetCe(value *zend.ZendClassEntry)    { this.ce = value }
