package core

/**
 * Application PHP应用服务实体，全局唯一
 */
type Application struct {
	sapiModule  ISapiModule
	sapiGlobals SapiGlobals
	config      *Configuration
}

var currentApp *Application

func App() *Application {
	// todo 获取当前 Application，后续需替换掉
	return currentApp
}

func NewApp() *Application {
	return &Application{}
}

func (app *Application) Startup(module ISapiModule) {
	module.SetIniEntries(nil)
	app.sapiModule = module
	app.sapiGlobals.Init()
}

func (app *Application) Shutdown() {
	app.sapiGlobals.Destroy()
}

func (app *Application) SG() *SapiGlobals { return &app.sapiGlobals }
func (app *Application) SM() ISapiModule  { return app.sapiModule }

/**
 * Context 单个PHP请求上下文
 */
type Context struct {
	app *Application
}
