// <<generate>>

package standard

// Source: <ext/standard/crypt_freesec.h>

/*
 * _crypt_extended_init() must be called explicitly before first use of
 * _crypt_extended_r().
 */

// Source: <ext/standard/crypt_freesec.c>

/*
 * This version is derived from the original implementation of FreeSec
 * (release 1.1) by David Burren.  I've reviewed the changes made in
 * OpenBSD (as of 2.7) and modified the original code in a similar way
 * where applicable.  I've also made it reentrant and made a number of
 * other changes.
 * - Solar Designer <solar at openwall.com>
 */

/*
 *    No E box is used, as it's replaced by some ANDs, shifts, and ORs.
 */

/*      0000000000111111111122222222223333333333444444444455555555556666 */

/*
 * When we choose to "support" invalid salts, nevertheless disallow those
 * containing characters that would violate the passwd file format.
 */
