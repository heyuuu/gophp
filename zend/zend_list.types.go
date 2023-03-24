package zend

/**
 * ZendRsrcListDtorsEntry
 */
type ZendRsrcListDtorsEntry struct {
	list_dtor_ex  RsrcDtorFuncT
	plist_dtor_ex RsrcDtorFuncT
	type_name     *byte
	module_number int
	resource_id   int
}

// func MakeZendRsrcListDtorsEntry(list_dtor_ex RsrcDtorFuncT, plist_dtor_ex RsrcDtorFuncT, type_name *byte, module_number int, resource_id int) ZendRsrcListDtorsEntry {
//     return ZendRsrcListDtorsEntry{
//         list_dtor_ex:list_dtor_ex,
//         plist_dtor_ex:plist_dtor_ex,
//         type_name:type_name,
//         module_number:module_number,
//         resource_id:resource_id,
//     }
// }
func (this *ZendRsrcListDtorsEntry) GetListDtorEx() RsrcDtorFuncT       { return this.list_dtor_ex }
func (this *ZendRsrcListDtorsEntry) SetListDtorEx(value RsrcDtorFuncT)  { this.list_dtor_ex = value }
func (this *ZendRsrcListDtorsEntry) GetPlistDtorEx() RsrcDtorFuncT      { return this.plist_dtor_ex }
func (this *ZendRsrcListDtorsEntry) SetPlistDtorEx(value RsrcDtorFuncT) { this.plist_dtor_ex = value }
func (this *ZendRsrcListDtorsEntry) GetTypeName() *byte                 { return this.type_name }
func (this *ZendRsrcListDtorsEntry) SetTypeName(value *byte)            { this.type_name = value }
func (this *ZendRsrcListDtorsEntry) GetModuleNumber() int               { return this.module_number }
func (this *ZendRsrcListDtorsEntry) SetModuleNumber(value int)          { this.module_number = value }
func (this *ZendRsrcListDtorsEntry) GetResourceId() int                 { return this.resource_id }
func (this *ZendRsrcListDtorsEntry) SetResourceId(value int)            { this.resource_id = value }
