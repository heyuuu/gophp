package zend

func (this *Zval) IsNotUndef() bool { return this.IsType(IS_UNDEF) }
