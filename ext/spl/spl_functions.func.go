package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

func SplRegisterProperty(class_entry *types.ClassEntry, prop_name string, prop_name_len int, prop_flags int) {
	zend.ZendDeclarePropertyNull(class_entry, prop_name, prop_name_len, prop_flags)
}
func SplAddClassName(list *types.Zval, pce *types.ClassEntry, allow int, ceFlags uint32) {
	if allow == 0 || allow > 0 && pce.HasCeFlags(ceFlags) || allow < 0 && !pce.HasCeFlags(ceFlags) {
		var tmp *types.Zval
		if b.Assign(&tmp, list.Array().KeyFind(pce.Name())) == nil {
			var t types.Zval
			t.SetStringVal(pce.Name())
			list.Array().KeyAdd(pce.Name(), &t)
		}
	}
}
func SplAddClassNameEx(list *types.Array, pce *types.ClassEntry, allow int, ceFlags uint32) {
	if allow == 0 || allow > 0 && pce.HasCeFlags(ceFlags) || allow < 0 && !pce.HasCeFlags(ceFlags) {
		if !list.KeyExists(pce.Name()) {
			list.KeyAdd(pce.Name(), types.NewZvalString(pce.Name()))
		}
	}
}
func SplAddInterfaces(list *types.Zval, pce *types.ClassEntry, allow int, ce_flags uint32) {
	if pce.GetNumInterfaces() != 0 {
		b.Assert(pce.HasCeFlags(types.AccLinked))
		for numInterfaces := 0; numInterfaces < pce.GetNumInterfaces(); numInterfaces++ {
			SplAddClassName(list, pce.GetInterfaces()[numInterfaces], allow, ce_flags)
		}
	}
}
func SplAddInterfacesEx(list *types.Array, pce *types.ClassEntry, allow int, ceFlags uint32) {
	if pce.GetNumInterfaces() != 0 {
		b.Assert(pce.HasCeFlags(types.AccLinked))
		for numInterfaces := 0; numInterfaces < pce.GetNumInterfaces(); numInterfaces++ {
			SplAddClassNameEx(list, pce.GetInterfaces()[numInterfaces], allow, ceFlags)
		}
	}
}
func SplAddTraits(list *types.Zval, pce *types.ClassEntry, allow int, ceFlags uint32) {
	var num_traits uint32
	var trait *types.ClassEntry
	for num_traits = 0; num_traits < pce.GetNumTraits(); num_traits++ {
		trait = zend.ZendFetchClassByName(pce.GetTraitNames()[num_traits].GetName(), pce.GetTraitNames()[num_traits].GetLcName(), zend.ZEND_FETCH_CLASS_TRAIT)
		b.Assert(trait != nil)
		SplAddClassName(list, trait, allow, ceFlags)
	}
}
func SplAddTraitsEx(list *types.Array, pce *types.ClassEntry, allow int, ceFlags uint32) {
	var numTraits uint32
	var trait *types.ClassEntry
	for numTraits = 0; numTraits < pce.GetNumTraits(); numTraits++ {
		trait = zend.ZendFetchClassByName_Ex(pce.GetTraitNames()[numTraits].GetName(), pce.GetTraitNames()[numTraits].GetLcName(), zend.ZEND_FETCH_CLASS_TRAIT)
		b.Assert(trait != nil)
		SplAddClassNameEx(list, trait, allow, ceFlags)
	}
}
func SplAddClasses(pce *types.ClassEntry, list *types.Zval, sub int, allow int, ceFlags uint32) int {
	if pce == nil {
		return 0
	}
	SplAddClassName(list, pce, allow, ceFlags)
	if sub != 0 {
		SplAddInterfaces(list, pce, allow, ceFlags)
		for pce.GetParent() {
			pce = pce.GetParent()
			SplAddClasses(pce, list, sub, allow, ceFlags)
		}
	}
	return 0
}
func SplGenPrivatePropName(ce *types.ClassEntry, prop_name string) *types.String {
	str := zend.ZendManglePropertyName_Ex(ce.Name(), prop_name)
	return types.NewString(str)
}
