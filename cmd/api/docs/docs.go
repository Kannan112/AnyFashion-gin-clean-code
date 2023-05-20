// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/admin/category/delete/{category_id}": {
            "delete": {
                "description": "Admin can delete a category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Category"
                ],
                "summary": "Admin can delete a category",
                "operationId": "delete-category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "category_id",
                        "name": "category_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "description": "Logs in an admin user and returns an authentication token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Admin Login",
                "parameters": [
                    {
                        "description": "Admin login details",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/address/add": {
            "post": {
                "description": "Login as a user to access the ecommerce site",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Address"
                ],
                "summary": "Add Address",
                "operationId": "AddAddress",
                "parameters": [
                    {
                        "description": "Input Field",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.Address"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/address/listall": {
            "get": {
                "description": "Login as a user to access the ecommerce site",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Address"
                ],
                "summary": "List Address",
                "operationId": "ListallAddres",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/address/update": {
            "patch": {
                "description": "Login as a user to access the ecommerce site",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Address"
                ],
                "summary": "Update Address",
                "operationId": "UpdateAddress",
                "parameters": [
                    {
                        "description": "Input Field",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.Address"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login as a user to access the ecommerce site",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User Login",
                "operationId": "UserLogin",
                "parameters": [
                    {
                        "description": "Input Field",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "get": {
                "description": "User logout to access the ecommerce site",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User Logout",
                "operationId": "UserLogout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/otp/send": {
            "post": {
                "description": "Send OTP to use's mobile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Otp"
                ],
                "summary": "Send OTP to user's mobile",
                "operationId": "send-otp",
                "parameters": [
                    {
                        "description": "User mobile number",
                        "name": "user_mobile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.OTPData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/otp/verify": {
            "post": {
                "description": "Validate the  OTP sent to use's mobile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Otp"
                ],
                "summary": "Validate the OTP to user's mobile",
                "operationId": "validate-otp",
                "parameters": [
                    {
                        "description": "OTP sent to user's mobile number",
                        "name": "otp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.VerifyOtp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/profile/edit": {
            "patch": {
                "description": "Edit user prodile ecommerce site",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Edit Profile",
                "operationId": "EditProfile",
                "parameters": [
                    {
                        "description": "Input Field",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/profile/view": {
            "get": {
                "description": "Login as a user to access the ecommerce site",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "View Profile",
                "operationId": "ViewProfile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/wishlist/add/:itemId": {
            "post": {
                "description": "Login as a user to access the ecommerce site",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wishlist"
                ],
                "summary": "Add Wishlist",
                "operationId": "AddToWishlist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "itemId",
                        "name": "itemId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "req.Address": {
            "type": "object",
            "required": [
                "city",
                "district",
                "house_number",
                "landmark",
                "pincode",
                "street"
            ],
            "properties": {
                "city": {
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "house_number": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isdefault": {
                    "type": "boolean"
                },
                "landmark": {
                    "type": "string"
                },
                "pincode": {
                    "type": "integer"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "req.LoginReq": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "req.OTPData": {
            "type": "object",
            "properties": {
                "phoneNumber": {
                    "type": "string"
                }
            }
        },
        "req.UserReq": {
            "type": "object",
            "required": [
                "email",
                "mobile",
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "req.VerifyOtp": {
            "type": "object",
            "required": [
                "code",
                "user"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/req.OTPData"
                }
            }
        },
        "res.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                },
                "stastuscode": {
                    "type": "integer"
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
