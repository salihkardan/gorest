package database

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v5"
	"fmt"
	"github.com/op/go-logging"
	"time"
	"strconv"
)

// Event struct to bind objects
type Event struct {
	APIKey    string `form:"apiKey" json:"apiKey" binding:"required"`
	UserID    string `form:"userID" json:"userID" binding:"required"`
	Timestamp int64  `form:"timestamp" json:"timestamp" binding:"required"`
	Duration  int64
}

// Response struct to bind objects
type Response struct {
	APIKey    string
	Duration  int64
	Timestamp int64
}

var keyspace = "peak"
var log = logging.MustGetLogger("example")
var redisCli = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func ExampleRedisSetGet() {
	redisCli.Set("test", "salih2", 0);
	log.Info("setting key")
	value, _ := redisCli.Get("test").Result();
	log.Info(value);
	pong, _ := redisCli.Ping().Result()
	fmt.Println(pong)
}

// GetEvents get evetns form Cassandra
func GetEvents() gin.HandlerFunc {
	return func(c *gin.Context) {
		var events []Event
		var event Event

		keys := redisCli.Keys("*").Val();

		for i := 0; i < len(keys); i++ {
			event.APIKey = keys[i];
			event.UserID = "1"
			event.Duration = 121
			event.Timestamp = time.Now().Unix();
			events = append(events, event)
		}
		c.JSON(200, events)
	}
}

// GetResponseTimes get response times form Cassandra
func GetResponseTimes() gin.HandlerFunc {
	return func(c *gin.Context) {
		var responses []Response
		var response Response

		response.APIKey = "test"
		response.Duration = 121
		response.Timestamp = 123123
		responses = append(responses, response)

		c.JSON(200, responses)
	}
}

// SaveEvents saves events to Cassandra
func Save(apiKey string) {
	currentTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	log.Info("Time", currentTime)
	log.Info("ApiKey", apiKey)
	redisCli.Set(currentTime+"_"+apiKey, apiKey, 0);
	log.Info("Saving key : ", currentTime+"_test")
}

func SaveEvents(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info(c)
		go Save(apiKey)
	}
}

func SaveIncomingRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		type IncomingRequest struct {
			Visitor       string `form:"visitor" 		json:"visitor" 			binding:"required"`
			Ip            string `form:"ip" 		json:"ip" 			binding:"required"`
		}
		var json IncomingRequest
		if c.Bind(&json) == nil {
			go Save(json.Ip)
		} else {
			c.String(400, "failedd")
		}
	}
}

