// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
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
        "/organization/{organization_id}/credential": {
            "post": {
                "description": "Creates access tokens which can be used to authenticate when sending feedback.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "credentials"
                ],
                "summary": "Creates access tokens",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "client_id",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Client Secret",
                        "name": "client_secret",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/credentials.CreateAccessTokenResponseBody"
                        }
                    }
                }
            }
        },
        "/organization/{organization_id}/feedback": {
            "post": {
                "description": "Creates a new feedback for the organization.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "feedbacks"
                ],
                "summary": "Creates a new feedback",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path"
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
                "description": "Basic healthcheck endpoint that can be used to check whether the service can be reached.",
                "tags": [
                    "healthchecks"
                ],
                "summary": "Basic healthcheck endpoint",
                "responses": {
                    "204": {
                        "description": "ok"
                    }
                }
            }
        }
    },
    "definitions": {
        "credentials.CreateAccessTokenResponseBody": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "https://api.provar.se",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Provar API",
	Description:      "Provar.se API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
