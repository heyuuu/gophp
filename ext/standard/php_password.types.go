package standard

import (
	"github.com/heyuuu/gophp/zend/types"
)

/**
 * PhpPasswordAlgo
 */
type PhpPasswordAlgo struct {
	name         *byte
	hash         func(password *types.String, options *types.Array) *types.String
	verify       func(password *types.String, hash *types.String) types.ZendBool
	needs_rehash func(password *types.String, options *types.Array) types.ZendBool
	get_info     func(return_value *types.Zval, hash *types.String) int
	valid        func(hash *types.String) types.ZendBool
}

func MakePhpPasswordAlgo(
	name *byte,
	hash func(password *types.String, options *types.Array) *types.String,
	verify func(password *types.String, hash *types.String) types.ZendBool,
	needs_rehash func(password *types.String, options *types.Array) types.ZendBool,
	get_info func(return_value *types.Zval, hash *types.String) int,
	valid func(hash *types.String) types.ZendBool,
) PhpPasswordAlgo {
	return PhpPasswordAlgo{
		name:         name,
		hash:         hash,
		verify:       verify,
		needs_rehash: needs_rehash,
		get_info:     get_info,
		valid:        valid,
	}
}
func (this *PhpPasswordAlgo) GetName() *byte { return this.name }

// func (this *PhpPasswordAlgo) SetName(value *byte) { this.name = value }
func (this *PhpPasswordAlgo) GetHash() func(password *types.String, options *types.Array) *types.String {
	return this.hash
}

// func (this *PhpPasswordAlgo) SetHash(value func(password *zend.String, options *zend.Array) *zend.String) { this.hash = value }
func (this *PhpPasswordAlgo) GetVerify() func(password *types.String, hash *types.String) types.ZendBool {
	return this.verify
}

// func (this *PhpPasswordAlgo) SetVerify(value func(password *zend.String, hash *zend.String) zend.ZendBool) { this.verify = value }
func (this *PhpPasswordAlgo) GetNeedsRehash() func(password *types.String, options *types.Array) types.ZendBool {
	return this.needs_rehash
}

// func (this *PhpPasswordAlgo) SetNeedsRehash(value func(password *zend.String, options *zend.Array) zend.ZendBool) { this.needs_rehash = value }
func (this *PhpPasswordAlgo) GetGetInfo() func(return_value *types.Zval, hash *types.String) int {
	return this.get_info
}

// func (this *PhpPasswordAlgo) SetGetInfo(value func(return_value *zend.Zval, hash *zend.String) int) { this.get_info = value }
func (this *PhpPasswordAlgo) GetValid() func(hash *types.String) types.ZendBool {
	return this.valid
}

// func (this *PhpPasswordAlgo) SetValid(value func(hash *zend.String) zend.ZendBool) { this.valid = value }
