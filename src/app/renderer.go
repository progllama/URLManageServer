package app

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

var Renderer = multitemplate.NewRenderer()

func loadTemplates(templatesDir string) {
	layouts, err := filepath.Glob(templatesDir + "/layout/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	contents, err := filepath.Glob(templatesDir + "/content/**/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layout/ and include/ directories
	for _, include := range contents {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)

		Renderer.AddFromFiles(include[17:], files...)
	}
}
