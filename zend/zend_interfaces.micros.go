// <<generate>>

package zend

// #define ZEND_INTERFACES_H

// # include "zend.h"

// # include "zend_API.h"

// #define REGISTER_MAGIC_INTERFACE(class_name,class_name_str) { zend_class_entry ce ; INIT_CLASS_ENTRY ( ce , # class_name_str , zend_funcs_ ## class_name ) zend_ce_ ## class_name = zend_register_internal_interface ( & ce ) ; zend_ce_ ## class_name -> interface_gets_implemented = zend_implement_ ## class_name ; }

// #define REGISTER_MAGIC_IMPLEMENT(class_name,interface_name) zend_class_implements ( zend_ce_ ## class_name , 1 , zend_ce_ ## interface_name )

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_interfaces.h"

// # include "zend_exceptions.h"
