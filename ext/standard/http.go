package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"strconv"
	"strings"
)

const URL_DEFAULT_ARG_SEP = "&"

func urlEncodeHash(ctx *php.Context, ht *types.Array, buf *strings.Builder, numPrefix string, keyPrefix string, keySuffix string, typ *types.Object, argSep string, encType int) bool {
	var propName string
	if ht == nil {
		return false
	}
	if ht.IsRecursive() {
		/* Prevent recursion */
		return true
	}

	if argSep == "" {
		argSep = URL_DEFAULT_ARG_SEP
	}
	ht.Each(func(key types.ArrayKey, zdata types.Zval) {
		/* handling for private & protected object properties */
		if key.IsStrKey() {
			propName = key.StrKey()
			if key.StrKey()[0] == '\000' && typ != nil {
				_, propName, _ = php.UnmanglePropertyName(ctx, key.StrKey())
			} else {
				propName = key.StrKey()
			}
		} else {
			propName = ""
		}
		zdata = zdata.DeRef()
		if zdata.IsArray() || zdata.IsObject() {
			var newPrefix string
			if key.IsStrKey() {
				var ekey string
				if encType == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(propName)
				} else {
					ekey = PhpUrlEncode(propName)
				}
				newPrefix = keyPrefix + ekey + keySuffix + "%5B"
			} else {
				/* Is an integer key */
				ekey := strconv.Itoa(key.IdxKey())
				newPrefix = keyPrefix + numPrefix + ekey + keySuffix + "%5B"
			}
			ht.ProtectRecursive()
			urlEncodeHash(ctx, php.HashOf(zdata), buf, "", newPrefix, "%5D", lang.Cond(zdata.IsObject(), zdata.Object(), nil), argSep, encType)
			ht.UnprotectRecursive()
		} else if zdata.IsNull() || zdata.IsResource() {
			/* Skip these types */
			return
		} else {
			if buf.Len() != 0 {
				buf.WriteString(argSep)
			}

			/* Simple key=value */
			buf.WriteString(keyPrefix)
			if key.IsStrKey() {
				var ekey string
				if encType == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(propName)
				} else {
					ekey = PhpUrlEncode(propName)
				}
				buf.WriteString(ekey)
			} else {
				/* Numeric key */
				buf.WriteString(numPrefix)
				buf.WriteString(strconv.Itoa(key.IdxKey()))
			}
			buf.WriteString(keySuffix)
			buf.WriteString("=")
			switch zdata.Type() {
			case types.IsString:
				var ekey string
				if encType == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(zdata.String())
				} else {
					ekey = PhpUrlEncode(zdata.String())
				}
				buf.WriteString(ekey)
			case types.IsLong:
				buf.WriteString(strconv.Itoa(zdata.Long()))
			case types.IsFalse:
				buf.WriteString("0")
			case types.IsTrue:
				buf.WriteString("1")
			default:
				var ekey string
				var str = php.ZvalGetStrVal(ctx, zdata)
				if encType == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(str)
				} else {
					ekey = PhpUrlEncode(str)
				}
				buf.WriteString(ekey)
			}
		}
	})
	return true
}

// @zif(onError=1)
func ZifHttpBuildQuery(ctx *php.Context, formdata types.Zval, _ zpp.Opt, prefix string, argSeparator string, encType_ *int) (string, bool) {
	var encType = lang.Option(encType_, PHP_QUERY_RFC1738)
	var buf strings.Builder
	ret := urlEncodeHash(ctx, php.HashOf(formdata), &buf, prefix, "", "", lang.CondF1(formdata.IsObject(), formdata.Object, nil), argSeparator, encType)
	return buf.String(), ret
}
