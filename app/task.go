package app

import (
	"fmt"

	"github.com/rohitkhatri1st/Task-API/database/psql"
	"github.com/rohitkhatri1st/Task-API/model"
	"github.com/rohitkhatri1st/Task-API/server/config"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Task interface {
	SomeTaskLogic()
	CreateTask(task *model.Task) error
}

type TaskImpl struct {
	App    *App
	Logger *zerolog.Logger
	Db     *gorm.DB
}

type TaskImplOpts struct {
	App    *App
	Config *config.ServiceConfig
}

func InitTask(opts *TaskImplOpts) Task {
	l := opts.App.Logger.With().Str("service", "Task").Logger()
	taskDbConfig := config.PsqlConfig{
		DbConfig: opts.App.Config.DatabaseConfig,
		DbName:   opts.Config.DBName,
	}
	psql := psql.InitPsql(&taskDbConfig)
	psql.AutoMigrate(&model.Task{})
	ti := TaskImpl{
		App:    opts.App,
		Logger: &l,
		Db:     psql,
	}
	return &ti
}

func (ti *TaskImpl) SomeTaskLogic() {
	fmt.Println("Implementing task logic")
}

func (ti *TaskImpl) CreateTask(task *model.Task) error {
	ti.Db.Create(task)
	return nil
}
