package standard

/* php_flock internally uses fcntl whether or not flock is available
 * This way our php_flock even works on NFS files.
 * More info: /usr/src/linux/Documentation
 */

/* Userland LOCK_* constants */

const PHP_LOCK_SH = 1
const PHP_LOCK_EX = 2
const PHP_LOCK_UN = 3
const PHP_LOCK_NB = 4
