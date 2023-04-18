package standard

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

/**
 * PhpPasswordAlgo
 */
type PhpPasswordAlgo struct {
	name         *byte
	hash         func(password *types2.String, options *types2.Array) *types2.String
	verify       func(password *types2.String, hash *types2.String) types2.ZendBool
	needs_rehash func(password *types2.String, options *types2.Array) types2.ZendBool
	get_info     func(return_value *types2.Zval, hash *types2.String) int
	valid        func(hash *types2.String) types2.ZendBool
}

func MakePhpPasswordAlgo(
	name *byte,
	hash func(password *types2.String, options *types2.Array) *types2.String,
	verify func(password *types2.String, hash *types2.String) types2.ZendBool,
	needs_rehash func(password *types2.String, options *types2.Array) types2.ZendBool,
	get_info func(return_value *types2.Zval, hash *types2.String) int,
	valid func(hash *types2.String) types2.ZendBool,
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
func (this *PhpPasswordAlgo) GetHash() func(password *types2.String, options *types2.Array) *types2.String {
	return this.hash
}

// func (this *PhpPasswordAlgo) SetHash(value func(password *zend.String, options *zend.Array) *zend.String) { this.hash = value }
func (this *PhpPasswordAlgo) GetVerify() func(password *types2.String, hash *types2.String) types2.ZendBool {
	return this.verify
}

// func (this *PhpPasswordAlgo) SetVerify(value func(password *zend.String, hash *zend.String) zend.ZendBool) { this.verify = value }
func (this *PhpPasswordAlgo) GetNeedsRehash() func(password *types2.String, options *types2.Array) types2.ZendBool {
	return this.needs_rehash
}

// func (this *PhpPasswordAlgo) SetNeedsRehash(value func(password *zend.String, options *zend.Array) zend.ZendBool) { this.needs_rehash = value }
func (this *PhpPasswordAlgo) GetGetInfo() func(return_value *types2.Zval, hash *types2.String) int {
	return this.get_info
}

// func (this *PhpPasswordAlgo) SetGetInfo(value func(return_value *zend.Zval, hash *zend.String) int) { this.get_info = value }
func (this *PhpPasswordAlgo) GetValid() func(hash *types2.String) types2.ZendBool {
	return this.valid
}

// func (this *PhpPasswordAlgo) SetValid(value func(hash *zend.String) zend.ZendBool) { this.valid = value }
