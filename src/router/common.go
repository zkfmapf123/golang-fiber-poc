package router

import (
	"fmt"
	"os"

	"github.com/gofiber/contrib/swagger"
)

func (router *Router) Swagger(path string, basePath string) {
	currentPath, _ := os.Getwd()
	router.r.Get(path, swagger.New(
		swagger.Config{
			BasePath: basePath,
			FilePath: fmt.Sprintf("%s/swagger/doc.json", currentPath),
		},
	))
}
