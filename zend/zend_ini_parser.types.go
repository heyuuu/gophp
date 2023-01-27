// <<generate>>

package zend

/**
 * IniUnionYyalloc
 */
type IniUnionYyalloc struct /* union */ {
	yyss_alloc yytype_int16
	yyvs_alloc Zval
}

func (this *IniUnionYyalloc) GetYyssAlloc() yytype_int16      { return this.yyss_alloc }
func (this *IniUnionYyalloc) SetYyssAlloc(value yytype_int16) { this.yyss_alloc = value }
func (this *IniUnionYyalloc) GetYyvsAlloc() Zval              { return this.yyvs_alloc }
func (this *IniUnionYyalloc) SetYyvsAlloc(value Zval)         { this.yyvs_alloc = value }
