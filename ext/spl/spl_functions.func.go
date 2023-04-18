package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

func SplRegisterInterface(ppce **types2.ClassEntry, class_name string, functions *types2.FunctionEntry) {
	var ce types2.ClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetNameVal(class_name)
	ce.SetBuiltinFunctions(functions)
	*ppce = zend.ZendRegisterInternalInterface(&ce)
}
func SplRegisterStdClass(ppce **types2.ClassEntry, class_name string, obj_ctor any, function_list *types2.FunctionEntry) {
	var ce types2.ClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetNameVal(class_name)
	ce.SetBuiltinFunctions(function_list)
	*ppce = zend.ZendRegisterInternalClass(&ce)

	/* entries changed by initialize */

	if obj_ctor {
		ppce.SetCreateObject(obj_ctor)
	}

	/* entries changed by initialize */
}
func SplRegisterSubClass(ppce **types2.ClassEntry, parent_ce *types2.ClassEntry, class_name string, obj_ctor any, function_list *types2.FunctionEntry) {
	var ce types2.ClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetNameVal(class_name)
	ce.SetBuiltinFunctions(function_list)
	*ppce = zend.ZendRegisterInternalClassEx(&ce, parent_ce)

	/* entries changed by initialize */

	if obj_ctor {
		ppce.SetCreateObject(obj_ctor)
	} else {
		ppce.SetCreateObject(parent_ce.GetCreateObject())
	}

	/* entries changed by initialize */
}
func SplRegisterProperty(class_entry *types2.ClassEntry, prop_name string, prop_name_len int, prop_flags int) {
	zend.ZendDeclarePropertyNull(class_entry, prop_name, prop_name_len, prop_flags)
}
func SplAddClassName(list *types2.Zval, pce *types2.ClassEntry, allow int, ce_flags int) {
	if allow == 0 || allow > 0 && pce.HasCeFlags(ce_flags) || allow < 0 && !pce.HasCeFlags(ce_flags) {
		var tmp *types2.Zval
		if b.Assign(&tmp, list.Array().KeyFind(pce.GetName().GetStr())) == nil {
			var t types2.Zval
			t.SetStringCopy(pce.GetName())
			list.Array().KeyAdd(pce.GetName().GetStr(), &t)
		}
	}
}
func SplAddInterfaces(list *types2.Zval, pce *types2.ClassEntry, allow int, ce_flags int) {
	var num_interfaces uint32
	if pce.GetNumInterfaces() != 0 {
		b.Assert(pce.HasCeFlags(zend.AccLinked))
		for num_interfaces = 0; num_interfaces < pce.GetNumInterfaces(); num_interfaces++ {
			SplAddClassName(list, pce.GetInterfaces()[num_interfaces], allow, ce_flags)
		}
	}
}
func SplAddTraits(list *types2.Zval, pce *types2.ClassEntry, allow int, ce_flags int) {
	var num_traits uint32
	var trait *types2.ClassEntry
	for num_traits = 0; num_traits < pce.GetNumTraits(); num_traits++ {
		trait = zend.ZendFetchClassByName(pce.GetTraitNames()[num_traits].GetName(), pce.GetTraitNames()[num_traits].GetLcName(), zend.ZEND_FETCH_CLASS_TRAIT)
		b.Assert(trait != nil)
		SplAddClassName(list, trait, allow, ce_flags)
	}
}
func SplAddClasses(pce *types2.ClassEntry, list *types2.Zval, sub int, allow int, ce_flags int) int {
	if pce == nil {
		return 0
	}
	SplAddClassName(list, pce, allow, ce_flags)
	if sub != 0 {
		SplAddInterfaces(list, pce, allow, ce_flags)
		for pce.GetParent() {
			pce = pce.GetParent()
			SplAddClasses(pce, list, sub, allow, ce_flags)
		}
	}
	return 0
}
func SplGenPrivatePropName(ce *types2.ClassEntry, prop_name string) *types2.String {
	str := zend.ZendManglePropertyName_Ex(ce.GetName().GetStr(), prop_name)
	return types2.NewString(str)
}
