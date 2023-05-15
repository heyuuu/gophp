package core

import (
	"github.com/heyuuu/gophp/php/types"
)

func Config() *Configuration {
	return App().config
}

type Configuration struct {
	hash *types.Array
}

func (this *Configuration) Init() {
	this.hash = types.NewArray(0)
}

func (this *Configuration) Set(key string, value string) {
	var zv types.Zval
	zv.SetStringVal(value)
	this.hash.KeyUpdate(key, &zv)
}

func (this *Configuration) KeyFind(key string) *types.Zval {
	return this.hash.KeyFind(key)
}

func (this *Configuration) GetHash() *types.Array {
	return this.hash
}

func (this *Configuration) Destroy() {
	this.hash.Destroy()
}
