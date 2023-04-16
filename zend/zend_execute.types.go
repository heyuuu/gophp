package zend

import (
	b "github.com/heyuuu/gophp/builtin"
)

type VmStack struct {
	elements []*ZendExecuteData
}

func (s *VmStack) Reset() { s.elements = nil }

func (s *VmStack) Push(ex *ZendExecuteData) {
	s.elements = append(s.elements, ex)
}

func (s *VmStack) Pop() *ZendExecuteData {
	if len(s.elements) == 0 {
		return nil
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top
}

func (s *VmStack) PopCheck(ex *ZendExecuteData) {
	b.Assert(s.Pop() == ex)
}
