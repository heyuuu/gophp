package streams

import (
	"github.com/heyuuu/gophp/core"
	types2 "github.com/heyuuu/gophp/php/types"
)

/**
 * PhpUserStreamWrapper
 */
type PhpUserStreamWrapper struct {
	protoname *byte
	classname *byte
	ce        *types2.ClassEntry
	wrapper   core.PhpStreamWrapper
}

func (this *PhpUserStreamWrapper) GetProtoname() *byte               { return this.protoname }
func (this *PhpUserStreamWrapper) SetProtoname(value *byte)          { this.protoname = value }
func (this *PhpUserStreamWrapper) GetClassname() *byte               { return this.classname }
func (this *PhpUserStreamWrapper) SetClassname(value *byte)          { this.classname = value }
func (this *PhpUserStreamWrapper) GetCe() *types2.ClassEntry         { return this.ce }
func (this *PhpUserStreamWrapper) GetWrapper() core.PhpStreamWrapper { return this.wrapper }

/**
 * _phpUserstreamData
 */
type _phpUserstreamData struct {
	wrapper *PhpUserStreamWrapper
	object  types2.Zval
}

func (this *_phpUserstreamData) GetWrapper() *PhpUserStreamWrapper      { return this.wrapper }
func (this *_phpUserstreamData) SetWrapper(value *PhpUserStreamWrapper) { this.wrapper = value }
func (this *_phpUserstreamData) GetObject() types2.Zval                 { return this.object }
