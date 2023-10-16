package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

func SplAddClassName(list *types.Array, pce *types.ClassEntry, allow int, ceFlags uint32) {
	if allow == 0 || allow > 0 && pce.HasCeFlags(ceFlags) || allow < 0 && !pce.HasCeFlags(ceFlags) {
		if !list.KeyExists(pce.Name()) {
			list.KeyAdd(pce.Name(), types.NewZvalString(pce.Name()))
		}
	}
}
func SplAddInterfaces(list *types.Array, pce *types.ClassEntry, allow int, ceFlags uint32) {
	if pce.HasInterfaces() {
		b.Assert(pce.HasCeFlags(types.AccLinked))
		for _, iface := range pce.GetInterfaces() {
			SplAddClassName(list, iface, allow, ceFlags)
		}
	}
}
func SplAddTraits(list *types.Array, pce *types.ClassEntry, allow int, ceFlags uint32) {
	for _, traitName := range pce.GetTraitNames() {
		trait := zend.ZendFetchClassByNameEx(traitName, zend.ZEND_FETCH_CLASS_TRAIT)
		b.Assert(trait != nil)
		SplAddClassName(list, trait, allow, ceFlags)
	}
}

func SplGenPrivatePropName(ce *types.ClassEntry, prop_name string) *types.String {
	str := zend.ZendManglePropertyName_Ex(ce.Name(), prop_name)
	return types.NewString(str)
}
