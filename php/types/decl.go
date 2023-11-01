package types

type blockInfo struct {
	filename   string
	lineStart  uint32
	lineEnd    uint32
	docComment string // 块注释，默认值空字符表示注释不存在
}

func (b *blockInfo) Filename() string                { return b.filename }
func (b *blockInfo) SetFilename(filename string)     { b.filename = filename }
func (b *blockInfo) LineStart() uint32               { return b.lineStart }
func (b *blockInfo) SetLineStart(lineStart uint32)   { b.lineStart = lineStart }
func (b *blockInfo) LineEnd() uint32                 { return b.lineEnd }
func (b *blockInfo) SetLineEnd(lineEnd uint32)       { b.lineEnd = lineEnd }
func (b *blockInfo) DocComment() string              { return b.docComment }
func (b *blockInfo) SetDocComment(docComment string) { b.docComment = docComment }
