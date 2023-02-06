// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/campaign": {
            "get": {
                "description": "Get all campaigns",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "campaign"
                ],
                "summary": "Get all campaigns",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Campaign"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a campaign",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "campaign"
                ],
                "summary": "Create a campaign",
                "parameters": [
                    {
                        "description": "New Campaign",
                        "name": "campaign",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Campaign"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Campaign"
                        }
                    }
                }
            }
        },
        "/campaign/{campaign_id}": {
            "get": {
                "description": "Get a single campaign by its UUID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "campaign"
                ],
                "summary": "Get a campaign",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Campaign ID",
                        "name": "campaign_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Campaign"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a campaign",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "campaign"
                ],
                "summary": "Update a campaign",
                "parameters": [
                    {
                        "description": "New Campaign",
                        "name": "campaign",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Campaign"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Campaign ID",
                        "name": "campaign_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Campaign"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a campaign",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "campaign"
                ],
                "summary": "Delete a campaign",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Campaign ID",
                        "name": "campaign_id",
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
        "models.Campaign": {
            "type": "object",
            "required": [
                "end_date",
                "mcc",
                "merchant",
                "name",
                "reward_amount",
                "reward_program",
                "start_date"
            ],
            "properties": {
                "end_date": {
                    "description": "should be later than Start",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "mcc": {
                    "type": "integer",
                    "maximum": 9999,
                    "minimum": 0
                },
                "merchant": {
                    "type": "string",
                    "minLength": 1
                },
                "min_spend": {
                    "type": "number",
                    "minimum": 0
                },
                "name": {
                    "type": "string",
                    "minLength": 1
                },
                "reward_amount": {
                    "type": "integer"
                },
                "reward_program": {
                    "type": "string",
                    "minLength": 1
                },
                "start_date": {
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