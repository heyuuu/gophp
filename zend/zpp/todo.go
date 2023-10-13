package zpp

import "github.com/heyuuu/gophp/zend"

func currExecuteData() ExecuteData { return zend.CurrEX() }

func existException() bool {
	return zend.EG__().HasException()
}
