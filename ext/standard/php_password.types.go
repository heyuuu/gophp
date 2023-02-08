// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * PhpPasswordAlgo
 */
type PhpPasswordAlgo struct {
	name         *byte
	hash         func(password *zend.ZendString, options *zend.ZendArray) *zend.ZendString
	verify       func(password *zend.ZendString, hash *zend.ZendString) zend.ZendBool
	needs_rehash func(password *zend.ZendString, options *zend.ZendArray) zend.ZendBool
	get_info     func(return_value *zend.Zval, hash *zend.ZendString) int
	valid        func(hash *zend.ZendString) zend.ZendBool
}

func MakePhpPasswordAlgo(
	name *byte,
	hash func(password *zend.ZendString, options *zend.ZendArray) *zend.ZendString,
	verify func(password *zend.ZendString, hash *zend.ZendString) zend.ZendBool,
	needs_rehash func(password *zend.ZendString, options *zend.ZendArray) zend.ZendBool,
	get_info func(return_value *zend.Zval, hash *zend.ZendString) int,
	valid func(hash *zend.ZendString) zend.ZendBool,
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
func (this *PhpPasswordAlgo) GetHash() func(password *zend.ZendString, options *zend.ZendArray) *zend.ZendString {
	return this.hash
}

// func (this *PhpPasswordAlgo) SetHash(value func(password *zend.ZendString, options *zend.ZendArray) *zend.ZendString) { this.hash = value }
func (this *PhpPasswordAlgo) GetVerify() func(password *zend.ZendString, hash *zend.ZendString) zend.ZendBool {
	return this.verify
}

// func (this *PhpPasswordAlgo) SetVerify(value func(password *zend.ZendString, hash *zend.ZendString) zend.ZendBool) { this.verify = value }
func (this *PhpPasswordAlgo) GetNeedsRehash() func(password *zend.ZendString, options *zend.ZendArray) zend.ZendBool {
	return this.needs_rehash
}

// func (this *PhpPasswordAlgo) SetNeedsRehash(value func(password *zend.ZendString, options *zend.ZendArray) zend.ZendBool) { this.needs_rehash = value }
func (this *PhpPasswordAlgo) GetGetInfo() func(return_value *zend.Zval, hash *zend.ZendString) int {
	return this.get_info
}

// func (this *PhpPasswordAlgo) SetGetInfo(value func(return_value *zend.Zval, hash *zend.ZendString) int) { this.get_info = value }
func (this *PhpPasswordAlgo) GetValid() func(hash *zend.ZendString) zend.ZendBool { return this.valid }

// func (this *PhpPasswordAlgo) SetValid(value func(hash *zend.ZendString) zend.ZendBool) { this.valid = value }
