package argparse

import "sik/zend"

func currExecuteData() ExecuteData { return zend.CurrEX() }

func existException() bool {
	return zend.EG__().GetException() != nil
}
