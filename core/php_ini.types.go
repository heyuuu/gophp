// <<generate>>

package core

/**
 * PhpExtensionLists
 */
type PhpExtensionLists struct {
	engine    zend.ZendLlist
	functions zend.ZendLlist
}

func (this PhpExtensionLists) GetEngine() zend.ZendLlist          { return this.engine }
func (this *PhpExtensionLists) SetEngine(value zend.ZendLlist)    { this.engine = value }
func (this PhpExtensionLists) GetFunctions() zend.ZendLlist       { return this.functions }
func (this *PhpExtensionLists) SetFunctions(value zend.ZendLlist) { this.functions = value }
