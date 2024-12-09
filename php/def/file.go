package def

type File struct {
	FilePath    string
	StrictTypes bool
}

func NewFile(filePath string, strictTypes bool) *File {
	return &File{
		FilePath:    filePath,
		StrictTypes: strictTypes,
	}
}

func (f *File) TopFn(namespace string, fn func(d TopDefiner) Val) {
	// TODO
}
