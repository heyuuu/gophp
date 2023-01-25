// <<generate>>

package standard

// Source: <ext/standard/crypt_sha256.c>

/* SHA256-based Unix crypt implementation.
   Released into the Public Domain by Ulrich Drepper <drepper@redhat.com>.  */

/* Structure to save state of computation between the single steps.  */

/* This array contains the bytes used to pad the buffer to the next
   64-byte boundary.  (FIPS 180-2:5.1.1)  */

var Fillbuf []uint8 = []uint8{0x80, 0}

/* Constants for SHA256 from FIPS 180-2:4.2.2.  */

var K32 []uint32 = []uint32{0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5, 0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174, 0xe49b69c1, 0xefbe4786, 0xfc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da, 0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x6ca6351, 0x14292967, 0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85, 0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070, 0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3, 0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2}

/* Process LEN bytes of BUFFER, accumulating context into CTX.
   It is assumed that LEN % 64 == 0.  */

/* Initialize structure containing state of computation.
   (FIPS 180-2:5.3.2)  */

/* Process the remaining bytes in the internal buffer and the usual
   prolog according to the standard and write the result to RESBUF.

   IMPORTANT: On some systems it is required that RESBUF is correctly
   aligned for a 32 bits value.  */

/* Define our magic string to mark salt for SHA256 "encryption"
   replacement.  */

var Sha256SaltPrefix []byte = "$5$"

/* Prefix for optional rounds specification.  */

var Sha256RoundsPrefix []byte = "rounds="

/* Maximum salt string length.  */

const SALT_LEN_MAX = 16

/* Default number of rounds if not explicitly specified.  */

const ROUNDS_DEFAULT = 5000

/* Minimum number of rounds.  */

const ROUNDS_MIN = 1000

/* Maximum number of rounds.  */

const ROUNDS_MAX = 999999999

/* Table with characters for base64 transformation.  */

var B64t []byte = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

/* This entry point is equivalent to the `crypt' function in Unix
   libcs.  */
