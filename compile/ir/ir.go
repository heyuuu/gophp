package ir

type File struct {
	// 初始化语句
	Init *InitStmt
	// 声明类语句，包括函数、类、接口、Trait等定义
	Decls []Stmt
}
