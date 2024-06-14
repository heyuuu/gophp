package ast

func MetaRawValue(n Node) string {
	if rawValue, ok := n.MetaValue("rawValue").(string); ok {
		return rawValue
	}
	return ""
}
