package database

// Bu dosyada, veritabanıyla yapılan iletişim kodlandı
// Buradaki tüm fonksiyonlar sadece veritabanıyla ilgili

import (
	"encoding/json"
	"todo-api/types"
)

func CreateTodo(todo *types.Todo) (string, error) {

	tx := db.Create(&todo)

	err := tx.Error
	if err != nil {
		return "", err
	}

	return todo.UUID, nil
}

func GetTodo() ([]byte, error) {

	var todos []types.Todo

	tx := db.Find(&todos)

	err := tx.Error
	if err != nil {
		return []byte{}, err
	}

	todosBytes, err := json.Marshal(&todos)
	if err != nil {
		return []byte{}, err
	}

	return todosBytes, nil

}

func GetTodoUUID(uuid string) ([]byte, error) {

	var todo types.Todo // := types.Todo{}
	tx := db.First(&todo, "uuid = ?", uuid)
	err := tx.Error

	if err != nil {
		return []byte{}, err
	}

	todoBytes, err := json.Marshal(&todo)
	if err != nil {
		return []byte{}, err
	}

	return todoBytes, nil
}

func UpdateTodo(uuid string, todoUpdate *types.UpdateTodo) (bool, error) {

	tx := db.Model(types.Todo{}).Where("uuid = ?", uuid).Updates(map[string]interface{}{
		"title":    todoUpdate.Title,
		"content":  todoUpdate.Content,
		"priority": todoUpdate.Priority,
	})

	err := tx.Error
	if err != nil {
		return false, err
	}

	affected := tx.RowsAffected

	if affected > 0 {
		return true, nil
	}

	return false, nil
}

func DeleteTodo(uuid string) (bool, error) {

	tx := db.Where("uuid = ?", uuid).Delete(types.Todo{})

	err := tx.Error
	if err != nil {
		return false, err
	}

	affected := tx.RowsAffected

	if affected > 0 {
		return true, nil
	}

	return false, nil
}

func StatusTodo(uuid string, todoStatus *types.StatusTodo) (bool, error) {

	tx := db.Model(types.Todo{}).Where("uuid = ?", uuid).Update("completed", todoStatus.Completed)
	err := tx.Error

	if err != nil {
		return false, err
	}

	affected := tx.RowsAffected
	if affected > 0 {
		return true, nil
	}

	return false, nil
}
