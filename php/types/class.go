package types

type Class struct {
	name string
	blockInfo
}

func (c *Class) Name() string { return c.name }
