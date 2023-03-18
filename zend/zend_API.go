package zend

// todo 使用 ModuleRegistryMap 替换 ModuleRegistry
var ModuleRegistryMap = make(map[string]*ZendModuleEntry)

func CopyRegistryModules() []*ZendModuleEntry {
	var modules = make([]*ZendModuleEntry, 0, len(ModuleRegistryMap))
	for _, module := range ModuleRegistryMap {
		modules = append(modules, module)
	}
	return modules
}
