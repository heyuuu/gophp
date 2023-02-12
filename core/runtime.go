package core

/**
 * App PHP应用服务实体，全局唯一
 */
type App struct {
	sapiModule  ISapiModule
	sapiGlobals SapiGlobals
}

func NewApp() *App {
	return &App{}
}

func (app *App) Startup(module ISapiModule) {
	sf := module.(*SapiModule)
	sf.SetIniEntries(nil)
	sapi_module = *sf
	SapiGlobalsCtor(&sapi_globals)
}

func (app *App) Shutdown() {
	SapiGlobalsDtor(&sapi_globals)
}

/**
 * Context 单个PHP请求上下文
 */
type Context struct {
	app *App
}
