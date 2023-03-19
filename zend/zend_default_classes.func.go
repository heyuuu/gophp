// <<generate>>

package zend

import "sik/zend/faults"

func ZendRegisterDefaultClasses() {
	ZendRegisterInterfaces()
	faults.ZendRegisterDefaultException()
	ZendRegisterIteratorWrapper()
	ZendRegisterClosureCe()
	ZendRegisterGeneratorCe()
	ZendRegisterWeakrefCe()
}
