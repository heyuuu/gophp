package core

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func Config() *Configuration {
	return App().config
}

type Configuration struct {
	hash types2.Array
}

func (this *Configuration) Init() {
	this.hash.Init(8, ConfigZvalDtor)
}

func (this *Configuration) Set(key string, value string) *types2.Zval {
	var zv types2.Zval
	zv.SetStringVal(value)
	this.hash.KeyUpdate(key, &zv)
}

func (this *Configuration) KeyFind(key string) *types2.Zval {
	return this.hash.KeyFind(key)
}

func (this *Configuration) GetHash() *types2.Array {
	return &this.hash
}

func (this *Configuration) Destroy() {
	this.hash.Destroy()
}
