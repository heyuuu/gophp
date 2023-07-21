package ir

type File struct {
	// 文件是否开启 strict_types
	StrictTypes bool
	//// 初始化语句
	//Init *InitStmt
	//// 声明类语句，包括函数、类、接口、Trait等定义
	//Decls []Stmt
	// 代码段
	Segments []Segment
}

type Segment struct {
	// 段命名空间
	Namespace string
	// 初始化语句
	Init *InitStmt
	// 声明类语句，包括函数、类、接口、Trait等定义
	Decls []Stmt
}
