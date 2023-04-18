package standard

import (
	"github.com/heyuuu/gophp/php/types"
)

var PhpPasswordAlgos *types.Array

var PhpPasswordAlgoBcrypt PhpPasswordAlgo = MakePhpPasswordAlgo("bcrypt", PhpPasswordBcryptHash, PhpPasswordBcryptVerify, PhpPasswordBcryptNeedsRehash, PhpPasswordBcryptGetInfo, PhpPasswordBcryptValid)
