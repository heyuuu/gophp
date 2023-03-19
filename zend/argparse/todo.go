package argparse

import "sik/zend"

func currExecuteData() ExecuteData { return zend.CurrEX() }

func isArgUseWeakTypes() bool { return !zend.CurrEX().IsArgUseStrictTypes() }
