package service

import (
	"github.com/mytheta/kanbanAPI/model"
)

var Todo = todoimpl{}

type todoimpl struct {}

func(t *todoimpl) FindByStatus(status string) ([]model.Todo,error){
	var todos []model.Todo
	err := db.Where("status = ?", status).Find(&todos).Error
	return todos, err
}