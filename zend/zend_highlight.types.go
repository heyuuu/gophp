// <<generate>>

package zend

/**
 * ZendSyntaxHighlighterIni
 */
type ZendSyntaxHighlighterIni struct {
	highlight_html    *byte
	highlight_comment *byte
	highlight_default *byte
	highlight_string  *byte
	highlight_keyword *byte
}

func (this ZendSyntaxHighlighterIni) GetHighlightHtml() *byte       { return this.highlight_html }
func (this *ZendSyntaxHighlighterIni) SetHighlightHtml(value *byte) { this.highlight_html = value }
func (this ZendSyntaxHighlighterIni) GetHighlightComment() *byte    { return this.highlight_comment }
func (this *ZendSyntaxHighlighterIni) SetHighlightComment(value *byte) {
	this.highlight_comment = value
}
func (this ZendSyntaxHighlighterIni) GetHighlightDefault() *byte { return this.highlight_default }
func (this *ZendSyntaxHighlighterIni) SetHighlightDefault(value *byte) {
	this.highlight_default = value
}
func (this ZendSyntaxHighlighterIni) GetHighlightString() *byte       { return this.highlight_string }
func (this *ZendSyntaxHighlighterIni) SetHighlightString(value *byte) { this.highlight_string = value }
func (this ZendSyntaxHighlighterIni) GetHighlightKeyword() *byte      { return this.highlight_keyword }
func (this *ZendSyntaxHighlighterIni) SetHighlightKeyword(value *byte) {
	this.highlight_keyword = value
}
