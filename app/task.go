package app

import (
	"log"

	"github.com/rohitkhatri1st/Task-API/database/psql"
	"github.com/rohitkhatri1st/Task-API/model"
	"github.com/rohitkhatri1st/Task-API/server/config"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Task interface {
	CreateTask(task *model.Task) (*model.Task, error)
	UpdateTask(task *model.Task) (*model.Task, error)
	DeleteTask(taskId uint) error
	GetTask(taskId uint) (*model.Task, error)
	GetAllTasks(pageNo uint) ([]model.Task, error)
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
	err := psql.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatal(err)
	}
	ti := TaskImpl{
		App:    opts.App,
		Logger: &l,
		Db:     psql,
	}
	return &ti
}

func (ti *TaskImpl) CreateTask(task *model.Task) (*model.Task, error) {
	db := ti.Db.Create(task)
	if err := db.Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (ti *TaskImpl) UpdateTask(task *model.Task) (*model.Task, error) {
	db := ti.Db.Where("id = ?", task.ID).Updates(task)
	if err := db.Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (ti *TaskImpl) DeleteTask(taskId uint) error {
	db := ti.Db.Where("id = ?", taskId).Delete(model.Task{})
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

func (ti *TaskImpl) GetTask(taskId uint) (*model.Task, error) {
	task := &model.Task{}
	db := ti.Db.First(task, taskId)
	if err := db.Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (ti *TaskImpl) GetAllTasks(pageNo uint) ([]model.Task, error) {
	tasks := []model.Task{}
	pageSize := 2
	offset := (int(pageNo) - 1) * pageSize
	db := ti.Db.Order("id").Offset(offset).Limit(pageSize).Find(&tasks)
	if err := db.Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
