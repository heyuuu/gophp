package zend

/**
 * 本文件代码由脚本生成，勿手动修改
 */

import (
	b "sik/builtin"
)

func (this *ZendArray) FindByZendString(key *ZendString) *Zval {
	return this.KeyFind(key.GetStr())
}

func (this *ZendArray) FindByStrPtr(key *byte, len_ int) *Zval {
	return this.KeyFind(b.CastStr(key, len_))
}
