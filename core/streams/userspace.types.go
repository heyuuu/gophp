package streams

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
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

func (this *PhpUserStreamWrapper) GetProtoname() *byte               { return this.protoname }
func (this *PhpUserStreamWrapper) SetProtoname(value *byte)          { this.protoname = value }
func (this *PhpUserStreamWrapper) GetClassname() *byte               { return this.classname }
func (this *PhpUserStreamWrapper) SetClassname(value *byte)          { this.classname = value }
func (this *PhpUserStreamWrapper) GetCe() *types.ClassEntry          { return this.ce }
func (this *PhpUserStreamWrapper) GetWrapper() core.PhpStreamWrapper { return this.wrapper }

/**
 * _phpUserstreamData
 */
type _phpUserstreamData struct {
	wrapper *PhpUserStreamWrapper
	object  types.Zval
}

func (this *_phpUserstreamData) GetWrapper() *PhpUserStreamWrapper      { return this.wrapper }
func (this *_phpUserstreamData) SetWrapper(value *PhpUserStreamWrapper) { this.wrapper = value }
func (this *_phpUserstreamData) GetObject() types.Zval                  { return this.object }
