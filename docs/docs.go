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
        "/api/songs": {
            "post": {
                "description": "Create a new song by providing the group and song title.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Create a new song",
                "parameters": [
                    {
                        "description": "Song data",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song was created successfully",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to create song",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        },
        "/api/songs/filter": {
            "get": {
                "description": "Get a list of songs based on filter criteria like artist, title, release date, text, and source link with pagination (limit and offset).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Get all songs with filtering and pagination",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Artist name",
                        "name": "artist",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Song title",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Release date",
                        "name": "release_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Text content",
                        "name": "text",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Source link",
                        "name": "source_link",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit of results",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset for pagination",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of songs",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.SongResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid query parameters",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch songs",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        },
        "/api/songs/{id}": {
            "get": {
                "description": "Get the text of a song by its ID. If the song is not found, returns a 404 error.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Get song text by song ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song text retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/handlers.SongTextResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid song ID",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing song by its ID. If the song is not found, returns a 404 error.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Update a song by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Song data to update",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song updated successfully",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid song ID or request body",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a song from the database using its unique ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Delete a song by its ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song was deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid song ID",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "404": {
                        "description": "Song with this ID isn't present",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to delete song",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.Request": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                }
            }
        },
        "handlers.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handlers.SongResponse": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "release_date": {
                    "type": "string"
                },
                "source_link": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handlers.SongTextResponse": {
            "type": "object",
            "properties": {
                "text_parts": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "handlers.UpdateRequest": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "source_link": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Song Library API",
	Description:      "API для управления музыкальной библиотекой",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
