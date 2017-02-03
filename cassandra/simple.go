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
		c.JSON(200, "Hello")
	}
}

// GetResponseTimesFromCassandra get response times form Cassandra
func GetResponseTimesFromCassandra() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Hello")
	}
}

// SaveEventsToCassandra saves events to Cassandra
func SaveEventsToCassandra() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Event successfully saved...")
	}
}
