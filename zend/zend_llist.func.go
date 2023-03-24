package zend

func ZendLlistGetFirst(l *ZendLlist) any { return ZendLlistGetFirstEx(l, nil) }
func ZendLlistGetLast(l *ZendLlist) any  { return ZendLlistGetLastEx(l, nil) }
func ZendLlistGetNext(l *ZendLlist) any  { return ZendLlistGetNextEx(l, nil) }
func ZendLlistInit(l *ZendLlist, size int, dtor LlistDtorFuncT, persistent uint8) {
	l.Init(size, dtor, persistent)
}
func ZendLlistAddElement(l *ZendLlist, element any) {
	l.AddElement(element)
}
func ZendLlistDelElement(l *ZendLlist, element any, compare func(element1 any, element2 any) int) {
	l.DelElementByData(element, compare)
}
func ZendLlistApplyWithDel(l *ZendLlist, func_ func(data any) int) {
	l.ApplyWithDel(func_)
}
func ZendLlistApplyWithArgument(l *ZendLlist, func_ LlistApplyWithArgFuncT, arg any) {
	l.ApplyWithArgument(func_, arg)
}
func ZendLlistGetFirstEx(l *ZendLlist, pos *ZendLlistPosition) []byte {
	return l.GetFirstEx(pos)
}
func ZendLlistGetLastEx(l *ZendLlist, pos *ZendLlistPosition) any {
	return l.GetLastEx(pos)
}
func ZendLlistGetNextEx(l *ZendLlist, pos *ZendLlistPosition) any {
	return l.GetNextEx(pos)
}
