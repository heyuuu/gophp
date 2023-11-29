package php

import (
	"github.com/heyuuu/gophp/php/types"
)

type Val = *types.Zval

var NewVal = types.NewZvalUndef
var Null = types.NewZvalNull
var False = types.NewZvalFalse
var True = types.NewZvalTrue
var Bool = types.NewZvalBool
var Long = types.NewZvalLong
var Double = types.NewZvalDouble
var String = types.NewZvalString
var Array = types.NewZvalArray
