package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

func SplRegisterProperty(classEntry *types.ClassEntry, propName string, propNameLen int, propFlags int) {
	zend.ZendDeclarePropertyNull(classEntry, propName, propNameLen, propFlags)
}
func SplAddClassName(list *types.Array, pce *types.ClassEntry, allow int, ceFlags uint32) {
	if allow == 0 || allow > 0 && pce.HasCeFlags(ceFlags) || allow < 0 && !pce.HasCeFlags(ceFlags) {
		if !list.KeyExists(pce.Name()) {
			list.KeyAdd(pce.Name(), types.NewZvalString(pce.Name()))
		}
	}
}
func SplAddInterfaces(list *types.Array, pce *types.ClassEntry, allow int, ceFlags uint32) {
	if pce.GetNumInterfaces() != 0 {
		b.Assert(pce.HasCeFlags(types.AccLinked))
		for numInterfaces := 0; numInterfaces < pce.GetNumInterfaces(); numInterfaces++ {
			SplAddClassName(list, pce.GetInterfaces()[numInterfaces], allow, ceFlags)
		}
	}
}
func SplAddTraits(list *types.Array, pce *types.ClassEntry, allow int, ceFlags uint32) {
	var numTraits uint32
	var trait *types.ClassEntry
	for numTraits = 0; numTraits < pce.GetNumTraits(); numTraits++ {
		trait = zend.ZendFetchClassByName_Ex(pce.GetTraitNames()[numTraits].GetName(), pce.GetTraitNames()[numTraits].GetLcName(), zend.ZEND_FETCH_CLASS_TRAIT)
		b.Assert(trait != nil)
		SplAddClassName(list, trait, allow, ceFlags)
	}
}

func SplGenPrivatePropName(ce *types.ClassEntry, prop_name string) *types.String {
	str := zend.ZendManglePropertyName_Ex(ce.Name(), prop_name)
	return types.NewString(str)
}
