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
        "/editors/news": {
            "get": {
                "security": [
                    {
                        "EditorsAuth": []
                    }
                ],
                "description": "editor get all news",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "editors"
                ],
                "summary": "Editor Get All News",
                "operationId": "editorGetAllNews",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/university.News"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    }
                }
            }
        },
        "/editors/news/{id}": {
            "get": {
                "security": [
                    {
                        "EditorsAuth": []
                    }
                ],
                "description": "editor get news by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "editors"
                ],
                "summary": "Editor Get News By ID",
                "operationId": "editorsGetNewsById",
                "parameters": [
                    {
                        "type": "string",
                        "description": "news id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/university.News"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    }
                }
            }
        },
        "/editors/refresh": {
            "post": {
                "security": [
                    {
                        "EditorsAuth": []
                    }
                ],
                "description": "editor refresh tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "editors"
                ],
                "summary": "Editor Refresh Token",
                "operationId": "editorRefresh",
                "parameters": [
                    {
                        "description": "sign up info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpv1.refreshInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpv1.tokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    }
                }
            }
        },
        "/editors/sign-in": {
            "post": {
                "description": "editor sign in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "editors"
                ],
                "summary": "Editor SignIn",
                "operationId": "editorSignIn",
                "parameters": [
                    {
                        "description": "sign in info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpv1.editorsSignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpv1.tokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    }
                }
            }
        },
        "/editors/sign-up": {
            "post": {
                "description": "create editor account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "editors"
                ],
                "summary": "Editor SignUp",
                "operationId": "editorSignUp",
                "parameters": [
                    {
                        "description": "sign up info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpv1.editorsSignUpInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    }
                }
            }
        },
        "/editors/verify/{code}": {
            "post": {
                "description": "editor verify registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "editors"
                ],
                "summary": "Editor Verify Registration",
                "operationId": "editorVerify",
                "parameters": [
                    {
                        "type": "string",
                        "description": "verification code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpv1.tokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/httpv1.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpv1.editorsSignInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "httpv1.editorsSignUpInput": {
            "type": "object",
            "required": [
                "email",
                "hash",
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "hash": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "httpv1.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "httpv1.refreshInput": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "httpv1.tokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.Editor": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "description": "id записи в MongoDB.",
                    "type": "string"
                },
                "last_visit_at": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "registered_at": {
                    "type": "string"
                },
                "session": {
                    "$ref": "#/definitions/models.Session"
                },
                "university_id": {
                    "type": "string"
                },
                "verification": {
                    "description": "Статус верификации",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Verification"
                        }
                    ]
                }
            }
        },
        "models.Session": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.Verification": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "verified": {
                    "type": "boolean"
                }
            }
        },
        "university.News": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "Основной текст новостной записи.",
                    "type": "string"
                },
                "created_at": {
                    "description": "Дата создания новостной записи.",
                    "type": "string"
                },
                "created_by": {
                    "description": "Автор новостной записи.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Editor"
                        }
                    ]
                },
                "description": {
                    "description": "Краткое описание новостной записи.",
                    "type": "string"
                },
                "header": {
                    "description": "Заголовок новостной записи.",
                    "type": "string"
                },
                "id": {
                    "description": "id записи в MongoDB.",
                    "type": "string"
                },
                "image_url": {
                    "description": "Ссылка на основное изображение новостной записи.",
                    "type": "string"
                },
                "published": {
                    "description": "Статус публикации новостной записи.",
                    "type": "boolean"
                },
                "updated_at": {
                    "description": "Дата последнего обновления новостной записи.",
                    "type": "string"
                },
                "updated_by": {
                    "description": "Редакторы новостной записи.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Editor"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "AdminAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "EditorsAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "University Platform API",
	Description:      "API Server for University Platform",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
