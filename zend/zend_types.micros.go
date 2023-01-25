// <<generate>>

package zend

// #define ZEND_TYPES_H

// # include "zend_portability.h"

// # include "zend_long.h"

// #define ZEND_ENDIAN_LOHI(lo,hi) lo ; hi ;

// #define ZEND_ENDIAN_LOHI_3(lo,mi,hi) lo ; mi ; hi ;

// #define ZEND_ENDIAN_LOHI_4(a,b,c,d) a ; b ; c ; d ;

// #define ZEND_ENDIAN_LOHI_C(lo,hi) lo , hi

// #define ZEND_ENDIAN_LOHI_C_3(lo,mi,hi) lo , mi , hi ,

// #define ZEND_ENDIAN_LOHI_C_4(a,b,c,d) a , b , c , d

// #define ZEND_TLS       static

// #define ZEND_EXT_TLS

// #define ZEND_TYPE_ENCODE_CLASS_CONST_Q1(allow_null,class_name) ZEND_TYPE_ENCODE_CLASS_CONST_Q2 ( ZEND_TYPE_ENCODE_CLASS_CONST_ ## allow_null , class_name )
