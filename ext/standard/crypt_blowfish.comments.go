// <<generate>>

package standard

// Source: <ext/standard/crypt_blowfish.h>

/*
 * Written by Solar Designer <solar at openwall.com> in 2000-2011.
 * No copyright is claimed, and the software is hereby placed in the public
 * domain. In case this attempt to disclaim copyright and place the software
 * in the public domain is deemed null and void, then the software is
 * Copyright (c) 2000-2011 Solar Designer and it is hereby released to the
 * general public under the following terms:
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted.
 *
 * There's ABSOLUTELY NO WARRANTY, express or implied.
 *
 * See crypt_blowfish.c for more information.
 */

// Source: <ext/standard/crypt_blowfish.c>

/*
 * The crypt_blowfish homepage is:
 *
 *    http://www.openwall.com/crypt/
 *
 * This code comes from John the Ripper password cracker, with reentrant
 * and crypt(3) interfaces added, but optimizations specific to password
 * cracking removed.
 *
 * Written by Solar Designer <solar at openwall.com> in 1998-2015.
 * No copyright is claimed, and the software is hereby placed in the public
 * domain. In case this attempt to disclaim copyright and place the software
 * in the public domain is deemed null and void, then the software is
 * Copyright (c) 1998-2015 Solar Designer and it is hereby released to the
 * general public under the following terms:
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted.
 *
 * There's ABSOLUTELY NO WARRANTY, express or implied.
 *
 * It is my intent that you should be able to use this on your system,
 * as part of a software package, or anywhere else to improve security,
 * ensure compatibility, or for any other purpose. I would appreciate
 * it if you give credit where it is due and keep your modifications in
 * the public domain as well, but I don't require that in order to let
 * you place this code and any modifications you make under a license
 * of your choice.
 *
 * This implementation is fully compatible with OpenBSD's bcrypt.c for prefix
 * "$2b$", originally by Niels Provos <provos at citi.umich.edu>, and it uses
 * some of his ideas. The password hashing algorithm was designed by David
 * Mazieres <dm at lcs.mit.edu>. For information on the level of
 * compatibility for bcrypt hash prefixes other than "$2b$", please refer to
 * the comments in BF_set_key() below and to the included crypt(3) man page.
 *
 * There's a paper on the algorithm that explains its design decisions:
 *
 *    http://www.usenix.org/events/usenix99/provos.html
 *
 * Some of the tricks in BF_ROUND might be inspired by Eric Young's
 * Blowfish library (I can't be sure if I would think of something if I
 * hadn't seen his code).
 */

/* Just to make sure the prototypes match the actual definitions */

/* Number of Blowfish rounds, this is also hardcoded into a few places */

/*
 * Magic IV for 64 Blowfish encryptions that we do at the end.
 * The string is "OrpheanBeholderScryDoubt" on big-endian.
 */

/*
 * P-box and S-box tables initialized with digits of Pi.
 */

/* Architectures with no complicated addressing modes supported */

/*
 * Encrypt one block, BF_N is hardcoded here.
 */

/*
 * Please preserve the runtime self-test. It serves two purposes at once:
 *
 * 1. We really can't afford the risk of producing incompatible hashes e.g.
 * when there's something like gcc bug 26587 again, whereas an application or
 * library integrating this code might not also integrate our external tests or
 * it might not run them after every build. Even if it does, the miscompile
 * might only occur on the production build, but not on a testing build (such
 * as because of different optimization settings). It is painful to recover
 * from incorrectly-computed hashes - merely fixing whatever broke is not
 * enough. Thus, a proactive measure like this self-test is needed.
 *
 * 2. We don't want to leave sensitive data from our actual password hash
 * computation on the stack or in registers. Previous revisions of the code
 * would do explicit cleanups, but simply running the self-test after hash
 * computation is more reliable.
 *
 * The performance cost of this quick self-test is around 0.6% at the "$2a$08"
 * setting.
 */
