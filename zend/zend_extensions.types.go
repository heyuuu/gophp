package zend

import (
	b "github.com/heyuuu/gophp/builtin"
)

/**
 * ZendExtensionVersionInfo
 */
type ZendExtensionVersionInfo struct {
	zend_extension_api_no int
	build_id              *byte
}

func (this *ZendExtensionVersionInfo) GetZendExtensionApiNo() int { return this.zend_extension_api_no }
func (this *ZendExtensionVersionInfo) GetBuildId() *byte          { return this.build_id }

/**
 * ZendExtension
 */
type ZendExtension struct {
	name                *byte
	version             *byte
	author              *byte
	URL                 *byte
	copyright           *byte
	startup             StartupFuncT
	shutdown            ShutdownFuncT
	activate            ActivateFuncT
	deactivate          DeactivateFuncT
	message_handler     MessageHandlerFuncT
	statement_handler   StatementHandlerFuncT
	fcall_begin_handler FcallBeginHandlerFuncT
	fcall_end_handler   FcallEndHandlerFuncT
	api_no_check        func(api_no int) int
	build_id_check      func(build_id *byte) int
	reserved5           any
	reserved6           any
	reserved7           any
	reserved8           any
	handle              any
	resource_number     int
}

func (this *ZendExtension) GetNameStr() string      { return b.CastStrAuto(this.name) }
func (this *ZendExtension) GetVersionStr() string   { return b.CastStrAuto(this.version) }
func (this *ZendExtension) GetAuthorStr() string    { return b.CastStrAuto(this.author) }
func (this *ZendExtension) GetURLStr() string       { return b.CastStrAuto(this.URL) }
func (this *ZendExtension) GetCopyrightStr() string { return b.CastStrAuto(this.copyright) }

/**
 * generate
 */
func (this *ZendExtension) GetName() *byte                             { return this.name }
func (this *ZendExtension) GetVersion() *byte                          { return this.version }
func (this *ZendExtension) GetAuthor() *byte                           { return this.author }
func (this *ZendExtension) GetURL() *byte                              { return this.URL }
func (this *ZendExtension) GetCopyright() *byte                        { return this.copyright }
func (this *ZendExtension) GetStartup() StartupFuncT                   { return this.startup }
func (this *ZendExtension) GetShutdown() ShutdownFuncT                 { return this.shutdown }
func (this *ZendExtension) GetActivate() ActivateFuncT                 { return this.activate }
func (this *ZendExtension) GetDeactivate() DeactivateFuncT             { return this.deactivate }
func (this *ZendExtension) GetMessageHandler() MessageHandlerFuncT     { return this.message_handler }
func (this *ZendExtension) GetStatementHandler() StatementHandlerFuncT { return this.statement_handler }
func (this *ZendExtension) GetFcallBeginHandler() FcallBeginHandlerFuncT {
	return this.fcall_begin_handler
}
func (this *ZendExtension) GetFcallEndHandler() FcallEndHandlerFuncT  { return this.fcall_end_handler }
func (this *ZendExtension) GetApiNoCheck() func(api_no int) int       { return this.api_no_check }
func (this *ZendExtension) GetBuildIdCheck() func(build_id *byte) int { return this.build_id_check }
func (this *ZendExtension) GetHandle() any                            { return this.handle }
func (this *ZendExtension) SetHandle(value any)                       { this.handle = value }
func (this *ZendExtension) SetResourceNumber(value int)               { this.resource_number = value }
