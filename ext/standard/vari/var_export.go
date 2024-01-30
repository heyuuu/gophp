package vari

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"strconv"
	"strings"
)

func BufferAppendSpaces(buf *strings.Builder, numSpaces int) {
	buf.WriteString(fmt.Sprintf("%*c", numSpaces, ' '))
}

type VarExportPrinter struct {
	ctx *php.Context
	buf strings.Builder
}

func NewVarExportPrinter(ctx *php.Context) *VarExportPrinter {
	return &VarExportPrinter{ctx: ctx}
}

func (p *VarExportPrinter) String() string {
	return p.buf.String()
}

func (p *VarExportPrinter) escape(s string) string {
	replacer := strings.NewReplacer(`'`, `\'`, `\`, `\\`)
	return replacer.Replace(s)
}

func (p *VarExportPrinter) ArrayElement(zv types.Zval, key types.ArrayKey, level int) {
	buf := &p.buf
	if key.IsStrKey() {
		ckey := p.escape(key.StrKey())
		tmpStr := strings.ReplaceAll(ckey, "0", "' . \"\\0\" . '")
		BufferAppendSpaces(buf, level+1)
		buf.WriteByte('\'')
		buf.WriteString(tmpStr)
		buf.WriteString("' => ")
	} else {
		BufferAppendSpaces(buf, level+1)
		buf.WriteString(strconv.Itoa(key.IdxKey()))
		buf.WriteString(" => ")
	}
	p.Zval(zv, level+2)
	buf.WriteByte(',')
	buf.WriteByte('\n')
}

func (p *VarExportPrinter) ObjectElement(zv types.Zval, key types.ArrayKey, level int) {
	//ctx := p.ctx
	buf := &p.buf
	BufferAppendSpaces(buf, level+2)
	if key.IsStrKey() {
		//_, propName, _ := php.UnmanglePropertyName(ctx, key.StrKey())
		propName := key.StrKey()
		propNameEscaped := p.escape(propName)
		buf.WriteByte('\'')
		buf.WriteString(propNameEscaped)
		buf.WriteByte('\'')
	} else {
		buf.WriteString(strconv.Itoa(key.IdxKey()))
	}
	buf.WriteString(" => ")
	p.Zval(zv, level+2)
	buf.WriteByte(',')
	buf.WriteByte('\n')
}

func (p *VarExportPrinter) Zval(struc types.Zval, level int) {
	ctx := p.ctx
	buf := &p.buf

	struc = struc.DeRef()
	switch struc.Type() {
	case types.IsFalse:
		buf.WriteString("false")
	case types.IsTrue:
		buf.WriteString("true")
	case types.IsNull:
		buf.WriteString("NULL")
	case types.IsLong:

		/* INT_MIN as a literal will be parsed as a float. Emit something like
		 * -9223372036854775807-1 to avoid this. */

		if struc.Long() == php.LongMin {
			buf.WriteString(strconv.Itoa(php.LongMin + 1))
			buf.WriteString("-1")
			break
		}
		buf.WriteString(strconv.Itoa(struc.Long()))
	case types.IsDouble:
		//var tmp_str []byte
		//php.PhpGcvt(struc.Double(), int(php.PG__(ctx).SerializePrecision()), '.', 'E', tmp_str)
		//buf.WriteString(builtin.CastStrAuto(tmp_str))
		//
		///* Without a decimal point, PHP treats a number literal as an int.
		// * This check even works for scientific notation, because the
		// * mantissa always contains a decimal point.
		// * We need to check for finiteness, because INF, -INF and NAN
		// * must not have a decimal point added.
		// */
		//if mathkit.IsFinite(struc.Double()) && nil == strchr(tmp_str, '.') {
		//	buf.WriteString(".0")
		//}
	case types.IsString:
		ztmp := p.escape(struc.String())
		ztmp2 := strings.ReplaceAll(ztmp, "0", "' . \"\\0\" . '")
		buf.WriteByte('\'')
		buf.WriteString(ztmp2)
		buf.WriteByte('\'')
		//types.ZendStringFree(ztmp)
		//types.ZendStringFree(ztmp2)
	case types.IsArray:
		myht := struc.Array()
		if myht.IsRecursive() {
			buf.WriteString("NULL")
			php.Error(ctx, perr.E_WARNING, "var_export does not handle circular references")
			return
		}
		myht.ProtectRecursive()
		if level > 1 {
			buf.WriteByte('\n')
			BufferAppendSpaces(buf, level-1)
		}
		buf.WriteString("array (\n")
		myht.Each(func(key types.ArrayKey, value types.Zval) {
			p.ArrayElement(value, key, level)
		})
		myht.UnprotectRecursive()
		if level > 1 {
			BufferAppendSpaces(buf, level-1)
		}
		buf.WriteByte(')')
	case types.IsObject:
		//myht = php.ZendGetPropertiesFor(struc, types.PropPurposeVarExport)
		//if myht != nil {
		//	if myht.IsRecursive() {
		//		buf.WriteString("NULL")
		//		php.Error(ctx, perr.E_WARNING, "var_export does not handle circular references")
		//		//zend.ZendReleaseProperties(myht)
		//		return
		//	} else {
		//		myht.ProtectRecursive()
		//	}
		//}
		//if level > 1 {
		//	buf.WriteByte('\n')
		//	BufferAppendSpaces(buf, level-1)
		//}
		//
		///* stdClass has no __set_state method, but can be casted to */
		//
		//if types.Z_OBJCE_P(struc) == php.ZendStandardClassDef {
		//	buf.WriteString("(object) array(\n")
		//} else {
		//	buf.WriteString(types.Z_OBJCE_P(struc).Name())
		//	buf.WriteString("::__set_state(array(\n")
		//}
		//if myht != nil {
		//	myht.Each(func(key types.ArrayKey, value types.Zval) {
		//		p.ObjectElement(value, key, level)
		//	})
		//}
		//if level > 1 {
		//	BufferAppendSpaces(buf, level-1)
		//}
		//if types.Z_OBJCE_P(struc) == php.ZendStandardClassDef {
		//	buf.WriteByte(')')
		//} else {
		//	buf.WriteString("))")
		//}
	default:
		buf.WriteString("NULL")
	}
}
