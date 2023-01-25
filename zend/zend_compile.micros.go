// <<generate>>

package zend

// #define ZEND_COMPILE_H

// # include "zend.h"

// # include "zend_ast.h"

// # include < stdarg . h >

// # include "zend_llist.h"

// #define SET_UNUSED(op) op ## _type = IS_UNUSED

// # include "zend_globals.h"

// # include "zend_vm_opcodes.h"

// # include < zend_language_parser . h >

// # include "zend.h"

// # include "zend_compile.h"

// # include "zend_constants.h"

// # include "zend_llist.h"

// # include "zend_API.h"

// # include "zend_exceptions.h"

// # include "zend_interfaces.h"

// # include "zend_virtual_cwd.h"

// # include "zend_multibyte.h"

// # include "zend_language_scanner.h"

// # include "zend_inheritance.h"

// # include "zend_vm.h"

// #define SET_NODE(target,src) do { target ## _type = ( src ) -> op_type ; if ( ( src ) -> op_type == IS_CONST ) { target . constant = zend_add_literal ( & ( src ) -> u . constant ) ; } else { target = ( src ) -> u . op ; } } while ( 0 )

// #define GET_NODE(target,src) do { ( target ) -> op_type = src ## _type ; if ( ( target ) -> op_type == IS_CONST ) { ZVAL_COPY_VALUE ( & ( target ) -> u . constant , CT_CONSTANT ( src ) ) ; } else { ( target ) -> u . op = src ; } } while ( 0 )
