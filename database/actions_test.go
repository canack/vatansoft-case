package database

// Bu dosya üzerinde, veritabanı bağlantılarının manuel testlerini yazdım.

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"testing"
	"time"
	"todo-api/types"
)

func TestCreateTodo(t *testing.T) {
	ConnectDB()
	for i := 0; i < 10; i++ {
		now := time.Now()
		todo := types.Todo{UUID: uuid.NewString(), CreatedAt: now, UpdatedAt: now,
			Title: "", Content: "", Completed: false, Priority: 0}

		CreateTodo(&todo)

	}

}

func TestDeleteTodo(t *testing.T) {
	ConnectDB()
	status, err := DeleteTodo("6c9c1545-8631-4b01-9ae1-8b8d11acd028")
	if err != nil {
		panic(err)
	}

	fmt.Println(status)

}

func TestStatusTodo(t *testing.T) {
	ConnectDB()

	statusTodo := types.StatusTodo{Completed: true}
	status, err := StatusTodo("d988a769-4fbe-4127-bcd3-e64f7fc31d83", &statusTodo)
	if err != nil {
		panic(err)
	}

	fmt.Println(status)

}

func TestUpdateTodo(t *testing.T) {
	ConnectDB()

	updateTodo := types.UpdateTodo{Title: "Sevqwvqewlam", Content: "werwer", Priority: 4}

	status, err := UpdateTodo("bb1ec4cb-919f-4380-a43e-9a9303db3b8e", &updateTodo)

	if err != nil {
		panic(err)
	}

	fmt.Println(status)

}

func TestGetTodoUUID(t *testing.T) {
	ConnectDB()

	status, err := GetTodoUUID("bb1ec4cb-919f-4380-a43e-9a9303db3b8e")

	if err == gorm.ErrRecordNotFound {
		fmt.Println("Böyle bir todo yok")
		return
	}
	if err != nil {
		panic(err)
	}

	var todo types.Todo

	json.Unmarshal(status, &todo)

	fmt.Println(todo)
}

func TestGetTodo(t *testing.T) {
	ConnectDB()

	status, err := GetTodo()

	if err != nil {
		panic(err)
	}

	var todos []types.Todo

	json.Unmarshal(status, &todos)

	fmt.Println(todos)
}
