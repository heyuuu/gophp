// <<generate>>

package cli

// #define php_http_parser_h

// # include < sys / types . h >

// # include "php_config.h"

// # include "php_stdint.h"

// # include < assert . h >

// # include < stddef . h >

// # include "php_http_parser.h"

// #define CALLBACK2(FOR) do { if ( settings -> on_ ## FOR ) { if ( 0 != settings -> on_ ## FOR ( parser ) ) return ( p - data ) ; } } while ( 0 )

// #define MARK(FOR) do { FOR ## _mark = p ; } while ( 0 )

// #define CALLBACK_NOCLEAR(FOR) do { if ( FOR ## _mark ) { if ( settings -> on_ ## FOR ) { if ( 0 != settings -> on_ ## FOR ( parser , FOR ## _mark , p - FOR ## _mark ) ) { return ( p - data ) ; } } } } while ( 0 )

// #define CALLBACK(FOR) do { CALLBACK_NOCLEAR ( FOR ) ; FOR ## _mark = NULL ; } while ( 0 )

// #define STRICT_CHECK(cond)
