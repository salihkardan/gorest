package cassandra

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
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
var cluster = gocql.NewCluster("localhost")

// GetEventsFromCassandra get evetns form Cassandra
func GetEventsFromCassandra() gin.HandlerFunc {
	return func(c *gin.Context) {
		cluster.Keyspace = keyspace
		cluster.Consistency = gocql.Quorum
		session, _ := cluster.CreateSession()

		defer session.Close()
		iter := session.Query(`SELECT * FROM events`).Iter()

		var events []Event
		var event Event
		for iter.Scan(&event.APIKey, &event.UserID, &event.Timestamp) {
			events = append(events, event)
		}
		if err := iter.Close(); err != nil {
			log.Fatal(err)
		}
		c.JSON(200, events)
	}
}

// GetResponseTimesFromCassandra get response times form Cassandra
func GetResponseTimesFromCassandra() gin.HandlerFunc {
	return func(c *gin.Context) {
		cluster.Keyspace = keyspace
		cluster.Consistency = gocql.Quorum
		session, _ := cluster.CreateSession()

		defer session.Close()
		iter := session.Query(`SELECT * FROM responses`).Iter()

		var responses []Response
		var response Response

		for iter.Scan(&response.APIKey, &response.Timestamp, &response.Duration) {
			responses = append(responses, response)
		}
		if err := iter.Close(); err != nil {
			log.Fatal(err)
		}
		c.JSON(200, responses)
	}
}

// SaveEventsToCassandra saves events to Cassandra
func SaveEventsToCassandra() gin.HandlerFunc {
	return func(c *gin.Context) {
		cluster.Keyspace = "peak"
		cluster.Consistency = gocql.Quorum
		session, _ := cluster.CreateSession()
		defer session.Close()

		var json Event
		json, ok := c.MustGet("event").(Event)
		if !ok {
			// handle error here...
		}

		if err := session.Query(`INSERT INTO events (api_key, user_id, timestamp) VALUES(?, ?, ?)`,
			json.APIKey, json.UserID, json.Timestamp).Exec(); err != nil {
			c.JSON(400, "Could not save event....")
			log.Fatal(err)
		} else {
			// sleep randomly
			start := time.Now()
			time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)
			elapsed := time.Since(start)
			fmt.Println(elapsed)
			temp := elapsed.Nanoseconds() / int64(time.Millisecond)
			// fmt.Printf("temp : %#v\n", temp)
			SaveRequestToCassandra(session, json.APIKey, temp)
			// save request to another table
			s10 := strconv.FormatInt(temp, 10)
			c.JSON(200, "Event successfully saved... Slept "+s10+"ms.")
		}
	}
}

// SaveRequestToCassandra save reqessts to Cassandra
func SaveRequestToCassandra(session *gocql.Session, apiKey string, temp int64) {
	if err := session.Query(`INSERT INTO responses (api_key, response_time, timestamp) VALUES(?, ?, ?)`,
		apiKey, temp, time.Now()).Exec(); err != nil {
		fmt.Println("err:", err)
		log.Panicln(err)
	}
}
