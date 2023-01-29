package zend

import b "sik/builtin"

func (this *Zval) IsNotUndef() bool { return this.IsType(IS_UNDEF) }

func (this *Zval) IsRefcounted() bool { return b.FlagMatch(this.GetTypeInfo(), Z_TYPE_FLAGS_MASK) }
