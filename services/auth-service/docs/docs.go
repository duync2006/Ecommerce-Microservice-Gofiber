// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/forgot": {
            "post": {
                "description": "User lost their password, they need to change another password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Forgot password",
                "parameters": [
                    {
                        "description": "Email",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EmailForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login to the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Login to the system",
                "parameters": [
                    {
                        "description": "LoginForm",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "get": {
                "description": "Logout to the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Logout to the system",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "get": {
                "description": "Refresh new access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Refresh new access token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "UserModel",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/submitOTP": {
            "post": {
                "description": "Submit new password to change password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Reset password",
                "parameters": [
                    {
                        "description": "Password",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.NewPasswordForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ChangePasswordForm": {
            "type": "object",
            "properties": {
                "new_pwd": {
                    "type": "string"
                },
                "old_pwd": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.EmailForm": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "models.LoginForm": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.NewPasswordForm": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.OTPForm": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "otp": {
                    "type": "string"
                }
            }
        },
        "models.RegisterForm": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Ecommerce Authentication Service API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
