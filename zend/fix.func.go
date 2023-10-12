package zend

import "github.com/heyuuu/gophp/php/types"

func INI_SCNG__() *ZendIniScannerGlobals  { return &IniScannerGlobals }
func LANG_SCNG__() *ZendPhpScannerGlobals { return &LanguageScannerGlobals }
func FC__() *ZendFileContext              { return CG__().GetFileContext() }
func CurrEX() *ZendExecuteData            { return EG__().GetCurrentExecuteData() }

func UninitializedZval() *types.Zval { return types.NewZvalNull() }
