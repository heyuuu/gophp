// <<generate>>

package standard

// #define _CRYPT_BLOWFISH_H

// # include < string . h >

// # include < errno . h >

// # include "crypt_blowfish.h"

// #define BF_safe_atoi64(dst,src) { tmp = ( unsigned char ) ( src ) ; if ( tmp == '$' ) break ; if ( ( unsigned int ) ( tmp -= 0x20 ) >= 0x60 ) return - 1 ; tmp = BF_atoi64 [ tmp ] ; if ( tmp > 63 ) return - 1 ; ( dst ) = tmp ; }

// #define BF_ENCRYPT       L ^= data . ctx . P [ 0 ] ; BF_ROUND ( L , R , 0 ) ; BF_ROUND ( R , L , 1 ) ; BF_ROUND ( L , R , 2 ) ; BF_ROUND ( R , L , 3 ) ; BF_ROUND ( L , R , 4 ) ; BF_ROUND ( R , L , 5 ) ; BF_ROUND ( L , R , 6 ) ; BF_ROUND ( R , L , 7 ) ; BF_ROUND ( L , R , 8 ) ; BF_ROUND ( R , L , 9 ) ; BF_ROUND ( L , R , 10 ) ; BF_ROUND ( R , L , 11 ) ; BF_ROUND ( L , R , 12 ) ; BF_ROUND ( R , L , 13 ) ; BF_ROUND ( L , R , 14 ) ; BF_ROUND ( R , L , 15 ) ; tmp4 = R ; R = L ; L = tmp4 ^ data . ctx . P [ BF_N + 1 ] ;
