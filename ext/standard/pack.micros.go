package standard

// #define PACK_H

// # include "php.h"

// # include < stdio . h >

// # include < stdlib . h >

// # include < errno . h >

// # include < sys / types . h >

// # include < sys / stat . h >

// # include < fcntl . h >

// # include < sys / param . h >

// # include "ext/standard/head.h"

// # include "php_string.h"

// # include "pack.h"

// # include < pwd . h >

// # include "fsock.h"

// # include < netinet / in . h >

// #define INC_OUTPUTPOS(a,b) if ( ( a ) < 0 || ( ( INT_MAX - outputpos ) / ( ( int ) b ) ) < ( a ) ) { efree ( formatcodes ) ; efree ( formatargs ) ; php_error_docref ( NULL , E_WARNING , "Type %c: integer overflow in format string" , code ) ; RETURN_FALSE ; } outputpos += ( a ) * ( b ) ;
