// <<generate>>

package zend

// #define ZEND_EXECUTE_H

// # include "zend_compile.h"

// # include "zend_hash.h"

// # include "zend_operators.h"

// # include "zend_variables.h"

// #define ZEND_ASSERT_VM_STACK(stack)

// #define ZEND_ASSERT_VM_STACK_GLOBAL

// #define ZEND_REF_FOREACH_TYPE_SOURCES(ref,prop) do { zend_property_info_source_list * _source_list = & ZEND_REF_TYPE_SOURCES ( ref ) ; zend_property_info * * _prop , * * _end ; zend_property_info_list * _list ; if ( _source_list -> ptr ) { if ( ZEND_PROPERTY_INFO_SOURCE_IS_LIST ( _source_list -> list ) ) { _list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST ( _source_list -> list ) ; _prop = _list -> ptr ; _end = _list -> ptr + _list -> num ; } else { _prop = & _source_list -> ptr ; _end = _prop + 1 ; } for ( ; _prop < _end ; _prop ++ ) { prop = * _prop ;

// #define ZEND_REF_FOREACH_TYPE_SOURCES_END() } } } while ( 0 )

// # include < stdio . h >

// # include < signal . h >

// # include "zend.h"

// # include "zend_compile.h"

// # include "zend_execute.h"

// # include "zend_API.h"

// # include "zend_ptr_stack.h"

// # include "zend_constants.h"

// # include "zend_extensions.h"

// # include "zend_ini.h"

// # include "zend_exceptions.h"

// # include "zend_interfaces.h"

// # include "zend_closures.h"

// # include "zend_generators.h"

// # include "zend_vm.h"

// # include "zend_dtrace.h"

// # include "zend_inheritance.h"

// # include "zend_type_info.h"

// # include "zend_virtual_cwd.h"

// #define EXECUTE_DATA_DC       , EXECUTE_DATA_D

// #define EXECUTE_DATA_CC       , EXECUTE_DATA_C

// #define NO_EXECUTE_DATA_CC       , NULL

// #define OPLINE_D       const zend_op * opline

// #define OPLINE_DC       , OPLINE_D

// #define OPLINE_CC       , OPLINE_C

// #define CHECK_SYMBOL_TABLES()

// #define ZEND_VM_NEXT_OPCODE_EX(check_exception,skip) CHECK_SYMBOL_TABLES ( ) if ( check_exception ) { OPLINE = EX ( opline ) + ( skip ) ; } else { ZEND_ASSERT ( ! EG ( exception ) ) ; OPLINE = opline + ( skip ) ; } ZEND_VM_CONTINUE ( )

// #define ZEND_VM_SET_NEXT_OPCODE(new_op) CHECK_SYMBOL_TABLES ( ) OPLINE = new_op

// #define ZEND_VM_SET_OPCODE(new_op) CHECK_SYMBOL_TABLES ( ) OPLINE = new_op ; ZEND_VM_INTERRUPT_CHECK ( )

// #define ZEND_VM_REPEATABLE_OPCODE       do {

// #define ZEND_VM_REPEAT_OPCODE(_opcode) } while ( UNEXPECTED ( ( ++ opline ) -> opcode == _opcode ) ) ; OPLINE = opline ; ZEND_VM_CONTINUE ( )

// #define ZEND_VM_GUARD(name)

// # include "zend_vm_execute.h"
