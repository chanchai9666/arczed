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
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/api/const/createConst": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "เพิ่มข้อมูลค่าคงที่",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Const"
                ],
                "summary": "เพิ่มข้อมูลค่าคงที่",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": " request body ",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.ConfigConstant"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/const/deleteConst/{group_id}/{const_id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ลบข้อมูล ค่าคงที่",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Const"
                ],
                "summary": "ลบข้อมูล ค่าคงที่",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "group id",
                        "name": "group_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "const id",
                        "name": "const_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/const/updateConst": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "แก้ไขข้อมูลค่าคงที่",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Const"
                ],
                "summary": "แก้ไขข้อมูลค่าคงที่",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": " request body ",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.ConfigConstant"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Login เข้าใช้งานสำหรับขอ token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login เข้าใช้งาน",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": " request body ",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.LoginResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show User ตามเงื่อนไข",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "ค้นหา User ตามเงื่อนไข",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "อีเมล",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ชื่อ",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "นามสกุล",
                        "name": "sur_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ผู้ใช้งาน",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/arczed_internal_entities_models.Users"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/users/createUsers": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "เพิ่มข้อมูล User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "เพิ่มข้อมูล User",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": " request body ",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.AddUsers"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/users/deleteUsers/{user_id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ลบข้อมูล User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "ลบข้อมูล User",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/users/updateUsers": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "แก้ไขข้อมูล User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "แก้ไขข้อมูล User",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": " request body ",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.AddUsers"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/users/usersAll": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show User ทั้งหมด",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "ค้นหา User ทั้งหมด",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/arczed_internal_entities_models.Users"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/users/{user_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show User ตาม UserId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "ค้นหา User ตาม UserId",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_models.Users"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.HTTPError"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Health checking for the service",
                "produces": [
                    "text/plain"
                ],
                "summary": "Health Check",
                "operationId": "HealthCheckHandler",
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
        "arczed_internal_entities_models.Users": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "วันที่สร้าง",
                    "type": "string"
                },
                "created_user": {
                    "description": "ผู้สร้าง",
                    "type": "string"
                },
                "deleted_at": {
                    "description": "วันเวลาที่ลบ",
                    "allOf": [
                        {
                            "$ref": "#/definitions/gorm.DeletedAt"
                        }
                    ]
                },
                "deleted_user": {
                    "description": "ผุ้ลบ",
                    "type": "string"
                },
                "email": {
                    "description": "อีเมล",
                    "type": "string"
                },
                "id_card": {
                    "description": "รหัสบัตรประจำตัว",
                    "type": "string"
                },
                "is_active": {
                    "description": "สถานะใช้งาน",
                    "type": "integer"
                },
                "level": {
                    "description": "ความสัมพันธ์กับ UserLevels",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/arczed_internal_entities_models.UsersLevels"
                    }
                },
                "name": {
                    "description": "ชื่อ",
                    "type": "string"
                },
                "password": {
                    "description": "ชื่อ โปรไฟล์",
                    "type": "string"
                },
                "phone_number": {
                    "description": "หมายเลขโทรศัพท์",
                    "type": "string"
                },
                "sur_name": {
                    "description": "นามสกุล",
                    "type": "string"
                },
                "updated_at": {
                    "description": "วันเวลาที่อัพเดทล่าสุด",
                    "type": "string"
                },
                "updated_user": {
                    "description": "ผู้อัพเดทล่าสุด",
                    "type": "string"
                },
                "user_id": {
                    "description": "ไอดี ของผู้ใช้งาน",
                    "type": "integer"
                }
            }
        },
        "arczed_internal_entities_models.UsersLevels": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "วันที่สร้าง",
                    "type": "string"
                },
                "created_user": {
                    "description": "ผู้สร้าง",
                    "type": "string"
                },
                "deleted_at": {
                    "description": "วันเวลาที่ลบ",
                    "allOf": [
                        {
                            "$ref": "#/definitions/gorm.DeletedAt"
                        }
                    ]
                },
                "deleted_user": {
                    "description": "ผุ้ลบ",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "description": "สถานะใช้งาน",
                    "type": "integer"
                },
                "level": {
                    "type": "string"
                },
                "updated_at": {
                    "description": "วันเวลาที่อัพเดทล่าสุด",
                    "type": "string"
                },
                "updated_user": {
                    "description": "ผู้อัพเดทล่าสุด",
                    "type": "string"
                },
                "userID": {
                    "description": "ต้องกำหนดประเภทและขนาดให้ตรงกับ Users.UserId",
                    "type": "integer"
                }
            }
        },
        "arczed_internal_entities_schemas.AddUsers": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "description": "อีเมล",
                    "type": "string"
                },
                "id_card": {
                    "description": "เลขบัตรประจำตัว",
                    "type": "string"
                },
                "name": {
                    "description": "ชื่อ",
                    "type": "string"
                },
                "password": {
                    "description": "รหัสผ่าน",
                    "type": "string"
                },
                "phone_number": {
                    "description": "เบอร์โทร",
                    "type": "string"
                },
                "sur_name": {
                    "description": "นามสกุล",
                    "type": "string"
                },
                "user_id": {
                    "description": "ผู้ใช้งาน",
                    "type": "integer"
                }
            }
        },
        "arczed_internal_entities_schemas.ConfigConstant": {
            "type": "object",
            "properties": {
                "const_id": {
                    "description": "ไอดีค่าคงที่",
                    "type": "string"
                },
                "group_id": {
                    "description": "ไอดีกลุ่มค่าคงที่",
                    "type": "string"
                },
                "name_en": {
                    "description": "ชื่อค่าคงที่ EN",
                    "type": "string"
                },
                "name_th": {
                    "description": "ชื่อค่าคงที่ TH",
                    "type": "string"
                },
                "ref_value1": {
                    "description": "ค่าอ้างอิง 1",
                    "type": "string"
                },
                "ref_value2": {
                    "description": "ค่าอ้างอิง 2",
                    "type": "string"
                },
                "ref_value3": {
                    "description": "ค่าอ้างอิง 3",
                    "type": "string"
                },
                "sort": {
                    "description": "ลำดับ",
                    "type": "integer"
                }
            }
        },
        "arczed_internal_entities_schemas.HTTPError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "arczed_internal_entities_schemas.LoginReq": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "ผู้ใช้งาน",
                    "type": "string"
                },
                "password": {
                    "description": "รหัสผ่าน",
                    "type": "string"
                }
            }
        },
        "arczed_internal_entities_schemas.LoginResp": {
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "Token เข้าใช้งาน",
                    "type": "string"
                },
                "user": {
                    "description": "ข้อมูลผู้ใช้",
                    "allOf": [
                        {
                            "$ref": "#/definitions/arczed_internal_entities_schemas.UserResp"
                        }
                    ]
                }
            }
        },
        "arczed_internal_entities_schemas.UserResp": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "อีเมล",
                    "type": "string"
                },
                "level": {
                    "description": "เลเวล",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "ชื่อ",
                    "type": "string"
                },
                "sur_name": {
                    "description": "นามสกุล",
                    "type": "string"
                },
                "user_id": {
                    "description": "ผู้ใช้งาน",
                    "type": "string"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "ใส่ค่า Bearer เว้นวรรคและตามด้วย TOKEN  ex(Bearer ?????????)"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "User API",
	Description:      "This is a sample server for user management.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
