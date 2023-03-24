package standard

// Source: <ext/standard/crypt_sha512.c>

/* SHA512-based Unix crypt implementation.
   Released into the Public Domain by Ulrich Drepper <drepper@redhat.com>.  */

/* See #51582 */

/* Structure to save state of computation between the single steps.  */

/* This array contains the bytes used to pad the buffer to the next
   64-byte boundary.  (FIPS 180-2:5.1.2)  */

/* Constants for SHA512 from FIPS 180-2:4.2.3.  */

var K64 []uint64 = []uint64{uint64(0x428a2f98d728ae22), uint64(0x7137449123ef65cd), uint64(-0x4a3f043013b2c4d1), uint64(-0x164a245a7e762444), uint64(0x3956c25bf348b538), uint64(0x59f111f1b605d019), uint64(-0x6dc07d5b50e6b065), uint64(-0x54e3a12a25927ee8), uint64(-0x27f855675cfcfdbe), uint64(0x12835b0145706fbe), uint64(0x243185be4ee4b28c), uint64(0x550c7dc3d5ffb4e2), uint64(0x72be5d74f27b896f), uint64(-0x7f214e01c4e9694f), uint64(-0x6423f958da38edcb), uint64(-0x3e640e8b3096d96c), uint64(-0x1b64963e610eb52e), uint64(-0x1041b879c7b0da1d), uint64(0xfc19dc68b8cd5b5), uint64(0x240ca1cc77ac9c65), uint64(0x2de92c6f592b0275), uint64(0x4a7484aa6ea6e483), uint64(0x5cb0a9dcbd41fbd4), uint64(0x76f988da831153b5), uint64(-0x67c1aead11992055), uint64(-0x57ce3992d24bcdf0), uint64(-0x4ffcd8376704dec1), uint64(-0x40a680384110f11c), uint64(-0x391ff40cc257703e), uint64(-0x2a586eb86cf558db), uint64(0x6ca6351e003826f), uint64(0x142929670a0e6e70), uint64(0x27b70a8546d22ffc), uint64(0x2e1b21385c26c926), uint64(0x4d2c6dfc5ac42aed), uint64(0x53380d139d95b3df), uint64(0x650a73548baf63de), uint64(0x766a0abb3c77b2a8), uint64(-0x7e3d36d1b812511a), uint64(-0x6d8dd37aeb7dcac5), uint64(-0x5d40175eb30efc9c), uint64(-0x57e599b443bdcfff), uint64(-0x3db4748f2f07686f), uint64(-0x3893ae5cf9ab41d0), uint64(-0x2e6d17e62910ade8), uint64(-0x2966f9dbaa9a56f0), uint64(-0xbf1ca7aa88edfd6), uint64(0x106aa07032bbd1b8), uint64(0x19a4c116b8d2d0c8), uint64(0x1e376c085141ab53), uint64(0x2748774cdf8eeb99), uint64(0x34b0bcb5e19b48a8), uint64(0x391c0cb3c5c95a63), uint64(0x4ed8aa4ae3418acb), uint64(0x5b9cca4f7763e373), uint64(0x682e6ff3d6b2b8a3), uint64(0x748f82ee5defb2fc), uint64(0x78a5636f43172f60), uint64(-0x7b3787eb5e0f548e), uint64(-0x7338fdf7e59bc614), uint64(-0x6f410005dc9ce1d8), uint64(-0x5baf9314217d4217), uint64(-0x41065c084d3986eb), uint64(-0x398e870d1c8dacd5), uint64(-0x35d8c13115d99e64), uint64(-0x2e794738de3f3df9), uint64(-0x15258229321f14e2), uint64(-0xa82b08011912e88), uint64(0x6f067aa72176fba), uint64(0xa637dc5a2c898a6), uint64(0x113f9804bef90dae), uint64(0x1b710b35131c471b), uint64(0x28db77f523047d84), uint64(0x32caab7b40c72493), uint64(0x3c9ebe0a15c9bebc), uint64(0x431d67c49c100d4c), uint64(0x4cc5d4becb3e42b6), uint64(0x597f299cfc657e2a), uint64(0x5fcb6fab3ad6faec), uint64(0x6c44198c4a475817)}

/* Process LEN bytes of BUFFER, accumulating context into CTX.
   It is assumed that LEN % 128 == 0.  */

/* Initialize structure containing state of computation.
   (FIPS 180-2:5.3.3)  */

/* Process the remaining bytes in the internal buffer and the usual
   prolog according to the standard and write the result to RESBUF.

   IMPORTANT: On some systems it is required that RESBUF is correctly
   aligned for a 32 bits value. */

/* Define our magic string to mark salt for SHA512 "encryption"
   replacement.  */

var Sha512SaltPrefix []byte = "$6$"

/* Prefix for optional rounds specification.  */

var Sha512RoundsPrefix []byte = "rounds="

/* Maximum salt string length.  */

/* Default number of rounds if not explicitly specified.  */

/* Minimum number of rounds.  */

/* Maximum number of rounds.  */

/* Table with characters for base64 transformation.  */

/* This entry point is equivalent to the `crypt' function in Unix
   libcs.  */
