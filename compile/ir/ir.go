package ir

// File
type File struct {
	// 文件是否开启 strict_types
	StrictTypes bool
	// 按命名空间保存的结果
	Namespaces []*Namespace
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

func (p *Project) getNamespaceOrInit(name string) *Namespace {
	if p.namespaces[name] == nil {
		p.namespaces[name] = NewNamespace(name)
	}
	return p.namespaces[name]
}

func (p *Project) AddFile(fileName string, f *File) (err error) {
	for _, ns := range f.Namespaces {
		err = p.getNamespaceOrInit(ns.Name).merge(ns)
		if err != nil {
			return err
		}
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

// Namespace
type Namespace struct {
	Name string
	// 代码段
	Segments []Segment
	// maps
	funcMap      map[string]*Func
	classLikeMap map[string]ClassLike
}

func NewNamespace(name string) *Namespace {
	return &Namespace{Name: name}
}

func (ns *Namespace) merge(newNs *Namespace) error {
	if ns.Name != newNs.Name {
		// todo use error
		panic("*Namespace.merge() must has same name")
	}
	// check name conflict
	for _, segment := range newNs.Segments {
		switch s := segment.(type) {
		case ClassLike:
			_, exists := ns.classLikeMap[s.GetName()]
			if exists {
				// todo use error
				panic("ClassLike conflict:" + s.GetName())
			}
			ns.classLikeMap[s.GetName()] = s
		case *Func:
			_, exists := ns.classLikeMap[s.Name.ToCodeString()]
			if exists {
				// todo use error
				panic("ClassLike conflict:" + s.GetName())
			}
			ns.funcMap[s.GetName()] = s
		}
	}
	ns.Segments = append(ns.Segments, newNs.Segments...)

	return nil
}

// Segment
type (
	Segment interface{ segment() }
	Decl    interface {
		Segment
		GetName() string
		decl()
	}
	ClassLike interface {
		Decl
		classLike()
	}

	InitFunc struct {
		Stmts []Stmt
	}

	Func struct {
		Name       Name
		ByRef      bool     // @var bool Whether function returns by reference
		Params     []*Param // @var Param[] Parameters
		ReturnType Type     // @var Type|null Return type
		Stmts      []Stmt   // @var Stmt[] Statements
	}

	Class struct {
		Name       Name
		Flags      Flags
		Extends    *Name
		Implements []Name
		Stmts      []Stmt
	}

	Interface struct {
		Name    Name
		Extends []Name
		Stmts   []Stmt
	}

	Trait struct {
		Name  Name
		Stmts []Stmt
	}
)

func (s *InitFunc) segment()  {}
func (s *Func) segment()      {}
func (s *Class) segment()     {}
func (s *Interface) segment() {}
func (s *Trait) segment()     {}

func (s *Func) decl()      {}
func (s *Class) decl()     {}
func (s *Interface) decl() {}
func (s *Trait) decl()     {}

func (s *Class) classLike()     {}
func (s *Interface) classLike() {}
func (s *Trait) classLike()     {}

func (s *Func) GetName() string      { return s.Name.ToCodeString() }
func (s *Class) GetName() string     { return s.Name.ToCodeString() }
func (s *Interface) GetName() string { return s.Name.ToCodeString() }
func (s *Trait) GetName() string     { return s.Name.ToCodeString() }
