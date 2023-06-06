package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

const ZEND_EXTENSION_API_NO = 320190902

/* Typedef's for zend_extension function pointers */

type StartupFuncT func(extension *ZendExtension) int
type ShutdownFuncT func(extension *ZendExtension)
type ActivateFuncT func()
type DeactivateFuncT func()
type MessageHandlerFuncT func(message int, arg any)
type StatementHandlerFuncT func(frame *ZendExecuteData)
type FcallBeginHandlerFuncT func(frame *ZendExecuteData)
type FcallEndHandlerFuncT func(frame *ZendExecuteData)

const ZEND_EXTMSG_NEW_EXTENSION = 1

var ZendExtensions ZendExtensionsT

type ZendExtensionsT struct {
	data []*ZendExtension
}

func (extensions *ZendExtensionsT) Init() {
	extensions.data = nil
}

func (extensions *ZendExtensionsT) Startup() {
	extensions.ApplyWithDel(func(ext *ZendExtension) int {
		if ext.startup != nil {
			if ext.startup(ext) != types.SUCCESS {
				return 1
			}
			ZendAppendVersionInfo(ext)
		}
		return 0
	})
}

func (extensions *ZendExtensionsT) Shutdown() {
	extensions.Apply(func(ext *ZendExtension) {
		if ext.shutdown != nil {
			ext.shutdown(ext)
		}
	})
	extensions.Destroy()
}

func (extensions *ZendExtensionsT) Destroy() {
	for _, ext := range extensions.data {
		if ext.handle != nil && !b.HasEnv("ZEND_DONT_UNLOAD_MODULES") {
			DL_UNLOAD(ext.handle)
		}
	}
	extensions.data = nil
}

func (extensions *ZendExtensionsT) Elements() []*ZendExtension {
	var result []*ZendExtension
	if len(extensions.data) != 0 {
		result = make([]*ZendExtension, len(extensions.data))
		copy(result, extensions.data)
	}
	return result
}

func (extensions *ZendExtensionsT) Register(newExt *ZendExtension, handle any) {
	var ext ZendExtension = *newExt
	ext.SetHandle(handle)
	extensions.DispatchMessage(ZEND_EXTMSG_NEW_EXTENSION, &ext)
	extensions.data = append(extensions.data, &ext)
}

func (extensions *ZendExtensionsT) Get(name string) *ZendExtension {
	for _, ext := range extensions.data {
		if ext.GetNameStr() == name {
			return ext
		}
	}
	return nil
}

func (extensions *ZendExtensionsT) Apply(f func(ext *ZendExtension)) {
	for _, ext := range extensions.data {
		f(ext)
	}
}

func (extensions *ZendExtensionsT) ApplyWithDel(f func(ext *ZendExtension) int) {
	if len(extensions.data) == 0 {
		return
	}

	var newData []*ZendExtension
	for _, ext := range extensions.data {
		if f(ext) == 0 {
			newData = append(newData, ext)
		}
	}
	extensions.data = newData
}

func (extensions *ZendExtensionsT) DispatchMessage(message int, arg any) {
	extensions.Apply(func(ext *ZendExtension) {
		if ext.GetMessageHandler() != nil {
			ext.GetMessageHandler()(message, arg)
		}
	})
}
