package standard

import (
	"github.com/heyuuu/gophp/zend"
)

func ZmInfoDl(zend_module *zend.ModuleEntry) {
	PhpInfoPrintTableRow(2, "Dynamic Library Support", "enabled")
}
