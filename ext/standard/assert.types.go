package standard

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

/**
 * ZendAssertGlobals
 */
type ZendAssertGlobals struct {
	callback   types2.Zval
	cb         *byte
	active     types2.ZendBool
	bail       types2.ZendBool
	warning    types2.ZendBool
	quiet_eval types2.ZendBool
	exception  types2.ZendBool
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
func (this *ZendAssertGlobals) GetCallback() types2.Zval { return this.callback }

// func (this *ZendAssertGlobals) SetCallback(value zend.Zval) { this.callback = value }
// func (this *ZendAssertGlobals)  GetCb() *byte      { return this.cb }
func (this *ZendAssertGlobals) SetCb(value *byte)          { this.cb = value }
func (this *ZendAssertGlobals) GetActive() types2.ZendBool { return this.active }

// func (this *ZendAssertGlobals) SetActive(value zend.ZendBool) { this.active = value }
func (this *ZendAssertGlobals) GetBail() types2.ZendBool { return this.bail }

// func (this *ZendAssertGlobals) SetBail(value zend.ZendBool) { this.bail = value }
func (this *ZendAssertGlobals) GetWarning() types2.ZendBool { return this.warning }

// func (this *ZendAssertGlobals) SetWarning(value zend.ZendBool) { this.warning = value }
func (this *ZendAssertGlobals) GetQuietEval() types2.ZendBool { return this.quiet_eval }

// func (this *ZendAssertGlobals) SetQuietEval(value zend.ZendBool) { this.quiet_eval = value }
func (this *ZendAssertGlobals) GetException() types2.ZendBool { return this.exception }

// func (this *ZendAssertGlobals) SetException(value zend.ZendBool) { this.exception = value }
