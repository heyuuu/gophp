// <<generate>>

package standard

// # include "php.h"

// # include "php_globals.h"

// # include "ext/standard/basic_functions.h"

// # include "ext/standard/file.h"

// # include "ext/standard/php_string.h"

// # include "zend_smart_str.h"

// #define bmask(a) ( 0xffff >> ( 16 - a ) )

// #define NEXT_CHAR(ps,icnt,lb_ptr,lb_cnt,lbchars) ( ( lb_ptr ) < ( lb_cnt ) ? ( lbchars ) [ ( lb_ptr ) ] : * ( ps ) )

// #define CONSUME_CHAR(ps,icnt,lb_ptr,lb_cnt) if ( ( lb_ptr ) < ( lb_cnt ) ) { ( lb_ptr ) ++ ; } else { ( lb_cnt ) = ( lb_ptr ) = 0 ; -- ( icnt ) ; ( ps ) ++ ; }

// #define GET_STR_PROP(ht,var,var_len,fldname,persistent) php_conv_get_string_prop_ex ( ht , & var , & var_len , fldname , sizeof ( fldname ) , persistent )

// #define GET_INT_PROP(ht,var,fldname) php_conv_get_int_prop_ex ( ht , & var , fldname , sizeof ( fldname ) )

// #define GET_UINT_PROP(ht,var,fldname) php_conv_get_uint_prop_ex ( ht , & var , fldname , sizeof ( fldname ) )

// #define GET_BOOL_PROP(ht,var,fldname) php_conv_get_bool_prop_ex ( ht , & var , fldname , sizeof ( fldname ) )
