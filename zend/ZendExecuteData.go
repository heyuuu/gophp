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

//             func MakeZendExecuteData(
// opline *ZendOp,
// call *ZendExecuteData,
// return_value *Zval,
// func_ *ZendFunction,
// This Zval,
// prev_execute_data *ZendExecuteData,
// symbol_table *ZendArray,
// run_time_cache *any,
// ) ZendExecuteData {
//                 return ZendExecuteData{
//                     opline:opline,
//                     call:call,
//                     return_value:return_value,
//                     func_:func_,
//                     This:This,
//                     prev_execute_data:prev_execute_data,
//                     symbol_table:symbol_table,
//                     run_time_cache:run_time_cache,
//                 }
//             }
func (this *ZendExecuteData) GetOpline() *ZendOp             { return this.opline }
func (this *ZendExecuteData) SetOpline(value *ZendOp)        { this.opline = value }
func (this *ZendExecuteData) GetCall() *ZendExecuteData      { return this.call }
func (this *ZendExecuteData) SetCall(value *ZendExecuteData) { this.call = value }

// func (this *ZendExecuteData)  GetReturnValue() *Zval      { return this.return_value }
func (this *ZendExecuteData) SetReturnValue(value *Zval)  { this.return_value = value }
func (this *ZendExecuteData) GetFunc() *ZendFunction      { return this.func_ }
func (this *ZendExecuteData) SetFunc(value *ZendFunction) { this.func_ = value }
func (this *ZendExecuteData) GetThis() Zval               { return this.This }

// func (this *ZendExecuteData) SetThis(value Zval) { this.This = value }
func (this *ZendExecuteData) GetPrevExecuteData() *ZendExecuteData { return this.prev_execute_data }
func (this *ZendExecuteData) SetPrevExecuteData(value *ZendExecuteData) {
	this.prev_execute_data = value
}
func (this *ZendExecuteData) GetSymbolTable() *ZendArray      { return this.symbol_table }
func (this *ZendExecuteData) SetSymbolTable(value *ZendArray) { this.symbol_table = value }

// func (this *ZendExecuteData)  GetRunTimeCache() *any      { return this.run_time_cache }
// func (this *ZendExecuteData) SetRunTimeCache(value *any) { this.run_time_cache = value }
