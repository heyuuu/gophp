package types

type blockInfo struct {
	filename   string `prop:""`
	lineStart  uint32 `prop:""`
	lineEnd    uint32 `prop:""`
	docComment string `prop:""` // 块注释，默认值空字符表示注释不存在
}
