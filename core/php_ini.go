package core

import (
	"sik/zend"
	"sik/zend/types"
)

func Config() *Configuration {
	return App().config
}

type Configuration struct {
	hash types.HashTable
}

func (this *Configuration) Init() {
	zend.ZendHashInit(&this.hash, 8, nil, ConfigZvalDtor, 1)
}

func (this *Configuration) Set(key string, value string) *types.Zval {
	var zv types.Zval
	zv.SetRawString(value)
	this.hash.KeyUpdate(key, &zv)
}

func (this *Configuration) KeyFind(key string) *types.Zval {
	return this.hash.KeyFind(key)
}

func (this *Configuration) GetHash() *types.HashTable {
	return &this.hash
}

func (this *Configuration) Destroy() {
	this.hash.Destroy()
}
