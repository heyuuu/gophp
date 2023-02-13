package core

import "sik/zend"

func Config() *Configuration {
	return App().config
}

type Configuration struct {
	hash zend.HashTable
}

func (this *Configuration) Init() {
	zend.ZendHashInit(&this.hash, 8, nil, ConfigZvalDtor, 1)
}

func (this *Configuration) Set(key string, value string) *zend.Zval {
	var zv zend.Zval
	zv.SetRawString(value)
	this.hash.KeyUpdate(key, &zv)
}

func (this *Configuration) KeyFind(key string) *zend.Zval {
	return this.hash.KeyFind(key)
}

func (this *Configuration) GetHash() *zend.HashTable {
	return &this.hash
}

func (this *Configuration) Destroy() {
	this.hash.Destroy()
}
