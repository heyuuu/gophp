// <<generate>>

package zend

// #define ZEND_SCANNER_H

// #define YYDEBUG(s,c)

// # include "zend_language_scanner_defs.h"

// # include < errno . h >

// # include "zend.h"

// # include "zend_alloc.h"

// # include < zend_language_parser . h >

// # include "zend_compile.h"

// # include "zend_language_scanner.h"

// # include "zend_highlight.h"

// # include "zend_constants.h"

// # include "zend_variables.h"

// # include "zend_operators.h"

// # include "zend_API.h"

// # include "zend_strtod.h"

// # include "zend_exceptions.h"

// # include "zend_virtual_cwd.h"

// #define YYCTYPE       unsigned char

// #define YYFILL(n) { if ( ( YYCURSOR + n ) >= ( YYLIMIT + ZEND_MMAP_AHEAD ) ) { return 0 ; } }

// #define STATE(name) yyc ## name

// # include < stdarg . h >

// # include < unistd . h >

// #define yy_push_state(state_and_tsrm) _yy_push_state ( yyc ## state_and_tsrm )

// #define RETURN_TOKEN(_token) do { token = _token ; goto emit_token ; } while ( 0 )
