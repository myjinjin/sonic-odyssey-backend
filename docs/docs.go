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
        "/api/v1/auth/sign-in": {
            "post": {
                "description": "Responds with a JWT token and expiration time upon successful login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "Login Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/auth.UnauthorizedResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/auth.UnauthorizedResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "post": {
                "description": "User SignUp",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User SignUp",
                "parameters": [
                    {
                        "description": "SignUp Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/me": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "JWT 인증 토큰 기반 내 유저 정보 조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get my user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.GetMyUserInfoResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "JWT 인증 토큰 기반 내 유저 정보 업데이트",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Patch my user info",
                "parameters": [
                    {
                        "description": "PatchMyUser Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.PatchMyUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.PatchMyUserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/password/recovery": {
            "post": {
                "description": "비밀번호 복구 이메일 전송",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Send password recovery email",
                "parameters": [
                    {
                        "description": "SendPasswordRecoveryEmailRequest Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.SendPasswordRecoveryEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.SendPasswordRecoveryEmailResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/password/reset": {
            "post": {
                "description": "비밀번호 재설정",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Reset password",
                "parameters": [
                    {
                        "description": "ResetPasswordRequest Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.ResetPasswordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "odyssey@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "Example123!"
                }
            }
        },
        "auth.LoginResponse": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string",
                    "example": "2024-05-30T09:00:00Z"
                },
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
                }
            }
        },
        "auth.UnauthorizedResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "incorrect Username or Password"
                }
            }
        },
        "v1.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "v1.GetMyUserInfoResponse": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string",
                    "example": "bio..."
                },
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "name": {
                    "type": "string",
                    "example": "name"
                },
                "nickname": {
                    "type": "string",
                    "example": "nickname"
                },
                "profile_image_url": {
                    "type": "string",
                    "example": "https://example.com/profile.png"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                },
                "website": {
                    "type": "string",
                    "example": "https://example.com"
                }
            }
        },
        "v1.PatchMyUserRequest": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string",
                    "example": "newbio"
                },
                "name": {
                    "type": "string",
                    "minLength": 2,
                    "example": "newname"
                },
                "nickname": {
                    "type": "string",
                    "minLength": 5,
                    "example": "newnickname"
                },
                "website": {
                    "type": "string",
                    "example": "https://example.com/new"
                }
            }
        },
        "v1.PatchMyUserResponse": {
            "type": "object"
        },
        "v1.ResetPasswordRequest": {
            "type": "object",
            "required": [
                "flow_id",
                "password"
            ],
            "properties": {
                "flow_id": {
                    "type": "string",
                    "example": "cc833698-4519-4873-b9b4-67d6fef70dcb:1717170088"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "password123"
                }
            }
        },
        "v1.ResetPasswordResponse": {
            "type": "object"
        },
        "v1.SendPasswordRecoveryEmailRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                }
            }
        },
        "v1.SendPasswordRecoveryEmailResponse": {
            "type": "object"
        },
        "v1.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "nickname",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "nickname": {
                    "type": "string",
                    "example": "johndoe"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "password123"
                }
            }
        },
        "v1.SignUpResponse": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "integer",
                    "example": 1
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
