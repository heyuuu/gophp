// <<generate>>

package zend

type LlistDtorFuncT func(any)
type LlistCompareFuncT func(**ZendLlistElement, **ZendLlistElement) int
type LlistApplyWithArgsFuncT func(data any, num_args int, args ...any)
type LlistApplyWithArgFuncT func(data any, arg any)
type LlistApplyFuncT func(any)

type ZendLlistPosition = *ZendLlistElement
