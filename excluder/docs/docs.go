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
        "/exclusion": {
            "get": {
                "description": "Get all exclusions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exclusion"
                ],
                "summary": "Get all exclusions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Exclusion"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create an exclusion",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exclusion"
                ],
                "summary": "Create an exclusion",
                "parameters": [
                    {
                        "description": "New Exclusion",
                        "name": "exclusion",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Exclusion"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Exclusion"
                        }
                    }
                }
            }
        },
        "/exclusion/{exclusion_id}": {
            "get": {
                "description": "Get a single exclusion by its UUID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exclusion"
                ],
                "summary": "Get an exclusion",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exclusion ID",
                        "name": "exclusion_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Exclusion"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an exclusion",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exclusion"
                ],
                "summary": "Update an exclusion",
                "parameters": [
                    {
                        "description": "New Exclusion",
                        "name": "exclusion",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Exclusion"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Exclusion ID",
                        "name": "exclusion_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Exclusion"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an exclusion",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exclusion"
                ],
                "summary": "Delete an exclusion",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exclusion ID",
                        "name": "exclusion_id",
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
        }
    },
    "definitions": {
        "models.Exclusion": {
            "type": "object",
            "required": [
                "mcc",
                "valid_from"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "mcc": {
                    "type": "integer",
                    "maximum": 9999,
                    "minimum": 1
                },
                "valid_from": {
                    "description": "should be later than time.Now()",
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
