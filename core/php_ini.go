package core

import (
	"sik/zend/types"
)

func Config() *Configuration {
	return App().config
}

type Configuration struct {
	hash types.Array
}

func (this *Configuration) Init() {
	&this.hash = types.MakeArrayEx(8, ConfigZvalDtor, 1)
}

func (this *Configuration) Set(key string, value string) *types.Zval {
	var zv types.Zval
	zv.SetRawString(value)
	this.hash.KeyUpdate(key, &zv)
}

func (this *Configuration) KeyFind(key string) *types.Zval {
	return this.hash.KeyFind(key)
}

func (this *Configuration) GetHash() *types.Array {
	return &this.hash
}

func (this *Configuration) Destroy() {
	this.hash.Destroy()
}
