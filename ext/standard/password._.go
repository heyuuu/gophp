package standard

import (
	"sik/zend/types"
)

var PhpPasswordAlgos types.Array

var PhpPasswordAlgoBcrypt PhpPasswordAlgo = MakePhpPasswordAlgo("bcrypt", PhpPasswordBcryptHash, PhpPasswordBcryptVerify, PhpPasswordBcryptNeedsRehash, PhpPasswordBcryptGetInfo, PhpPasswordBcryptValid)
