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
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/feedback.CreateFeedbackRequestBody"
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
        "/feedback/import": {
            "post": {
                "description": "Imports feedback from an uploaded CSV file.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "feedback"
                ],
                "summary": "Imports feedback from an uploaded CSV file.",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/feedback.ImportFeedbackRequestBody"
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
        "feedback.CreateFeedbackRequestBody": {
            "type": "object",
            "required": [
                "data",
                "type"
            ],
            "properties": {
                "data": {
                    "$ref": "#/definitions/feedback.FeedbackData"
                },
                "meta": {
                    "$ref": "#/definitions/feedback.FeedbackMeta"
                },
                "projectId": {
                    "type": "string"
                },
                "tags": {
                    "$ref": "#/definitions/feedback.FeedbackTags"
                },
                "type": {
                    "enum": [
                        "text",
                        "cnps",
                        "csat"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/feedback.FeedbackType"
                        }
                    ]
                },
                "user": {
                    "$ref": "#/definitions/feedback.FeedbackUser"
                }
            }
        },
        "feedback.FeedbackData": {
            "type": "object",
            "properties": {
                "cnps": {
                    "type": "number",
                    "maximum": 1,
                    "minimum": 0
                },
                "csat": {
                    "type": "number",
                    "maximum": 1,
                    "minimum": 0
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "feedback.FeedbackMeta": {
            "type": "object",
            "additionalProperties": true
        },
        "feedback.FeedbackTags": {
            "type": "object",
            "additionalProperties": {
                "type": "string"
            }
        },
        "feedback.FeedbackType": {
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
        },
        "feedback.FeedbackUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "feedback.ImportFeedbackRequestBody": {
            "type": "object",
            "required": [
                "link"
            ],
            "properties": {
                "link": {
                    "type": "string"
                },
                "projectId": {
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
