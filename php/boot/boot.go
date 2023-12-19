package boot

import (
	_ "github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php"
)

func init() {
	php.MarkIsBoot(true)
}
