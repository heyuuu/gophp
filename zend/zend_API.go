package zend

import "log"

// todo 使用 ModuleRegistryMap 替换 ModuleRegistry
var ModuleRegistryMap = make(map[string]*ZendModuleEntry)

func CopyRegistryModules() []*ZendModuleEntry {
	var modules = make([]*ZendModuleEntry, 0, len(ModuleRegistryMap))
	for _, module := range ModuleRegistryMap {
		modules = append(modules, module)
	}
	return modules
}

/**
 *
 */
type receiveArgs struct {
	args []any
	pos  int
}

func newReceiveArgs(args []any) *receiveArgs {
	return &receiveArgs{args: args, pos: 0}
}
func putReceiveArg[T any](r *receiveArgs, val T) {
	if r.pos >= len(r.args) {
		log.Fatal("解析参数异常，超过获取长度")
	}

	if ptr, ok := r.args[r.pos].(*T); ok {
		*ptr = val
	} else {
		log.Fatalf("解析参数异常: 类型不匹配，pos=%d", r.pos)
	}

	r.pos++
}
