package core

/**
 * App PHP应用服务实体，全局唯一
 */
type App struct {
	sapiModule  *SapiModule
	sapiGlobals SapiGlobals
}

var currentApp *App

func CurrentApp() *App {
	// todo 获取当前 App，后续需替换掉
	return currentApp
}

func NewApp() *App {
	return &App{}
}

func (app *App) Startup(module *SapiModule) {
	module.SetIniEntries(nil)
	app.sapiModule = module
	app.sapiGlobals.Init()
}

func (app *App) Shutdown() {
	app.sapiGlobals.Destroy()
}

func (app *App) SG() *SapiGlobals { return &app.sapiGlobals }
func (app *App) SM() *SapiModule  { return app.sapiModule }

/**
 * Context 单个PHP请求上下文
 */
type Context struct {
	app *App
}
