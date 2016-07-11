package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strconv"
	"strings"

	cassandra "gorest/cassandra"
	jade "gorest/jade"

	"github.com/gin-gonic/gin"
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

	c.JSON(code, resp)
	c.Abort()
}

//TokenAuthMiddleware apiKey authentaication
func TokenAuthMiddleware(apiKeyMap map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(apiKeyMap)
		var json cassandra.Event
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

func main() {
	apiKeyMap := loadAPIKeys()

	port := 8080
	StaticFiles := "./web"

	r := gin.New()
	r.Use(gin.Recovery())
	// logger := gin.Logger()
	gin.SetMode(gin.ReleaseMode)

	// Static files mapping
	r.Static("/css", path.Join(StaticFiles, "/public/css"))
	r.Static("/fonts", path.Join(StaticFiles, "/public/fonts"))
	r.Static("/js", path.Join(StaticFiles, "/public/js"))
	r.Static("/vendor", path.Join(StaticFiles, "/public/vendor"))

	// Jade renderings
	r.GET("/", jade.RenderJadeFromDirectPath(path.Join(StaticFiles, "/views"), "index.html"))
	r.GET("/partials/*filepath", jade.RenderJadeFromBasePath(path.Join(StaticFiles, "/views/partials")))

	r.GET("/api/events", cassandra.GetEventsFromCassandra())
	r.GET("/api/requests", cassandra.GetResponseTimesFromCassandra())

	api := r.Group("/api", TokenAuthMiddleware(apiKeyMap))
	{
		api.POST("/endpoint", cassandra.SaveEventsToCassandra())
	}

	r.Run(":" + strconv.Itoa(port))
}
