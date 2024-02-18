package php

import "github.com/heyuuu/gophp/php/types"

/* zend_internal_function_handler */
type ZifHandler func(ex *ExecuteData, returnValue *types.Zval)
