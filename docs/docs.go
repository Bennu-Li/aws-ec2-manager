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
        "/create": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Cretae a rc2 instance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "Create a ec2 instance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user name",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "instance image id",
                        "name": "imageId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "instance type",
                        "name": "instanceType",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "instance rootDiskSize",
                        "name": "rootDiskSize",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "instance deviceDiskSize",
                        "name": "deviceDiskSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "user department",
                        "name": "department",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/describe": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Describe a EC2 Instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "Describe a EC2 Instance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "instance Id",
                        "name": "instanceId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/key": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Cretae a rsa keyPair for the user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "Create a keyPair",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user name",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/list": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "DList EC2 Instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "List EC2 Instance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserName",
                        "name": "userName",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/reboot": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Reboot a EC2 Instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "Reboot a EC2 Instance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "instance Id",
                        "name": "instanceId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/start": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Start a EC2 Instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "Start a EC2 Instance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "instance Id",
                        "name": "instanceId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/stop": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Stop a EC2 Instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "Stop a EC2 Instance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "instance Id",
                        "name": "instanceId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/terminate": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Terminate a EC2 Instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "Terminate a EC2 Instance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "instance Id",
                        "name": "instanceId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
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
	BasePath:         "/v1/ec2",
	Schemes:          []string{},
	Title:            "Manage EC2 API",
	Description:      "This API is used to manager aws ec2.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
