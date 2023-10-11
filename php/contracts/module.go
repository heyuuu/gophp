package contracts

type ModuleLifeCycle interface {
	StartUp()    // call on module startup
	Shutdown()   // call on module shutdown
	Activate()   // call on request startup
	Deactivate() // call on request shutdown
}
