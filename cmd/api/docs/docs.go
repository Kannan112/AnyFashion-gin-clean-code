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
        "/admin/adminlogin": {
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
        "/admin/category/add": {
            "post": {
                "description": "Admin can create new category from admin panel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Category"
                ],
                "summary": "Create new product category",
                "operationId": "create-category",
                "parameters": [
                    {
                        "description": "New category name",
                        "name": "category_name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.Category"
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
        "/admin/coupon": {
            "get": {
                "description": "Admins and users can see all available coupons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coupon"
                ],
                "summary": "Admins and users can see all available coupons",
                "operationId": "view-coupons",
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
        "/admin/coupon/add": {
            "post": {
                "description": "Admin can create new coupons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coupon"
                ],
                "summary": "Admin can create new coupon",
                "operationId": "create-coupon",
                "parameters": [
                    {
                        "description": "details of new coupon to be created",
                        "name": "new_coupon_details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.Coupons"
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
        "/admin/coupon/delete/{couponId}": {
            "delete": {
                "description": "Delete coupon",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coupon"
                ],
                "summary": "Delete a coupon",
                "operationId": "DeleteCoupon",
                "parameters": [
                    {
                        "type": "string",
                        "description": "New Admin details",
                        "name": "couponId",
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
        "/admin/coupon/update/{couponId}": {
            "patch": {
                "description": "admin coupon update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coupon"
                ],
                "summary": "Update a existing coupon",
                "operationId": "UpdateCoupon",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Coupon ID",
                        "name": "couponId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New Admin details",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UpdateCoupon"
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
        "/admin/dashbord/list": {
            "get": {
                "description": "Admin can access dashboard and view details regarding orders, products, etc.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Admin Dashboard",
                "operationId": "admin-dashboard",
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
        "/admin/user/block": {
            "patch": {
                "description": "admin block user access to the store",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Admin BlockUser",
                "parameters": [
                    {
                        "description": "User bolocking details",
                        "name": "blocking_details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.BlockData"
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
        "/admin/user/unblock/{user_id}": {
            "patch": {
                "description": "Admins can block users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Admin can unbolock a blocked user",
                "operationId": "unblock-users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the user to be blocked",
                        "name": "user_id",
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
        "/user/cart/add/:product_item_id": {
            "post": {
                "description": "User can add product item to the cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "User add product-item to cart",
                "operationId": "add-to-cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product_item_id",
                        "name": "product-items-id",
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
                            "$ref": "#/definitions/req.Otpverifier"
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
        "/user/wishlist/add/{itemId}": {
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
        },
        "/user/wishlist/remove/:itemId": {
            "delete": {
                "description": "Remove item from wishlist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wishlist"
                ],
                "summary": "Remove Item",
                "operationId": "RemoveFromWishlis",
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
        "req.BlockData": {
            "type": "object",
            "required": [
                "reason",
                "userid"
            ],
            "properties": {
                "reason": {
                    "type": "string"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "req.Category": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "req.Coupons": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "discountmaximumamount": {
                    "type": "number"
                },
                "discountpercent": {
                    "type": "number"
                },
                "expirationdate": {
                    "type": "string"
                },
                "minimumpurchaseamount": {
                    "type": "number"
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
        "req.Otpverifier": {
            "type": "object",
            "required": [
                "phoneNumber",
                "pin"
            ],
            "properties": {
                "phoneNumber": {
                    "type": "string"
                },
                "pin": {
                    "type": "string"
                }
            }
        },
        "req.UpdateCoupon": {
            "type": "object",
            "properties": {
                "discountmaximumamount": {
                    "type": "number"
                },
                "discountpercent": {
                    "type": "number"
                },
                "expirationdate": {
                    "type": "string"
                },
                "minimumpurchaseamount": {
                    "type": "number"
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
