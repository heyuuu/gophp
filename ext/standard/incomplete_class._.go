package standard

import (
	"sik/zend"
)

const INCOMPLETE_CLASS_MSG string = "The script tried to execute a method or " + "access a property of an incomplete object. " + "Please ensure that the class definition \"%s\" of the object " + "you are trying to operate on was loaded _before_ " + "unserialize() gets called or provide an autoloader " + "to load the class definition"

var PhpIncompleteObjectHandlers zend.ZendObjectHandlers
