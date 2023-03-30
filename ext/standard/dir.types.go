package standard

import (
	"github.com/heyuuu/gophp/zend/types"
)

/**
 * PhpDirGlobals
 */
type PhpDirGlobals struct {
	default_dir *types.ZendResource
}

// func MakePhpDirGlobals(default_dir *zend.ZendResource) PhpDirGlobals {
//     return PhpDirGlobals{
//         default_dir:default_dir,
//     }
// }
// func (this *PhpDirGlobals)  GetDefaultDir() *zend.ZendResource      { return this.default_dir }
// func (this *PhpDirGlobals) SetDefaultDir(value *zend.ZendResource) { this.default_dir = value }
