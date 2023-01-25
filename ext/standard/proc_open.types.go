// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * PhpProcessEnvT
 */
type PhpProcessEnvT struct {
	envp     *byte
	envarray **byte
}

func (this PhpProcessEnvT) GetEnvp() *byte            { return this.envp }
func (this *PhpProcessEnvT) SetEnvp(value *byte)      { this.envp = value }
func (this PhpProcessEnvT) GetEnvarray() **byte       { return this.envarray }
func (this *PhpProcessEnvT) SetEnvarray(value **byte) { this.envarray = value }

/**
 * PhpProcessHandle
 */
type PhpProcessHandle struct {
	child         PhpProcessIdT
	npipes        int
	pipes         **zend.ZendResource
	command       *byte
	is_persistent int
	env           PhpProcessEnvT
}

func (this PhpProcessHandle) GetChild() PhpProcessIdT             { return this.child }
func (this *PhpProcessHandle) SetChild(value PhpProcessIdT)       { this.child = value }
func (this PhpProcessHandle) GetNpipes() int                      { return this.npipes }
func (this *PhpProcessHandle) SetNpipes(value int)                { this.npipes = value }
func (this PhpProcessHandle) GetPipes() **zend.ZendResource       { return this.pipes }
func (this *PhpProcessHandle) SetPipes(value **zend.ZendResource) { this.pipes = value }
func (this PhpProcessHandle) GetCommand() *byte                   { return this.command }
func (this *PhpProcessHandle) SetCommand(value *byte)             { this.command = value }
func (this PhpProcessHandle) GetIsPersistent() int                { return this.is_persistent }
func (this *PhpProcessHandle) SetIsPersistent(value int)          { this.is_persistent = value }
func (this PhpProcessHandle) GetEnv() PhpProcessEnvT              { return this.env }
func (this *PhpProcessHandle) SetEnv(value PhpProcessEnvT)        { this.env = value }

/**
 * PhpProcOpenDescriptorItem
 */
type PhpProcOpenDescriptorItem struct {
	index      int
	parentend  PhpFileDescriptorT
	childend   PhpFileDescriptorT
	mode       int
	mode_flags int
}

func (this PhpProcOpenDescriptorItem) GetIndex() int                          { return this.index }
func (this *PhpProcOpenDescriptorItem) SetIndex(value int)                    { this.index = value }
func (this PhpProcOpenDescriptorItem) GetParentend() PhpFileDescriptorT       { return this.parentend }
func (this *PhpProcOpenDescriptorItem) SetParentend(value PhpFileDescriptorT) { this.parentend = value }
func (this PhpProcOpenDescriptorItem) GetChildend() PhpFileDescriptorT        { return this.childend }
func (this *PhpProcOpenDescriptorItem) SetChildend(value PhpFileDescriptorT)  { this.childend = value }
func (this PhpProcOpenDescriptorItem) GetMode() int                           { return this.mode }
func (this *PhpProcOpenDescriptorItem) SetMode(value int)                     { this.mode = value }
func (this PhpProcOpenDescriptorItem) GetModeFlags() int                      { return this.mode_flags }
func (this *PhpProcOpenDescriptorItem) SetModeFlags(value int)                { this.mode_flags = value }

/* PhpProcOpenDescriptorItem.mode_flags */
func (this *PhpProcOpenDescriptorItem) AddModeFlags(value int)     { this.mode_flags |= value }
func (this *PhpProcOpenDescriptorItem) SubModeFlags(value int)     { this.mode_flags &^= value }
func (this PhpProcOpenDescriptorItem) HasModeFlags(value int) bool { return this.mode_flags&value != 0 }
func (this *PhpProcOpenDescriptorItem) SwitchModeFlags(value int, cond bool) {
	if cond {
		this.AddModeFlags(value)
	} else {
		this.SubModeFlags(value)
	}
}
