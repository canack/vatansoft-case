package types

// Bu dosya, sadece kullanılan tipleri içeriyor

import "time"

type Todo struct {
	UUID      string    `json:"uuid" example:"6c9c1545-8631-4b01-9ae1-8b8d11acd028"`
	Title     string    `json:"title" example:"Günlük antrenman"`
	Content   string    `json:"content" example:"30 Dakika koşu, 30 dakika meditasyon"`
	CreatedAt time.Time `json:"created_at" example:"2022-06-21T19:37:56+03:00" gorm:"type:time"`
	UpdatedAt time.Time `json:"updated_at" example:"2022-06-21T19:37:56+03:00" gorm:"type:time"`
	Completed bool      `json:"completed" example:"true"`
	Priority  uint8     `json:"priority" example:"10"`
}

type CreateTodo struct {
	Title    string `json:"title" example:"Egzersiz planım"`
	Content  string `json:"content" example:"Bugün 30 dakika koşu ve 30 dakika esneme yapacağım"`
	Priority uint8  `json:"priority" example:"3"`
}

type UpdateTodo struct {
	Title    string `json:"title" example:"Değiştirilecek örnek bir başlık"`
	Content  string `json:"content" example:"Değiştirilecek örnek bir içerik"`
	Priority uint8  `json:"priority" example:"5"`
}

type StatusTodo struct {
	Completed bool `json:"completed" example:"true"`
}
