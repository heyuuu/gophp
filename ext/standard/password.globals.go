// <<generate>>

package standard

import (
	"sik/zend"
)

var PhpPasswordAlgos zend.ZendArray
var PhpPasswordAlgoBcrypt PhpPasswordAlgo = PhpPasswordAlgo{"bcrypt", PhpPasswordBcryptHash, PhpPasswordBcryptVerify, PhpPasswordBcryptNeedsRehash, PhpPasswordBcryptGetInfo, PhpPasswordBcryptValid}
