// <<generate>>

package zend

// #define _ZEND_INI_SCANNER_H

// # include < errno . h >

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_globals.h"

// # include < zend_ini_parser . h >

// # include "zend_ini_scanner.h"

// #define YYDEBUG(s,c)

// # include "zend_ini_scanner_defs.h"

// #define YYCTYPE       unsigned char

// #define YYFILL(n) { if ( YYCURSOR > YYLIMIT ) return 0 ; }

// #define STATE(name) yyc ## name

// #define RETURN_TOKEN(type,str,len) { if ( SCNG ( scanner_mode ) == ZEND_INI_SCANNER_TYPED && ( YYSTATE == STATE ( ST_VALUE ) || YYSTATE == STATE ( ST_RAW ) ) ) { zend_ini_copy_typed_value ( ini_lval , type , str , len ) ; } else { zend_ini_copy_value ( ini_lval , str , len ) ; } return type ; }

// #define yy_push_state(state_and_tsrm) _yy_push_state ( yyc ## state_and_tsrm )
