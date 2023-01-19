// <<generate>>

package streams

/**
 * PhpUserStreamWrapper
 */
type PhpUserStreamWrapper struct {
	protoname *byte
	classname *byte
	ce        *zend.ZendClassEntry
	wrapper   core.PhpStreamWrapper
}

func (this PhpUserStreamWrapper) GetProtoname() *byte                     { return this.protoname }
func (this *PhpUserStreamWrapper) SetProtoname(value *byte)               { this.protoname = value }
func (this PhpUserStreamWrapper) GetClassname() *byte                     { return this.classname }
func (this *PhpUserStreamWrapper) SetClassname(value *byte)               { this.classname = value }
func (this PhpUserStreamWrapper) GetCe() *zend.ZendClassEntry             { return this.ce }
func (this *PhpUserStreamWrapper) SetCe(value *zend.ZendClassEntry)       { this.ce = value }
func (this PhpUserStreamWrapper) GetWrapper() core.PhpStreamWrapper       { return this.wrapper }
func (this *PhpUserStreamWrapper) SetWrapper(value core.PhpStreamWrapper) { this.wrapper = value }

/**
 * _phpUserstreamData
 */
type _phpUserstreamData struct {
	wrapper *PhpUserStreamWrapper
	object  zend.Zval
}

func (this _phpUserstreamData) GetWrapper() *PhpUserStreamWrapper       { return this.wrapper }
func (this *_phpUserstreamData) SetWrapper(value *PhpUserStreamWrapper) { this.wrapper = value }
func (this _phpUserstreamData) GetObject() zend.Zval                    { return this.object }
func (this *_phpUserstreamData) SetObject(value zend.Zval)              { this.object = value }
