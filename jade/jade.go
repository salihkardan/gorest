package jade

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/Joker/jade"
	"github.com/gin-gonic/gin"
)

// RenderJade renders given path to jade, expects path ends with .html and looks up for .jade file
func RenderJade(base, filePath string) (string, error) {
	// TODO: Escape path
	newFilePath := strings.Replace(filePath, ".html", ".jade", -1)
	// TODO: Join properly, dont allow directroy traversing
	newPath := path.Join(base, newFilePath)

	// Read the jade file
	file, err := ioutil.ReadFile(newPath)
	if err != nil {
		return "", err
	}

	// Convert thed file to string
	k := string(file)

	// Parse the template
	tpl, err := jade.Parse(filePath, k)
	if err != nil {
		return "", err
	}

	return tpl, nil
}

// RenderJadeFromBasePath is meant to be used as a HTTP handler in gin
func RenderJadeFromBasePath(basepath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Param("filepath")
		s, _ := RenderJade(basepath, path)
		c.Header("Content-type", "text/html")
		c.String(200, s)
		c.Next()
	}
}

// RenderJadeFromDirectPath renders a preset jade file for a route
func RenderJadeFromDirectPath(basepath, path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		s, _ := RenderJade(basepath, path)
		c.Header("Content-type", "text/html")
		c.String(200, s)
		c.Next()
	}
}
