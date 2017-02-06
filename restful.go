package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"

	"github.com/salihkardan/gorest/database"
	"github.com/salihkardan/gorest/jade"

	"github.com/gin-gonic/gin"
	"strconv"
)

//loadApiKeys load api keys to memory
func loadAPIKeys() map[string]bool {
	apiKeyMap := make(map[string]bool)
	fileName := "apikey.txt"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Could not ready " + fileName + " file. Exiting")
	}
	lines := strings.Split(string(content), "\n")
	for j := 0; j < len(lines); j++ {
		apiKeyMap[lines[j]] = true
	}
	return apiKeyMap
}

//RespondWithError resonses.
func RespondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	fmt.Print("Adfadf")
	c.JSON(code, resp)
	c.Abort()
}

//TokenAuthMiddleware apiKey authentaication
func TokenAuthMiddleware(apiKeyMap map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json database.Event
		if c.BindJSON(&json) == nil {
			if _, ok := apiKeyMap[json.APIKey]; ok {
				c.Set("event", json)
				c.Next()
			} else {
				log.Println("Could not authenticate user with API key")
				RespondWithError(401, "Could not authenticate user with API key", c)
			}
		} else {
			fmt.Println("Could not bind incoming hey data...")
			RespondWithError(400, "Wrong request... Missing data in request", c)
		}
	}
}

func Sum(x, y int) int {
	return x + y
}

func main() {
	apiKeyMap := loadAPIKeys()

	port := 8080
	StaticFiles := "./web"

	r := gin.New()
	r.Use(gin.Recovery())
	// logger := gin.Logger()
	gin.SetMode(gin.DebugMode)

	// Static files mapping
	r.Static("/css", path.Join(StaticFiles, "/public/css"))
	r.Static("/fonts", path.Join(StaticFiles, "/public/fonts"))
	r.Static("/js", path.Join(StaticFiles, "/public/js"))
	r.Static("/vendor", path.Join(StaticFiles, "/public/vendor"))
	r.Static("/images", path.Join(StaticFiles, "/public/images"))

	// Jade renderings
	r.GET("/", jade.RenderJadeFromDirectPath(path.Join(StaticFiles, "/views"), "index.html"))
	r.GET("/partials/*filepath", jade.RenderJadeFromBasePath(path.Join(StaticFiles, "/views/partials")))

	r.GET("/api/events", database.GetEvents())
	r.GET("/api/requests", database.GetResponseTimes())

	r.POST("/incoming", database.SaveIncomingRequest())

	api := r.Group("/api", TokenAuthMiddleware(apiKeyMap))
	{
		api.POST("/endpoint", database.SaveEvents("test-api-key"))
	}

	r.Run(":" + strconv.Itoa(port))
}
