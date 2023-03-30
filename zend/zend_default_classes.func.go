package zend

import "github.com/heyuuu/gophp/zend/faults"

func ZendRegisterDefaultClasses() {
	ZendRegisterInterfaces()
	faults.RegisterDefaultException()
	ZendRegisterIteratorWrapper()
	ZendRegisterClosureCe()
	ZendRegisterGeneratorCe()
	ZendRegisterWeakrefCe()
}
