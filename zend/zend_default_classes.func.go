// <<generate>>

package zend

import "sik/zend/faults"

func ZendRegisterDefaultClasses() {
	ZendRegisterInterfaces()
	faults.RegisterDefaultException()
	ZendRegisterIteratorWrapper()
	ZendRegisterClosureCe()
	ZendRegisterGeneratorCe()
	ZendRegisterWeakrefCe()
}
