// <<generate>>

package streams

import (
	"sik/core"
	"sik/zend/types"
)

/**
 * PhpUserStreamWrapper
 */
type PhpUserStreamWrapper struct {
	protoname *byte
	classname *byte
	ce        *types.ClassEntry
	wrapper   core.PhpStreamWrapper
}

// func MakePhpUserStreamWrapper(protoname *byte, classname *byte, ce *zend.ClassEntry, wrapper core.PhpStreamWrapper) PhpUserStreamWrapper {
//     return PhpUserStreamWrapper{
//         protoname:protoname,
//         classname:classname,
//         ce:ce,
//         wrapper:wrapper,
//     }
// }
func (this *PhpUserStreamWrapper) GetProtoname() *byte      { return this.protoname }
func (this *PhpUserStreamWrapper) SetProtoname(value *byte) { this.protoname = value }
func (this *PhpUserStreamWrapper) GetClassname() *byte      { return this.classname }
func (this *PhpUserStreamWrapper) SetClassname(value *byte) { this.classname = value }
func (this *PhpUserStreamWrapper) GetCe() *types.ClassEntry { return this.ce }

// func (this *PhpUserStreamWrapper) SetCe(value *zend.ClassEntry) { this.ce = value }
func (this *PhpUserStreamWrapper) GetWrapper() core.PhpStreamWrapper { return this.wrapper }

// func (this *PhpUserStreamWrapper) SetWrapper(value core.PhpStreamWrapper) { this.wrapper = value }

/**
 * _phpUserstreamData
 */
type _phpUserstreamData struct {
	wrapper *PhpUserStreamWrapper
	object  types.Zval
}

// func Make_phpUserstreamData(wrapper *PhpUserStreamWrapper, object zend.Zval) _phpUserstreamData {
//     return _phpUserstreamData{
//         wrapper:wrapper,
//         object:object,
//     }
// }
func (this *_phpUserstreamData) GetWrapper() *PhpUserStreamWrapper      { return this.wrapper }
func (this *_phpUserstreamData) SetWrapper(value *PhpUserStreamWrapper) { this.wrapper = value }
func (this *_phpUserstreamData) GetObject() types.Zval                  { return this.object }

// func (this *_phpUserstreamData) SetObject(value zend.Zval) { this.object = value }
