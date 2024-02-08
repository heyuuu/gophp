package printer

import (
	"github.com/heyuuu/gophp/kits/mathkit"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"io"
	"strconv"
	"strings"
)

type VarExportPrinter struct {
	basePrinter
	ctx *php.Context
}

func NewVarExportPrinter(ctx *php.Context, w io.Writer) *VarExportPrinter {
	p := &VarExportPrinter{}
	p.ctx = ctx
	p.w = w
	return p
}

func (p *VarExportPrinter) Zval(zv types.Zval, level int) {
	zv = zv.DeRef()
	switch zv.Type() {
	case types.IsFalse:
		p.print("false")
	case types.IsTrue:
		p.print("true")
	case types.IsNull:
		p.print("NULL")
	case types.IsLong:

		/* INT_MIN as a literal will be parsed as a float. Emit something like
		 * -9223372036854775807-1 to avoid this. */

		if zv.Long() == php.LongMin {
			p.print(strconv.Itoa(php.LongMin + 1))
			p.print("-1")
			break
		}
		p.print(strconv.Itoa(zv.Long()))
	case types.IsDouble:
		doubleStr := php.SerializeDouble(zv.Double(), p.ctx.PG().SerializePrecision())
		p.print(doubleStr)

		/* Without a decimal point, PHP treats a number literal as an int.
		 * This check even works for scientific notation, because the
		 * mantissa always contains a decimal point.
		 * We need to check for finiteness, because INF, -INF and NAN
		 * must not have a decimal point added.
		 */
		if mathkit.IsFinite(zv.Double()) && strings.IndexByte(doubleStr, '.') < 0 {
			p.print(".0")
		}
	case types.IsString:
		ztmp := p.escape(zv.String())
		ztmp2 := strings.ReplaceAll(ztmp, "0", `' . "\0" . '`)
		p.print(`'`)
		p.print(ztmp2)
		p.print(`'`)
	case types.IsArray:
		myht := zv.Array()
		if myht.IsRecursive() {
			p.print("NULL")
			php.Error(p.ctx, perr.E_WARNING, "var_export does not handle circular references")
			return
		}
		myht.ProtectRecursive()
		p.print("\n")
		p.printIdent(level)
		p.print("array (\n")
		myht.Each(func(key types.ArrayKey, value types.Zval) {
			p.arrayElement(value, key, level)
		})
		myht.UnprotectRecursive()
		p.printIdent(level)
		p.print(")")
	case types.IsObject:
		myht := php.ZendGetPropertiesFor(zv, types.PropPurposeVarExport)
		if myht != nil {
			if myht.IsRecursive() {
				p.print("NULL")
				php.Error(p.ctx, perr.E_WARNING, "var_export does not handle circular references")
				//zend.ZendReleaseProperties(myht)
				return
			} else {
				myht.ProtectRecursive()
			}
		}
		if level > 1 {
			p.print("\n")
			p.printIdent(level)
		}

		/* stdClass has no __set_state method, but can be casted to */
		if zv.Object().ClassName() == php.StdClassName {
			p.print("(object) array(\n")
		} else {
			p.print(zv.Object().ClassName())
			p.print("::__set_state(array(\n")
		}
		if myht != nil {
			myht.Each(func(key types.ArrayKey, value types.Zval) {
				p.objectElement(value, key, level)
			})
		}
		if level > 1 {
			p.printIdent(level)
		}
		if zv.Object().ClassName() == php.StdClassName {
			p.print(")")
		} else {
			p.print("))")
		}
	default:
		p.print("NULL")
	}
}

func (p *VarExportPrinter) escape(s string) string {
	replacer := strings.NewReplacer(`'`, `\'`, `\`, `\\`)
	return replacer.Replace(s)
}

func (p *VarExportPrinter) arrayElement(zv types.Zval, key types.ArrayKey, level int) {
	if key.IsStrKey() {
		ckey := p.escape(key.StrKey())
		tmpStr := strings.ReplaceAll(ckey, "0", `' . "\0" . '`)
		p.printIdent(level + 2)
		p.print(`'`)
		p.print(tmpStr)
		p.print("' => ")
	} else {
		p.printIdent(level + 2)
		p.print(strconv.Itoa(key.IdxKey()))
		p.print(" => ")
	}
	p.Zval(zv, level+2)
	p.print(",\n")
}

func (p *VarExportPrinter) objectElement(zv types.Zval, key types.ArrayKey, level int) {
	p.printIdent(level + 2)
	if key.IsStrKey() {
		_, propName, _ := php.UnmanglePropertyName(p.ctx, key.StrKey())
		propNameEscaped := p.escape(propName)
		p.print(`'` + propNameEscaped + `'`)
	} else {
		p.print(strconv.Itoa(key.IdxKey()))
	}
	p.print(" => ")
	p.Zval(zv, level+2)
	p.print(",\n")
}
