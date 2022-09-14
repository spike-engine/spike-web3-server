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
        "/query-api/v1/balance": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "query wallet balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wallet bsc address",
                        "name": "wallet_address",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/query-api/v1/nft/list": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "query single nft list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wallet bsc address",
                        "name": "wallet_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "nft contract address",
                        "name": "contract_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "nft type",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/query-api/v1/nft/type": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "query all nft type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wallet bsc address",
                        "name": "wallet_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "nft contract address",
                        "name": "contract_address",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/query-api/v1/txRecord/erc20": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "query wallet ERC20 tx list(7 days)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wallet bsc address",
                        "name": "wallet_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "erc20 contract address",
                        "name": "contract_address",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/query-api/v1/txRecord/native": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "query wallet native tx list(7 days)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wallet bsc address",
                        "name": "wallet_address",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/tx-api/v1/client/importNft": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "import nft",
                "parameters": [
                    {
                        "type": "string",
                        "description": "game orderId",
                        "name": "order_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tx fromAddress",
                        "name": "from_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "nft contract address",
                        "name": "contract_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "nft token id",
                        "name": "token_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tx hash",
                        "name": "tx_hash",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "game callBack url address",
                        "name": "cb",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/tx-api/v1/client/rechargeToken": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "recharge token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "game orderId",
                        "name": "order_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tx fromAddress",
                        "name": "from_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tx token amount",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token contract address(native : 0x0000000000000000000000000000000000000000)",
                        "name": "contract_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tx hash",
                        "name": "tx_hash",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "game callBack url address",
                        "name": "cb",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/tx-api/v1/hotWallet/mint": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "mint nft",
                "parameters": [
                    {
                        "type": "string",
                        "description": "game orderId",
                        "name": "order_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "nft tokenUri",
                        "name": "token_uri",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "game callBack url address",
                        "name": "cb",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/tx-api/v1/hotWallet/withdrawNFT": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "withdraw nft",
                "parameters": [
                    {
                        "type": "string",
                        "description": "game orderId",
                        "name": "order_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tx toAddress",
                        "name": "to_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "nft token id",
                        "name": "token_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "nft contract address",
                        "name": "contract_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "game callBack url address",
                        "name": "cb",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/tx-api/v1/hotWallet/withdrawToken": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "withdraw token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "game orderId",
                        "name": "order_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tx toAddress",
                        "name": "to_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tx token amount",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token contract address(native : 0x0000000000000000000000000000000000000000)",
                        "name": "contract_address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "game callBack url address",
                        "name": "cb",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
