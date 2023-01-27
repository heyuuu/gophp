// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * PhpDirGlobals
 */
type PhpDirGlobals struct {
	default_dir *zend.ZendResource
}

func (this *PhpDirGlobals) GetDefaultDir() *zend.ZendResource      { return this.default_dir }
func (this *PhpDirGlobals) SetDefaultDir(value *zend.ZendResource) { this.default_dir = value }
