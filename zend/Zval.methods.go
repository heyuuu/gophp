package zend

import b "sik/builtin"

func (this *Zval) IsRefcounted() bool  { return this.GetTypeFlags() != 0 }
func (this *Zval) IsCollectable() bool { return b.FlagMatch(this.GetTypeFlags(), IS_TYPE_COLLECTABLE) }
