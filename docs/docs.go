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
        "/": {
            "get": {
                "description": "get Books",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter. e.g. col1:v1,col2:v2",
                        "name": "query",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Fields returned. e.g. col1,col2",
                        "name": "fields",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sorted-by fields. e.g. col1,col2",
                        "name": "sortby",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Limit the size of result set. Must be an integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Start position of result set. Must be an integer",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Books"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    }
                }
            },
            "post": {
                "description": "create Books",
                "parameters": [
                    {
                        "description": "body for Books content",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Books"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "int"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "body"
                        }
                    }
                }
            }
        },
        "/:id": {
            "get": {
                "description": "get Books by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The key for staticblock",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Books"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "id"
                        }
                    }
                }
            },
            "put": {
                "description": "update the Books",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The id you want to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body for Books content",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Books"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Books"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "id"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete the Books",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The id you want to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "id"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Books": {
            "type": "object",
            "properties": {
                "booksAuthor": {
                    "type": "string"
                },
                "booksDescription": {
                    "type": "string"
                },
                "booksPublished": {
                    "type": "string"
                },
                "booksPublisher": {
                    "type": "string"
                },
                "booksSubtitle": {
                    "type": "string"
                },
                "booksTitle": {
                    "type": "string"
                },
                "id": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
