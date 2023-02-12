package core

/**
 * App PHP应用服务实体，全局唯一
 */
type App struct {
	sapiModule  ISapiModule
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

func (app *App) Startup(module ISapiModule) {
	sf := module.(*SapiModule)
	sf.SetIniEntries(nil)
	sapi_module = *sf
	app.SG().Init()
}

func (app *App) Shutdown() {
	app.SG().Destroy()
}

func (app *App) SG() *SapiGlobals { return &app.sapiGlobals }

/**
 * Context 单个PHP请求上下文
 */
type Context struct {
	app *App
}
