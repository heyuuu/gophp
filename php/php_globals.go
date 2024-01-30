package php

// PhpCoreGlobals
type PhpCoreGlobals struct {
	serializePrecision int
}

func (p *PhpCoreGlobals) Init() {
	p.serializePrecision = -1
}

func (p *PhpCoreGlobals) SerializePrecision() int {
	return p.serializePrecision
}
func (p *PhpCoreGlobals) SetSerializePrecision(serializePrecision int) {
	p.serializePrecision = serializePrecision
}
