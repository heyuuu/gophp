package standard

import (
	"github.com/heyuuu/gophp/php/types"
)

const PHP_PASSWORD_BCRYPT_COST = 10

var PhpPasswordAlgos *types.Array
var passwordAlgos map[string]*PhpPasswordAlgo

var passwordAlgoBcrypt *PhpPasswordAlgo = NewPhpPasswordAlgo(
	"bcrypt",
	PhpPasswordBcryptHash,
	PhpPasswordBcryptVerify,
	PhpPasswordBcryptNeedsRehash,
	PhpPasswordBcryptGetInfo,
	PhpPasswordBcryptValid,
)

/**
 * PhpPasswordAlgo
 */
type PhpPasswordAlgo struct {
	name       string
	hash       func(password string, options *types.Array) (string, bool)
	verify     func(password string, hash string) bool
	needRehash func(password string, options *types.Array) bool
	getInfo    func(hash string) *types.Array
	valid      func(hash string) bool
}

func NewPhpPasswordAlgo(
	name string,
	hash func(password string, options *types.Array) (string, bool),
	verify func(password string, hash string) bool,
	needRehash func(password string, options *types.Array) bool,
	getInfo func(hash string) *types.Array,
	valid func(hash string) bool,
) *PhpPasswordAlgo {
	return &PhpPasswordAlgo{name: name, hash: hash, verify: verify, needRehash: needRehash, getInfo: getInfo, valid: valid}
}

func (algo *PhpPasswordAlgo) Name() string { return algo.name }
func (algo *PhpPasswordAlgo) Hash(password string, options *types.Array) (string, bool) {
	return algo.hash(password, options)
}
func (algo *PhpPasswordAlgo) Verify(password string, hash string) bool {
	return algo.verify(password, hash)
}
func (algo *PhpPasswordAlgo) NeedsRehash(password string, options *types.Array) bool {
	return algo.needRehash(password, options)
}
func (algo *PhpPasswordAlgo) GetInfo(hash string) *types.Array {
	return algo.getInfo(hash)
}
func (algo *PhpPasswordAlgo) Valid(hash string) bool {
	return algo.valid(hash)
}
