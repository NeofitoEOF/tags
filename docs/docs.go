package docs

import "github.com/swaggo/swag"

const docTemplate = `{
	"schemes": {{ marshal .Schemes }},
	"swagger": "2.0",
	"info": {
		"description": "{{espace .Description}}",
		"title": "{{.Title}}",
		"contact": {},
		"version": "{{.Version}}"
	},
	"host": "{{.Host}}",
	"basePath": "{{.BasePath}}",
	"paths": {}
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Tag Service API",
	Description:      "A Tag service API in Go using Gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
