// <<generate>>

package core

import (
	"sik/zend"
)

/**
 * PhpExtensionLists
 */
type PhpExtensionLists struct {
	engine    zend.ZendLlist
	functions zend.ZendLlist
}

// func MakePhpExtensionLists(engine zend.ZendLlist, functions zend.ZendLlist) PhpExtensionLists {
//     return PhpExtensionLists{
//         engine:engine,
//         functions:functions,
//     }
// }
func (this *PhpExtensionLists) GetEngine() zend.ZendLlist { return this.engine }

// func (this *PhpExtensionLists) SetEngine(value zend.ZendLlist) { this.engine = value }
func (this *PhpExtensionLists) GetFunctions() zend.ZendLlist { return this.functions }

// func (this *PhpExtensionLists) SetFunctions(value zend.ZendLlist) { this.functions = value }
