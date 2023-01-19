// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * ZendAssertGlobals
 */
type ZendAssertGlobals struct {
	callback   zend.Zval
	cb         *byte
	active     zend.ZendBool
	bail       zend.ZendBool
	warning    zend.ZendBool
	quiet_eval zend.ZendBool
	exception  zend.ZendBool
}

func (this ZendAssertGlobals) GetCallback() zend.Zval            { return this.callback }
func (this *ZendAssertGlobals) SetCallback(value zend.Zval)      { this.callback = value }
func (this ZendAssertGlobals) GetCb() *byte                      { return this.cb }
func (this *ZendAssertGlobals) SetCb(value *byte)                { this.cb = value }
func (this ZendAssertGlobals) GetActive() zend.ZendBool          { return this.active }
func (this *ZendAssertGlobals) SetActive(value zend.ZendBool)    { this.active = value }
func (this ZendAssertGlobals) GetBail() zend.ZendBool            { return this.bail }
func (this *ZendAssertGlobals) SetBail(value zend.ZendBool)      { this.bail = value }
func (this ZendAssertGlobals) GetWarning() zend.ZendBool         { return this.warning }
func (this *ZendAssertGlobals) SetWarning(value zend.ZendBool)   { this.warning = value }
func (this ZendAssertGlobals) GetQuietEval() zend.ZendBool       { return this.quiet_eval }
func (this *ZendAssertGlobals) SetQuietEval(value zend.ZendBool) { this.quiet_eval = value }
func (this ZendAssertGlobals) GetException() zend.ZendBool       { return this.exception }
func (this *ZendAssertGlobals) SetException(value zend.ZendBool) { this.exception = value }
