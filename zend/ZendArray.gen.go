package zend

/**
 * 本文件代码由脚本生成，勿手动修改
 */

import (
	b "sik/builtin"
)

func (this *ZendArray) FindBucketByIndex(index int) *Bucket {
	var key_ = NewIndexKey(index)
	return this.FindBucket(key_)
}

func (this *ZendArray) FindBucketByStr(key string) *Bucket {
	var key_ = NewStrKey(key)
	return this.FindBucket(key_)
}

func (this *ZendArray) FindBucketByZendString(key *ZendString) *Bucket {
	var key_ = NewStrKey(key.GetStr())
	return this.FindBucket(key_)
}

func (this *ZendArray) FindBucketByStrPtr(key *byte, len_ int) *Bucket {
	var key_ = NewStrKey(b.CastStr(key, len_))
	return this.FindBucket(key_)
}

func (this *ZendArray) FindByIndex(index int) *Zval {
	var key_ = NewIndexKey(index)
	return this.Find(key_)
}

func (this *ZendArray) FindByStr(key string) *Zval {
	var key_ = NewStrKey(key)
	return this.Find(key_)
}

func (this *ZendArray) FindByZendString(key *ZendString) *Zval {
	var key_ = NewStrKey(key.GetStr())
	return this.Find(key_)
}

func (this *ZendArray) FindByStrPtr(key *byte, len_ int) *Zval {
	var key_ = NewStrKey(b.CastStr(key, len_))
	return this.Find(key_)
}

func (this *ZendArray) ExistsByIndex(index int) bool {
	var key_ = NewIndexKey(index)
	return this.Exists(key_)
}

func (this *ZendArray) ExistsByStr(key string) bool {
	var key_ = NewStrKey(key)
	return this.Exists(key_)
}

func (this *ZendArray) ExistsByZendString(key *ZendString) bool {
	var key_ = NewStrKey(key.GetStr())
	return this.Exists(key_)
}

func (this *ZendArray) ExistsByStrPtr(key *byte, len_ int) bool {
	var key_ = NewStrKey(b.CastStr(key, len_))
	return this.Exists(key_)
}
