// <<generate>>

package zend

// #define ZEND_OPERATORS_H

// # include < errno . h >

// # include < math . h >

// # include < assert . h >

// # include < stddef . h >

// # include "zend_portability.h"

// # include "zend_strtod.h"

// # include "zend_multiply.h"

// # include "zend_object_handlers.h"

// #define convert_to_ex_master(pzv,lower_type,upper_type) if ( Z_TYPE_P ( pzv ) != upper_type ) { convert_to_ ## lower_type ( pzv ) ; }

// #define convert_to_long_ex(pzv) convert_to_ex_master ( pzv , long , IS_LONG )

// #define convert_to_double_ex(pzv) convert_to_ex_master ( pzv , double , IS_DOUBLE )

// #define zend_update_current_locale()

// #define ZEND_TRY_BINARY_OP1_OBJECT_OPERATION(opcode,binary_op) if ( UNEXPECTED ( Z_TYPE_P ( op1 ) == IS_OBJECT ) && op1 == result && UNEXPECTED ( Z_OBJ_HANDLER_P ( op1 , get ) ) && EXPECTED ( Z_OBJ_HANDLER_P ( op1 , set ) ) ) { int ret ; zval rv ; zval * objval = Z_OBJ_HANDLER_P ( op1 , get ) ( op1 , & rv ) ; Z_TRY_ADDREF_P ( objval ) ; ret = binary_op ( objval , objval , op2 ) ; Z_OBJ_HANDLER_P ( op1 , set ) ( op1 , objval ) ; zval_ptr_dtor ( objval ) ; return ret ; } else if ( UNEXPECTED ( Z_TYPE_P ( op1 ) == IS_OBJECT ) && UNEXPECTED ( Z_OBJ_HANDLER_P ( op1 , do_operation ) ) ) { if ( EXPECTED ( SUCCESS == Z_OBJ_HANDLER_P ( op1 , do_operation ) ( opcode , result , op1 , op2 ) ) ) { return SUCCESS ; } }

// #define ZEND_TRY_BINARY_OP2_OBJECT_OPERATION(opcode) if ( UNEXPECTED ( Z_TYPE_P ( op2 ) == IS_OBJECT ) && UNEXPECTED ( Z_OBJ_HANDLER_P ( op2 , do_operation ) ) && EXPECTED ( SUCCESS == Z_OBJ_HANDLER_P ( op2 , do_operation ) ( opcode , result , op1 , op2 ) ) ) { return SUCCESS ; }

// #define ZEND_TRY_BINARY_OBJECT_OPERATION(opcode,binary_op) ZEND_TRY_BINARY_OP1_OBJECT_OPERATION ( opcode , binary_op ) else ZEND_TRY_BINARY_OP2_OBJECT_OPERATION ( opcode )

// #define ZEND_TRY_UNARY_OBJECT_OPERATION(opcode) if ( UNEXPECTED ( Z_TYPE_P ( op1 ) == IS_OBJECT ) && UNEXPECTED ( Z_OBJ_HANDLER_P ( op1 , do_operation ) ) && EXPECTED ( SUCCESS == Z_OBJ_HANDLER_P ( op1 , do_operation ) ( opcode , result , op1 , NULL ) ) ) { return SUCCESS ; }

// # include < ctype . h >

// # include "zend.h"

// # include "zend_operators.h"

// # include "zend_variables.h"

// # include "zend_globals.h"

// # include "zend_list.h"

// # include "zend_API.h"

// # include "zend_strtod.h"

// # include "zend_exceptions.h"

// # include "zend_closures.h"

// #define convert_op1_op2_long(op1,op1_lval,op2,op2_lval,result,op,op_func) do { if ( UNEXPECTED ( Z_TYPE_P ( op1 ) != IS_LONG ) ) { if ( Z_ISREF_P ( op1 ) ) { op1 = Z_REFVAL_P ( op1 ) ; if ( Z_TYPE_P ( op1 ) == IS_LONG ) { op1_lval = Z_LVAL_P ( op1 ) ; break ; } } ZEND_TRY_BINARY_OP1_OBJECT_OPERATION ( op , op_func ) ; op1_lval = _zval_get_long_func_noisy ( op1 ) ; if ( UNEXPECTED ( EG ( exception ) ) ) { if ( result != op1 ) { ZVAL_UNDEF ( result ) ; } return FAILURE ; } } else { op1_lval = Z_LVAL_P ( op1 ) ; } } while ( 0 ) ; do { if ( UNEXPECTED ( Z_TYPE_P ( op2 ) != IS_LONG ) ) { if ( Z_ISREF_P ( op2 ) ) { op2 = Z_REFVAL_P ( op2 ) ; if ( Z_TYPE_P ( op2 ) == IS_LONG ) { op2_lval = Z_LVAL_P ( op2 ) ; break ; } } ZEND_TRY_BINARY_OP2_OBJECT_OPERATION ( op ) ; op2_lval = _zval_get_long_func_noisy ( op2 ) ; if ( UNEXPECTED ( EG ( exception ) ) ) { if ( result != op1 ) { ZVAL_UNDEF ( result ) ; } return FAILURE ; } } else { op2_lval = Z_LVAL_P ( op2 ) ; } } while ( 0 ) ;
