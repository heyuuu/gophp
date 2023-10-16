package core

import (
	"github.com/heyuuu/gophp/zend"
)

/**
 * PostVarDataT
 */
type PostVarDataT struct {
	str            zend.SmartStr
	ptr            *byte
	end            *byte
	cnt            uint64
	alreadyScanned int
}

func (d *PostVarDataT) GetStr() *zend.SmartStr      { return &d.str }
func (d *PostVarDataT) GetPtr() *byte               { return d.ptr }
func (d *PostVarDataT) SetPtr(value *byte)          { d.ptr = value }
func (d *PostVarDataT) GetEnd() *byte               { return d.end }
func (d *PostVarDataT) SetEnd(value *byte)          { d.end = value }
func (d *PostVarDataT) GetCnt() uint64              { return d.cnt }
func (d *PostVarDataT) GetAlreadyScanned() int      { return d.alreadyScanned }
func (d *PostVarDataT) SetAlreadyScanned(value int) { d.alreadyScanned = value }
