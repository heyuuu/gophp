// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/zend"
)

func SplRegisterInterface(ppce **zend.ZendClassEntry, class_name string, functions *zend.ZendFunctionEntry) {
	var ce zend.ZendClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(zend.ZendStringInitInterned(class_name, strlen(class_name), 1))
	ce.SetBuiltinFunctions(functions)
	*ppce = zend.ZendRegisterInternalInterface(&ce)
}
func SplRegisterStdClass(ppce **zend.ZendClassEntry, class_name string, obj_ctor any, function_list *zend.ZendFunctionEntry) {
	var ce zend.ZendClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(zend.ZendStringInitInterned(class_name, strlen(class_name), 1))
	ce.SetBuiltinFunctions(function_list)
	*ppce = zend.ZendRegisterInternalClass(&ce)

	/* entries changed by initialize */

	if obj_ctor {
		ppce.create_object = obj_ctor
	}

	/* entries changed by initialize */
}
func SplRegisterSubClass(ppce **zend.ZendClassEntry, parent_ce *zend.ZendClassEntry, class_name string, obj_ctor any, function_list *zend.ZendFunctionEntry) {
	var ce zend.ZendClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(zend.ZendStringInitInterned(class_name, strlen(class_name), 1))
	ce.SetBuiltinFunctions(function_list)
	*ppce = zend.ZendRegisterInternalClassEx(&ce, parent_ce)

	/* entries changed by initialize */

	if obj_ctor {
		ppce.create_object = obj_ctor
	} else {
		ppce.create_object = parent_ce.create_object
	}

	/* entries changed by initialize */
}
func SplRegisterProperty(class_entry *zend.ZendClassEntry, prop_name string, prop_name_len int, prop_flags int) {
	zend.ZendDeclarePropertyNull(class_entry, prop_name, prop_name_len, prop_flags)
}
func SplAddClassName(list *zend.Zval, pce *zend.ZendClassEntry, allow int, ce_flags int) {
	if allow == 0 || allow > 0 && pce.HasCeFlags(ce_flags) || allow < 0 && !pce.HasCeFlags(ce_flags) {
		var tmp *zend.Zval
		if b.Assign(&tmp, zend.ZendHashFind(zend.Z_ARRVAL_P(list), pce.GetName())) == nil {
			var t zend.Zval
			zend.ZVAL_STR_COPY(&t, pce.GetName())
			zend.ZendHashAdd(zend.Z_ARRVAL_P(list), pce.GetName(), &t)
		}
	}
}
func SplAddInterfaces(list *zend.Zval, pce *zend.ZendClassEntry, allow int, ce_flags int) {
	var num_interfaces uint32
	if pce.GetNumInterfaces() != 0 {
		zend.ZEND_ASSERT(pce.HasCeFlags(zend.ZEND_ACC_LINKED))
		for num_interfaces = 0; num_interfaces < pce.GetNumInterfaces(); num_interfaces++ {
			SplAddClassName(list, pce.interfaces[num_interfaces], allow, ce_flags)
		}
	}
}
func SplAddTraits(list *zend.Zval, pce *zend.ZendClassEntry, allow int, ce_flags int) {
	var num_traits uint32
	var trait *zend.ZendClassEntry
	for num_traits = 0; num_traits < pce.GetNumTraits(); num_traits++ {
		trait = zend.ZendFetchClassByName(pce.GetTraitNames()[num_traits].GetName(), pce.GetTraitNames()[num_traits].GetLcName(), zend.ZEND_FETCH_CLASS_TRAIT)
		zend.ZEND_ASSERT(trait != nil)
		SplAddClassName(list, trait, allow, ce_flags)
	}
}
func SplAddClasses(pce *zend.ZendClassEntry, list *zend.Zval, sub int, allow int, ce_flags int) int {
	if pce == nil {
		return 0
	}
	SplAddClassName(list, pce, allow, ce_flags)
	if sub != 0 {
		SplAddInterfaces(list, pce, allow, ce_flags)
		for pce.parent {
			pce = pce.parent
			SplAddClasses(pce, list, sub, allow, ce_flags)
		}
	}
	return 0
}
func SplGenPrivatePropName(ce *zend.ZendClassEntry, prop_name string, prop_len int) *zend.ZendString {
	return zend.ZendManglePropertyName(ce.GetName().GetVal(), ce.GetName().GetLen(), prop_name, prop_len, 0)
}
