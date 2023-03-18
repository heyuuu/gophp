// <<generate>>

package standard

import (
	"sik/zend/types"
)

/**
 * ZendAssertGlobals
 */
type ZendAssertGlobals struct {
	callback   types.Zval
	cb         *byte
	active     types.ZendBool
	bail       types.ZendBool
	warning    types.ZendBool
	quiet_eval types.ZendBool
	exception  types.ZendBool
}

//             func MakeZendAssertGlobals(
// callback zend.Zval,
// cb *byte,
// active zend.ZendBool,
// bail zend.ZendBool,
// warning zend.ZendBool,
// quiet_eval zend.ZendBool,
// exception zend.ZendBool,
// ) ZendAssertGlobals {
//                 return ZendAssertGlobals{
//                     callback:callback,
//                     cb:cb,
//                     active:active,
//                     bail:bail,
//                     warning:warning,
//                     quiet_eval:quiet_eval,
//                     exception:exception,
//                 }
//             }
func (this *ZendAssertGlobals) GetCallback() types.Zval { return this.callback }

// func (this *ZendAssertGlobals) SetCallback(value zend.Zval) { this.callback = value }
// func (this *ZendAssertGlobals)  GetCb() *byte      { return this.cb }
func (this *ZendAssertGlobals) SetCb(value *byte)         { this.cb = value }
func (this *ZendAssertGlobals) GetActive() types.ZendBool { return this.active }

// func (this *ZendAssertGlobals) SetActive(value zend.ZendBool) { this.active = value }
func (this *ZendAssertGlobals) GetBail() types.ZendBool { return this.bail }

// func (this *ZendAssertGlobals) SetBail(value zend.ZendBool) { this.bail = value }
func (this *ZendAssertGlobals) GetWarning() types.ZendBool { return this.warning }

// func (this *ZendAssertGlobals) SetWarning(value zend.ZendBool) { this.warning = value }
func (this *ZendAssertGlobals) GetQuietEval() types.ZendBool { return this.quiet_eval }

// func (this *ZendAssertGlobals) SetQuietEval(value zend.ZendBool) { this.quiet_eval = value }
func (this *ZendAssertGlobals) GetException() types.ZendBool { return this.exception }

// func (this *ZendAssertGlobals) SetException(value zend.ZendBool) { this.exception = value }
