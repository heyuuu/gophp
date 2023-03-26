package standard

/* sha512 crypt has the maximal salt length of 123 characters */

const PHP_MAX_SALT_LEN = 123

/* Used to check DES salts to ensure that they contain only valid characters */

const DES_INVALID_SALT_ERROR = "Supplied salt is not valid for DES. Possible bug in provided salt format."

const Itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
