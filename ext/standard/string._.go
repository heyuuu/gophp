package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZifPathinfo(path string, _ zpp.Opt, options *int) *types.Zval {
	opt := b.Option(options, PHP_PATHINFO_ALL)

	var tmp types.Zval
	var ret *types.String = nil

	zend.ArrayInit(&tmp)
	if (opt & PHP_PATHINFO_DIRNAME) == PHP_PATHINFO_DIRNAME {
		dirname := zend.ZendDirname(path)
		if dirname != "" {
			zend.AddAssocStr(&tmp, "dirname", dirname)
		}
		zend.Efree(dirname)
	}

	haveBasename := (opt & PHP_PATHINFO_BASENAME) == PHP_PATHINFO_BASENAME
	if haveBasename {
		ret = PhpBasenameZStr(b.CastStr(path, path_len), "")
		zend.AddAssocStr(&tmp, "basename", ret.GetStr())
	}
	if (opt & PHP_PATHINFO_EXTENSION) == PHP_PATHINFO_EXTENSION {
		var p *byte
		var idx ptrdiff_t
		if haveBasename == 0 {
			ret = PhpBasenameZStr(b.CastStr(path, path_len), "")
		}
		p = zend.ZendMemrchr(ret.GetVal(), '.', ret.GetLen())
		if p != nil {
			idx = p - ret.GetVal()
			zend.AddAssocStringl(&tmp, "extension", ret.GetVal()+idx+1, ret.GetLen()-idx-1)
		}
	}
	if (opt & PHP_PATHINFO_FILENAME) == PHP_PATHINFO_FILENAME {
		var p *byte
		var idx ptrdiff_t

		/* Have we already looked up the basename? */

		if haveBasename == 0 && ret == nil {
			ret = PhpBasenameZStr(b.CastStr(path, path_len), "")
		}
		p = zend.ZendMemrchr(ret.GetVal(), '.', ret.GetLen())
		if p != nil {
			idx = p - ret.GetVal()
		} else {
			idx = ptrdiff_t(ret.GetLen())
		}
		zend.AddAssocStringl(&tmp, "filename", ret.GetVal(), idx)
	}

	if opt == PHP_PATHINFO_ALL {
		return &tmp
	} else {
		var element *types.Zval = types.ZendHashGetCurrentData(tmp.GetArr())
		if element != nil {
			return element
		} else {
			return types.NewZvalString("")
		}
	}
}
