// <<generate>>

package zend

// #define ZEND_PTR_STACK_H

// #define ZEND_PTR_STACK_RESIZE_IF_NEEDED(stack,count) if ( stack -> top + count > stack -> max ) { do { stack -> max += PTR_STACK_BLOCK_SIZE ; } while ( stack -> top + count > stack -> max ) ; stack -> elements = ( void * * ) perealloc ( stack -> elements , ( sizeof ( void * ) * ( stack -> max ) ) , stack -> persistent ) ; stack -> top_element = stack -> elements + stack -> top ; }

// # include "zend.h"

// # include "zend_ptr_stack.h"

// # include < stdarg . h >
