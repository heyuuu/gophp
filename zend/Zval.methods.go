package zend

func (this *Zval) IsNotUndef() bool { return this.GetType() != IS_UNDEF }
