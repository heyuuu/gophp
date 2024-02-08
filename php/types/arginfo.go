package types

// ArgInfo
type ArgInfo struct {
	name     string `get:""`
	byRef    bool   `get:""`
	variadic bool   `get:""`
}

func MakeArgInfo(name string, byRef bool, variadic bool) ArgInfo {
	return ArgInfo{
		name:     name,
		byRef:    byRef,
		variadic: variadic,
	}
}
