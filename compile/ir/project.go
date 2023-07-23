package ir

// Namespace
type Namespace struct{}

// Project
type Project struct {
	namespaces map[string]*Namespace
}

func NewProject() *Project {
	// todo
	return nil
}

func (p *Project) AddFile(fileName string, f *File) error {
	// todo
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
