package models

import (
	"tomatoList/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"` // 完成状态 0-待完成 1-已完成
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
// CreateATodo 创建todo任务
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// GetATodo
//  @Description: 获取一条待做事项
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateATodo
//  @Description: 更新待做事项
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

// DeleteATodo
//  @Description: 删除待做事项
//  @param id
//  @return err
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
