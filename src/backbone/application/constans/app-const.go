package constans

const (
	AppName         = "go-tel"
	BaseDocTemplate = `{
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
    "paths": {{Paths}}
}`

	BaseConfigPaths = `"{{Path}}": {
            "{{method}}": {
                "description": "{{description}}",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "{{tag}}"
                ],
                "summary": "{{summary}}",
                "responses": {
                    "{{status}}": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },`
)
