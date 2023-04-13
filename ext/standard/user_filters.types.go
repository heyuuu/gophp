package standard

import (
	"github.com/heyuuu/gophp/zend/types"
)

/**
 * PhpUserFilterData
 */
type PhpUserFilterData struct {
	ce        *types.ClassEntry
	classname string
}

func NewPhpUserFilterData(ce *types.ClassEntry, classname string) *PhpUserFilterData {
	return &PhpUserFilterData{ce: ce, classname: classname}
}

func (this *PhpUserFilterData) GetCe() *types.ClassEntry { return this.ce }
func (this *PhpUserFilterData) GetClassname() string     { return this.classname }
