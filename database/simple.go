package database

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v5"
	"github.com/op/go-logging"
	"time"
	"strconv"
)

// Event struct to bind objects
type Event struct {
	APIKey    string `form:"apiKey" json:"apiKey" binding:"required"`
	UserID    string `form:"userID" json:"userID" binding:"required"`
}

type IncomingRequest struct {
	Ip            string `form:"ip" json:"ip" binding:"required"`
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

// GetEvents get evetns form Cassandra
func GetEvents() gin.HandlerFunc {
	return func(c *gin.Context) {
		var events []Event
		var event Event
		keys := redisCli.Keys("*").Val();
		for i := 0; i < len(keys); i++ {
			val := redisCli.Get(keys[i]).Val();
			event.APIKey = keys[i];
			event.UserID = val;
			events = append(events, event)
		}
		c.JSON(200, events)
	}
}

// GetResponseTimes get response times
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

func SaveVisitors(visitor IncomingRequest) {
	currentTime := time.Now().Format(time.RFC850);
	//strconv.FormatInt(time.Now().UnixNano(), 10)
	log.Info("Time", currentTime)
	redisCli.Set(currentTime, visitor.Ip, 0);
}

func SaveEvents(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info(c)
		go Save(apiKey)
	}
}

func SaveIncomingRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json IncomingRequest
		if c.Bind(&json) == nil {
			go SaveVisitors(json)
		} else {
			c.String(400, "failed")
		}
	}
}

