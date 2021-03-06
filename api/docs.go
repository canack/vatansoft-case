// Package api GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/todo/": {
            "get": {
                "description": "json array olarak todo'ları döner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Bütün todo'ları döner",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Todo"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Bir todo oluşturur ve o todo'ya ait UUIDv4 döner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Yeni bir todo oluşturur",
                "parameters": [
                    {
                        "description": "Todo içeriği",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateTodo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "UUIDv4",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/todo/{UUIDv4}": {
            "get": {
                "description": "json olarak todo döner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Verilen UUIDv4 göre todo döner",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo UUIDv4",
                        "name": "UUIDv4",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Todo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Todo içeriğini günceller ve boolean döner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Todo içeriğini günceller",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo UUIDv4",
                        "name": "UUIDv4",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Todo içeriği",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UpdateTodo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Todo siler ve boolean dönüş yapar",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Todo siler",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo UUIDv4",
                        "name": "UUIDv4",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Devam eden todo'nun durumunu günceller ve başarılıysa boolean döner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Todo durumunu günceller",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo UUIDv4",
                        "name": "UUIDv4",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Todo durumu",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.StatusTodo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.CreateTodo": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "example": "Bugün 30 dakika koşu ve 30 dakika esneme yapacağım"
                },
                "priority": {
                    "type": "integer",
                    "example": 3
                },
                "title": {
                    "type": "string",
                    "example": "Egzersiz planım"
                }
            }
        },
        "types.StatusTodo": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "types.Todo": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": true
                },
                "content": {
                    "type": "string",
                    "example": "30 Dakika koşu, 30 dakika meditasyon"
                },
                "created_at": {
                    "type": "string",
                    "example": "2022-06-21T19:37:56+03:00"
                },
                "priority": {
                    "type": "integer",
                    "example": 10
                },
                "title": {
                    "type": "string",
                    "example": "Günlük antrenman"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2022-06-21T19:37:56+03:00"
                },
                "uuid": {
                    "type": "string",
                    "example": "6c9c1545-8631-4b01-9ae1-8b8d11acd028"
                }
            }
        },
        "types.UpdateTodo": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "example": "Değiştirilecek örnek bir içerik"
                },
                "priority": {
                    "type": "integer",
                    "example": 5
                },
                "title": {
                    "type": "string",
                    "example": "Değiştirilecek örnek bir başlık"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Todo api",
	Description:      "Vatansoft tarafından gönderilen case için yazılmış interaktif talimatlar bütünü\nTodo api kullanım talimatlarını içerir.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
