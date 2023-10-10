package zend

import (
	b "github.com/heyuuu/gophp/builtin"
)

/**
 * ZendStack
 */
type ZendStack[T any] struct {
	elements []T
}

const ZEND_STACK_APPLY_TOPDOWN = 1
const ZEND_STACK_APPLY_BOTTOMUP = 2

func (this *ZendStack[T]) Init() {
	this.elements = nil
}

func (this *ZendStack[T]) Slice(len_ int) {
	b.Assert(len_ <= len(this.elements))
	if len_ == 0 {
		this.elements = nil
	} else {
		this.elements = this.elements[:len_]
	}
}

func (this *ZendStack[T]) Push(element T) int {
	this.elements = append(this.elements, element)
	return len(this.elements)
}

func (this *ZendStack[T]) Pop() T {
	var top = this.Top()

	var len_ = len(this.elements)
	if len_ > 0 {
		this.elements = this.elements[:len_-1]
	}

	return top
}

func (this *ZendStack[T]) Top() T {
	var lastIndex = len(this.elements) - 1
	if lastIndex >= 0 {
		return this.elements[lastIndex]
	}
	return nil
}

func (this *ZendStack[T]) DelTop() {
	this.Pop()
}

func (this *ZendStack[T]) Size() int {
	return len(this.elements)
}

func (this *ZendStack[T]) IsEmpty() bool {
	return len(this.elements) == 0
}

func (this *ZendStack[T]) Destroy() {
	this.Init()
}

func (this *ZendStack[T]) ApplyWithArgument(type_ int, apply_function func(element any, arg any) int, arg any) {
	var len_ = len(this.elements)
	switch type_ {
	case ZEND_STACK_APPLY_TOPDOWN:
		for i := len_ - 1; i >= 0; i-- {
			element := this.elements[i]
			if apply_function(element, arg) != 0 {
				break
			}
		}
	case ZEND_STACK_APPLY_BOTTOMUP:
		for i := 0; i < len_; i++ {
			element := this.elements[i]
			if apply_function(element, arg) != 0 {
				break
			}
		}
	}
}

func (this *ZendStack[T]) Clean(func_ func(any), freeElements bool) {
	if func_ != nil {
		for _, element := range this.elements {
			func_(element)
		}
	}
	if freeElements {
		this.Init()
	}
}

/**
 * <<generate>>
 */
func (this *ZendStack[T]) GetTop() int       { return len(this.elements) }
func (this *ZendStack[T]) SetTop(offset int) { this.Slice(offset) }
func (this *ZendStack[T]) GetElements() []T  { return this.elements }
