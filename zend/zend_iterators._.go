package zend

import "sik/zend/types"

var ZendIteratorClassEntry types.ClassEntry
var IteratorObjectHandlers ZendObjectHandlers = MakeZendObjectHandlers(0, IterWrapperFree, IterWrapperDtor, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, IterWrapperGetGc, nil, nil)
