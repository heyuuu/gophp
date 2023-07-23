package ir

// Namespace
type Namespace struct {
	Name string
	// 代码段
	Segments []Segment
}

func NewNamespace(name string) *Namespace {
	return &Namespace{Name: name}
}

// Project
type Project struct {
	namespaces map[string]*Namespace
}

func NewProject() *Project {
	return &Project{
		namespaces: make(map[string]*Namespace),
	}
}

func (p *Project) AddFile(fileName string, f *File) error {
	// todo 临时开发调试代码，待修正
	for _, segment := range f.Segments {
		if p.namespaces[segment.Namespace] == nil {
			p.namespaces[segment.Namespace] = NewNamespace(segment.Namespace)
		}

		ns := p.namespaces[segment.Namespace]
		ns.Segments = append(ns.Segments, segment)
	}

	return nil
}

func (p *Project) GetConstant(name string) {
	// todo
}

func (p *Project) GetFunction(name string) {
	// todo
}

func (p *Project) GetClass(name string) {
	// todo
}
