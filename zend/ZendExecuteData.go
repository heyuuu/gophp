// <<generate>>

package zend

/**
 * ZendExecuteData
 */
type ZendExecuteData struct {
	opline            *ZendOp
	call              *ZendExecuteData
	return_value      *Zval
	func_             *ZendFunction
	This              Zval
	prev_execute_data *ZendExecuteData
	symbol_table      *ZendArray
	run_time_cache    *any
}

func (this *ZendExecuteData) NumArgs() uint32 { return this.This.GetNumArgs() }

/**
 * Getter/Setter
 */
func (this *ZendExecuteData) GetOpline() *ZendOp                   { return this.opline }
func (this *ZendExecuteData) SetOpline(value *ZendOp)              { this.opline = value }
func (this *ZendExecuteData) GetCall() *ZendExecuteData            { return this.call }
func (this *ZendExecuteData) SetCall(value *ZendExecuteData)       { this.call = value }
func (this *ZendExecuteData) SetReturnValue(value *Zval)           { this.return_value = value }
func (this *ZendExecuteData) GetFunc() *ZendFunction               { return this.func_ }
func (this *ZendExecuteData) SetFunc(value *ZendFunction)          { this.func_ = value }
func (this *ZendExecuteData) GetThis() *Zval                       { return &this.This }
func (this *ZendExecuteData) GetPrevExecuteData() *ZendExecuteData { return this.prev_execute_data }
func (this *ZendExecuteData) SetPrevExecuteData(value *ZendExecuteData) {
	this.prev_execute_data = value
}
func (this *ZendExecuteData) GetSymbolTable() *ZendArray      { return this.symbol_table }
func (this *ZendExecuteData) SetSymbolTable(value *ZendArray) { this.symbol_table = value }
