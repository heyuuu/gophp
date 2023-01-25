// <<generate>>

package standard

// Source: <ext/standard/crypt_sha256.c>

/* SHA256-based Unix crypt implementation.
   Released into the Public Domain by Ulrich Drepper <drepper@redhat.com>.  */

/* Structure to save state of computation between the single steps.  */

/* This array contains the bytes used to pad the buffer to the next
   64-byte boundary.  (FIPS 180-2:5.1.1)  */

/* Constants for SHA256 from FIPS 180-2:4.2.2.  */

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

/* Prefix for optional rounds specification.  */

/* Maximum salt string length.  */

/* Default number of rounds if not explicitly specified.  */

/* Minimum number of rounds.  */

/* Maximum number of rounds.  */

/* Table with characters for base64 transformation.  */

/* This entry point is equivalent to the `crypt' function in Unix
   libcs.  */
