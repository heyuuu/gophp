package php

import "github.com/heyuuu/gophp/php/perr"

var isBoot = false

func MarkIsBoot(v bool) {
	isBoot = v
}

func CheckIsBoot() {
	if !isBoot {
		perr.Panic(`请在使用 gophp 前初始化 (import _ "github.com/heyuuu/gophp/php/boot")`)
	}
}
