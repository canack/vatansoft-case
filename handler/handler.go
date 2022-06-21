package handler

// Bu fonksiyonlar, endpoint üzerine gelen istekleri işleyip; veritabanıyla olan iletişimlerine hazır hale getiriyor

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"todo-api/database"
	"todo-api/types"
)

// @Summary      Yeni bir todo oluşturur
// @Description  Bir todo oluşturur ve o todo'ya ait UUIDv4 döner
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param todo body types.CreateTodo true "Todo içeriği"
// @Success      201  {object}  string "UUIDv4"
// @Failure      400  {object}  string
// @Failure      502  {object}  string
// @Router       /todo/ [post]
func CreateTodo(c echo.Context) error {

	todoRequest := new(types.Todo)

	if err := c.Bind(todoRequest); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Gelen istek ayıklanamıyor."))
	}

	todoUUID := uuid.NewString()
	todoRequest.Completed = false
	todoRequest.UUID = todoUUID

	// ...
	// database access
	status, err := database.CreateTodo(todoRequest)

	if err != nil {
		return c.String(http.StatusBadGateway, err.Error())
	}
	// ...

	return c.String(http.StatusCreated, status)
}

// @Summary      Bütün todo'ları döner
// @Description  json array olarak todo'ları döner
// @Tags         todo
// @Accept       json
// @Produce      json
// @Success      200  {object}  []types.Todo
// @Failure      400  {object}  string
// @Router       /todo/ [get]
func GetTodo(c echo.Context) error {
	// ...
	// database access
	todosBytes, err := database.GetTodo()
	if err != nil {
		return c.String(http.StatusBadGateway, err.Error())
	}

	var todos []types.Todo

	unmarshalErr := json.Unmarshal(todosBytes, &todos)
	if unmarshalErr != nil {
		return c.String(http.StatusBadGateway, unmarshalErr.Error())
	}
	// ...

	return c.JSON(http.StatusOK, todos)
}

// @Summary      Verilen UUIDv4 göre todo döner
// @Description  json olarak todo döner
// @Tags         todo
// @Param        UUIDv4 path string true "Todo UUIDv4"
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Todo
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Failure      405  {object}  string
// @Router       /todo/{UUIDv4} [get]
func GetTodoUUID(c echo.Context) error {
	todoUUID := c.Param("todo")

	// ...
	// database access
	todoBytes, err := database.GetTodoUUID(todoUUID)

	if err == gorm.ErrRecordNotFound {
		return c.String(http.StatusNotFound, "Böyle bir todo yok")
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var todo types.Todo

	unmarshallErr := json.Unmarshal(todoBytes, &todo)
	if unmarshallErr != nil {
		return c.String(http.StatusNotAcceptable, unmarshallErr.Error())
	}
	// ...

	return c.JSON(http.StatusOK, todo)
}

// @Summary      Todo içeriğini günceller
// @Description  Todo içeriğini günceller ve boolean döner
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param        UUIDv4 path string true "Todo UUIDv4"
// @Param todo body types.UpdateTodo true "Todo içeriği"
// @Success      200  {object}  bool
// @Failure      400 {object} string
// @Failure      500 {object} string
// @Router       /todo/{UUIDv4} [put]
func UpdateTodo(c echo.Context) error {
	todoUUID := c.Param("todo")

	todoRequest := new(types.UpdateTodo)

	if err := c.Bind(todoRequest); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Gelen istek ayıklanamıyor."))
	}

	// ...
	// database access
	status, err := database.UpdateTodo(todoUUID, todoRequest)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// ...

	return c.String(http.StatusOK, fmt.Sprintf("%v", status))
}

// @Summary      Todo siler
// @Description  Todo siler ve boolean dönüş yapar
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param        UUIDv4 path string true "Todo UUIDv4"
// @Success      200  {object} bool
// @Failure		 500 {object} string
// @Router       /todo/{UUIDv4} [delete]
func DeleteTodo(c echo.Context) error {
	todoUUID := c.Param("todo")

	// ...
	// database access
	status, err := database.DeleteTodo(todoUUID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// ...

	return c.String(http.StatusOK, fmt.Sprintf("%v", status))
}

// @Summary      Todo durumunu günceller
// @Description  Devam eden todo'nun durumunu günceller ve başarılıysa boolean döner
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param        UUIDv4 path string true "Todo UUIDv4"
// @Param todo body types.StatusTodo true "Todo durumu"
// @Success      200  {object} bool
// @Failure		 500 {object} string
// @Router       /todo/{UUIDv4} [patch]
func StatusTodo(c echo.Context) error {
	todoUUID := c.Param("todo")

	todoRequest := new(types.StatusTodo)

	if err := c.Bind(todoRequest); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Gelen istek ayıklanamıyor."))
	}

	// ...
	// database access
	status, err := database.StatusTodo(todoUUID, todoRequest)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// ...

	return c.String(http.StatusOK, fmt.Sprintf("%v", status))
}
