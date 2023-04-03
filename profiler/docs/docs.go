// Code generated by swaggo/swag. DO NOT EDIT
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
        "/auth/login": {
            "post": {
                "description": "Returns user when provided credentials are valid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/password": {
            "post": {
                "description": "Endpoint allows a user to change their default password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Change default password",
                "parameters": [
                    {
                        "description": "Credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ChangeDefaultPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/card": {
            "get": {
                "description": "Get all cards",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "Get all cards",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Card"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a card",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "Create a card",
                "parameters": [
                    {
                        "description": "New Card",
                        "name": "card",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CardInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Card"
                        }
                    }
                }
            }
        },
        "/card/type": {
            "get": {
                "description": "Get all card types",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card_type"
                ],
                "summary": "Get all card types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.CardType"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a card type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card_type"
                ],
                "summary": "Create a card type",
                "parameters": [
                    {
                        "description": "New Card Type",
                        "name": "card_type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CardType"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.CardType"
                        }
                    }
                }
            }
        },
        "/card/type/{card_type_pk}": {
            "put": {
                "description": "Update a card type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card_type"
                ],
                "summary": "Update a card type",
                "parameters": [
                    {
                        "description": "New Card Type",
                        "name": "card_type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CardType"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Campaign Type PK",
                        "name": "card_type_pk",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CardType"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a card type",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card_type"
                ],
                "summary": "Delete a card type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card Type PK",
                        "name": "card_type_pk",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/card/type/{card_type}": {
            "get": {
                "description": "Get a single card type by its PK",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card_type"
                ],
                "summary": "Get a card type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card Type PK",
                        "name": "card_type",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CardType"
                        }
                    }
                }
            }
        },
        "/card/{card_id}": {
            "get": {
                "description": "Get a single card by its UUID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "Get a card",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "card_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Card"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a card",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "Delete a card",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "card_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "health check",
                "produces": [
                    "application/json"
                ],
                "summary": "health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Get all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "New User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/user/{user_id}": {
            "get": {
                "description": "Get a single user by its UUID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "description": "Updated User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AuthResponse": {
            "type": "object",
            "properties": {
                "is_new": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.Card": {
            "type": "object",
            "required": [
                "card_pan",
                "card_type",
                "id",
                "user_id"
            ],
            "properties": {
                "card_pan": {
                    "type": "string"
                },
                "card_type": {
                    "description": "card belongs to one card type",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "user_id": {
                    "description": "card belongs to one user",
                    "type": "string"
                }
            }
        },
        "models.CardInput": {
            "type": "object",
            "required": [
                "card_pan",
                "card_type",
                "id",
                "user_id"
            ],
            "properties": {
                "card_pan": {
                    "type": "string"
                },
                "card_type": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.CardType": {
            "type": "object",
            "required": [
                "card_type",
                "name",
                "reward_program",
                "reward_unit"
            ],
            "properties": {
                "card_type": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "reward_program": {
                    "type": "string"
                },
                "reward_unit": {
                    "type": "string"
                }
            }
        },
        "models.ChangeDefaultPassword": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "old_password",
                "password"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.SignIn": {
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
        "models.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "creditCards": {
                    "description": "one user has many credit cards",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Card"
                    }
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isNew": {
                    "description": "is new user, password not changed from default",
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.UserInput": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "first_name",
                "id",
                "last_name",
                "password",
                "phone"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
