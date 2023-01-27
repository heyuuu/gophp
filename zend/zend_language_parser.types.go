// <<generate>>

package zend

/**
 * LangUnionYyalloc
 */
type LangUnionYyalloc struct /* union */ {
	yyss_alloc yytype_int16
	yyvs_alloc ZendParserStackElem
}

func (this *LangUnionYyalloc) GetYyssAlloc() yytype_int16             { return this.yyss_alloc }
func (this *LangUnionYyalloc) SetYyssAlloc(value yytype_int16)        { this.yyss_alloc = value }
func (this *LangUnionYyalloc) GetYyvsAlloc() ZendParserStackElem      { return this.yyvs_alloc }
func (this *LangUnionYyalloc) SetYyvsAlloc(value ZendParserStackElem) { this.yyvs_alloc = value }
