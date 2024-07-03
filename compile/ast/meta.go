package ast

// NodeMeta 节点元数据
type NodeMeta struct {
	RawValue string
	Comments []*Comment
	Others   map[string]any
}

func NewNodeMeta() *NodeMeta {
	return &NodeMeta{
		Others: make(map[string]any),
	}
}

func MetaRawValue(n Node) string {
	if meta := n.Meta(); meta != nil {
		return meta.RawValue
	}
	return ""
}

func MetaComments(n Node) []*Comment {
	if meta := n.Meta(); meta != nil {
		return meta.Comments
	}
	return nil
}
