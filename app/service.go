package app

// InitService this initializes all the busines logic services
func InitService(a *App) {

	a.Task = InitTask(&TaskImplOpts{
		App:    a,
		Config: &a.Config.TaskConfig,
	})
	a.Utility = InitUtility(&UtilityOpts{
		App:    a,
		Logger: a.Logger,
	})
}
