// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

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
        "/metrics": {
            "get": {
                "description": "Получение метрик по фильтрам (временной интервал, ID устройства и т. д.).\nОБЯЗАТЕЛЬНО юзать device_id и interval",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "Получить метрики",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"minute\"",
                        "description": "Интервал ('minute hour day week')",
                        "name": "interval",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "example": 100500,
                        "description": "ID устройства",
                        "name": "device_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "2006-01-02T15:04:05Z",
                        "description": "Дата начала (time.RFC3339)",
                        "name": "from_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2006-01-02T15:04:05Z",
                        "description": "Дата окончания (time.RFC3339)",
                        "name": "to_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ",
                        "schema": {
                            "$ref": "#/definitions/domains.SuccessGet"
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации запроса",
                        "schema": {
                            "$ref": "#/definitions/domains.ErrorBody"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/domains.ErrorBody"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domains.ErrorBody": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "domains.LoadCell": {
            "type": "object",
            "properties": {
                "output1": {
                    "type": "number"
                },
                "output2": {
                    "type": "number"
                }
            }
        },
        "domains.Metrics": {
            "type": "object",
            "required": [
                "device_id"
            ],
            "properties": {
                "device_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "load_cell": {
                    "description": "Тензодатчик",
                    "allOf": [
                        {
                            "$ref": "#/definitions/domains.LoadCell"
                        }
                    ]
                },
                "muscle_activity": {
                    "$ref": "#/definitions/domains.MuscleActivity"
                },
                "pulse": {
                    "type": "number"
                },
                "temperature": {
                    "type": "number"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        },
        "domains.MuscleActivity": {
            "type": "object",
            "properties": {
                "output1": {
                    "type": "number"
                },
                "output2": {
                    "type": "number"
                }
            }
        },
        "domains.SuccessGet": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domains.Metrics"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization.",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Vet clinic metrics service",
	Description:      "metrics service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
