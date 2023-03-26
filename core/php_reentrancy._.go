package core

/* currently, PHP does not check for these functions, but assumes
   that they are available on all systems. */

const HAVE_LOCALTIME = 1
const HAVE_GMTIME = 1
const HAVE_ASCTIME = 1
const HAVE_CTIME = 1
const PhpLocaltimeR = localtime_r
const PhpCtimeR = ctime_r
const PhpAsctimeR = asctime_r
const PhpGmtimeR = gmtime_r
const PhpStrtokR = strtok_r
const PhpRandR = rand_r
