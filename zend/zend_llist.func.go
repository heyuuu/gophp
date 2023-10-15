package zend

func ZendLlistGetFirst[T any](l *ZendLlist[T]) any          { return ZendLlistGetFirstEx(l, nil) }
func ZendLlistGetNext[T any](l *ZendLlist[T]) any           { return ZendLlistGetNextEx(l, nil) }
func ZendLlistAddElement[T any](l *ZendLlist[T], element T) { l.AddElement(element) }
func ZendLlistDelElement[T any](l *ZendLlist[T], element T, compare func(element1 T, element2 T) int) {
	l.DelElementByData(element, compare)
}
func ZendLlistApplyWithArgument[T any](l *ZendLlist[T], func_ LlistApplyWithArgFuncT, arg any) {
	l.ApplyWithArgument(func_, arg)
}
func ZendLlistGetFirstEx[T any](l *ZendLlist[T], pos *ZendLlistPosition[T]) T {
	return l.GetFirstEx(pos)
}
func ZendLlistGetNextEx[T any](l *ZendLlist[T], pos *ZendLlistPosition[T]) T {
	return l.GetNextEx(pos)
}
