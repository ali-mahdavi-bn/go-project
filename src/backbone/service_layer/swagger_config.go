package service_layer

import (
	"go-tel/src/backbone/application/constans"
	"strings"
)

var PathsSwagger = ""

func setupSwaggerConfig(e *APIRouter, fullPath string, method string, handler HandlerFunc, swaggerConfig *SwaggerConfig) {
	if swaggerConfig.Path == "" {
		swaggerConfig.Path = fullPath
	}
	if swaggerConfig.Summary == "" {
		swaggerConfig.Summary = getNameFunc(handler)
	}
	if swaggerConfig.Method == "" {
		swaggerConfig.Method = method
	}
	if swaggerConfig.Tag == "" {
		swaggerConfig.Tag = e.Tag
	}
}

func AddSwaggerDoc() string {

	docTemplate := strings.Replace(constans.BaseDocTemplate, "{{Paths}}", "{"+PathsSwagger+"}", -1)
	return docTemplate
}

func GenerateSwaggerConfig(swaggerConfig *SwaggerConfig) {

	configSwagger := constans.BaseConfigPaths
	replacements := map[string]string{
		"{{Path}}":        isString(swaggerConfig.Path, "/"),
		"{{method}}":      isString(strings.ToLower(swaggerConfig.Method), "get"),
		"{{status}}":      isString(swaggerConfig.Status, "200"),
		"{{summary}}":     isString(swaggerConfig.Summary, ""),
		"{{description}}": isString(swaggerConfig.Description, "description"),
		"{{tag}}":         isString(swaggerConfig.Tag, "root"),
	}
	for placeholder, value := range replacements {
		configSwagger = strings.Replace(configSwagger, placeholder, value, -1)
	}
	PathsSwagger += configSwagger
}

func isString(status any, defaultValue string) string {
	value, ok := status.(string)
	if ok == false {
		return defaultValue
	}
	return value
}
