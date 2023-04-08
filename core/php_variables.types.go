package core

import (
	"github.com/heyuuu/gophp/zend"
)

/**
 * PostVarDataT
 */
type PostVarDataT struct {
	str             zend.SmartStr
	ptr             *byte
	end             *byte
	cnt             uint64
	already_scanned int
}

func (this *PostVarDataT) GetStr() zend.SmartStr       { return this.str }
func (this *PostVarDataT) GetPtr() *byte               { return this.ptr }
func (this *PostVarDataT) SetPtr(value *byte)          { this.ptr = value }
func (this *PostVarDataT) GetEnd() *byte               { return this.end }
func (this *PostVarDataT) SetEnd(value *byte)          { this.end = value }
func (this *PostVarDataT) GetCnt() uint64              { return this.cnt }
func (this *PostVarDataT) GetAlreadyScanned() int      { return this.already_scanned }
func (this *PostVarDataT) SetAlreadyScanned(value int) { this.already_scanned = value }
