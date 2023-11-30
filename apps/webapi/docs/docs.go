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
        "/feedback": {
            "post": {
                "description": "Create a new feedback event for an organization.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "feedback"
                ],
                "summary": "Create a new feedback event for an organization.",
                "parameters": [
                    {
                        "description": "The request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/feedback.RequestBody"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "ok"
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Ensure that the service can be reached.",
                "tags": [
                    "health"
                ],
                "summary": "Ensure that the service can be reached.",
                "responses": {
                    "204": {
                        "description": "ok"
                    }
                }
            }
        },
        "/ping/secure": {
            "get": {
                "description": "Ensure that the api key is valid",
                "tags": [
                    "health"
                ],
                "summary": "Ensure that the api key is valid",
                "responses": {
                    "204": {
                        "description": "ok"
                    }
                }
            }
        }
    },
    "definitions": {
        "feedback.RequestBody": {
            "type": "object",
            "required": [
                "data",
                "type"
            ],
            "properties": {
                "data": {
                    "$ref": "#/definitions/repository.FeedbackData"
                },
                "projectId": {
                    "type": "string"
                },
                "tags": {
                    "$ref": "#/definitions/repository.FeedbackTags"
                },
                "type": {
                    "enum": [
                        "text",
                        "cnps",
                        "csat"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/repository.FeedbackType"
                        }
                    ]
                }
            }
        },
        "repository.FeedbackData": {
            "type": "object",
            "properties": {
                "cnps": {
                    "type": "number"
                },
                "csat": {
                    "type": "number"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "repository.FeedbackTags": {
            "type": "object",
            "additionalProperties": {
                "type": "string"
            }
        },
        "repository.FeedbackType": {
            "type": "string",
            "enum": [
                "text",
                "cnps",
                "csat"
            ],
            "x-enum-varnames": [
                "FeedbackTypeText",
                "FeedbackTypeCNPS",
                "FeedbackTypeCSAT"
            ]
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
