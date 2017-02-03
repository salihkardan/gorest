package cassandra

import "github.com/gin-gonic/gin"

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

// GetEventsFromCassandra get evetns form Cassandra
func GetEventsFromCassandra() gin.HandlerFunc {
	return func(c *gin.Context) {
		var events []Event
		var event Event
		event.APIKey = "test"
		event.UserID = "1"
		event.Duration = 121
		event.Timestamp = 123123
		events = append(events, event)

		c.JSON(200, events)
	}
}

// GetResponseTimesFromCassandra get response times form Cassandra
func GetResponseTimesFromCassandra() gin.HandlerFunc {
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

// SaveEventsToCassandra saves events to Cassandra
func SaveEventsToCassandra() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Event successfully saved...")
	}
}
