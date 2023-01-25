// <<generate>>

package standard

// # include "php.h"

// # include "php_metaphone.h"

// #define End_Phoned_Word       { if ( p_idx == max_buffer_len ) { * phoned_word = zend_string_extend ( * phoned_word , 1 * sizeof ( char ) + max_buffer_len , 0 ) ; max_buffer_len += 1 ; } ZSTR_VAL ( * phoned_word ) [ p_idx ] = '\0' ; ZSTR_LEN ( * phoned_word ) = p_idx ; }
