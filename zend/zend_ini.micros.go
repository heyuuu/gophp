// <<generate>>

package zend

// #define ZEND_INI_H

// #define ZEND_INI_MH(name) int name ( zend_ini_entry * entry , zend_string * new_value , void * mh_arg1 , void * mh_arg2 , void * mh_arg3 , int stage )

// #define ZEND_INI_DISP(name) ZEND_COLD void name ( zend_ini_entry * ini_entry , int type )

// #define ZEND_INI_BEGIN() static const zend_ini_entry_def ini_entries [ ] = {

// #define ZEND_INI_END() { NULL , NULL , NULL , NULL , NULL , NULL , NULL , 0 , 0 , 0 } } ;

// #define ZEND_INI_ENTRY3_EX(name,default_value,modifiable,on_modify,arg1,arg2,arg3,displayer) { name , on_modify , arg1 , arg2 , arg3 , default_value , displayer , sizeof ( default_value ) - 1 , sizeof ( name ) - 1 , modifiable } ,

// #define ZEND_INI_ENTRY3(name,default_value,modifiable,on_modify,arg1,arg2,arg3) ZEND_INI_ENTRY3_EX ( name , default_value , modifiable , on_modify , arg1 , arg2 , arg3 , NULL )

// #define ZEND_INI_ENTRY2_EX(name,default_value,modifiable,on_modify,arg1,arg2,displayer) ZEND_INI_ENTRY3_EX ( name , default_value , modifiable , on_modify , arg1 , arg2 , NULL , displayer )

// #define ZEND_INI_ENTRY2(name,default_value,modifiable,on_modify,arg1,arg2) ZEND_INI_ENTRY2_EX ( name , default_value , modifiable , on_modify , arg1 , arg2 , NULL )

// #define ZEND_INI_ENTRY1_EX(name,default_value,modifiable,on_modify,arg1,displayer) ZEND_INI_ENTRY3_EX ( name , default_value , modifiable , on_modify , arg1 , NULL , NULL , displayer )

// #define ZEND_INI_ENTRY1(name,default_value,modifiable,on_modify,arg1) ZEND_INI_ENTRY1_EX ( name , default_value , modifiable , on_modify , arg1 , NULL )

// #define ZEND_INI_ENTRY_EX(name,default_value,modifiable,on_modify,displayer) ZEND_INI_ENTRY3_EX ( name , default_value , modifiable , on_modify , NULL , NULL , NULL , displayer )

// #define ZEND_INI_ENTRY(name,default_value,modifiable,on_modify) ZEND_INI_ENTRY_EX ( name , default_value , modifiable , on_modify , NULL )

// #define STD_ZEND_INI_ENTRY(name,default_value,modifiable,on_modify,property_name,struct_type,struct_ptr) ZEND_INI_ENTRY2 ( name , default_value , modifiable , on_modify , ( void * ) XtOffsetOf ( struct_type , property_name ) , ( void * ) & struct_ptr )

// #define STD_ZEND_INI_ENTRY_EX(name,default_value,modifiable,on_modify,property_name,struct_type,struct_ptr,displayer) ZEND_INI_ENTRY2_EX ( name , default_value , modifiable , on_modify , ( void * ) XtOffsetOf ( struct_type , property_name ) , ( void * ) & struct_ptr , displayer )

// #define STD_ZEND_INI_BOOLEAN(name,default_value,modifiable,on_modify,property_name,struct_type,struct_ptr) ZEND_INI_ENTRY3_EX ( name , default_value , modifiable , on_modify , ( void * ) XtOffsetOf ( struct_type , property_name ) , ( void * ) & struct_ptr , NULL , zend_ini_boolean_displayer_cb )

// # include "zend.h"

// # include "zend_sort.h"

// # include "zend_API.h"

// # include "zend_ini.h"

// # include "zend_alloc.h"

// # include "zend_operators.h"

// # include "zend_strtod.h"
