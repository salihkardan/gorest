
https://circleci.com/gh/salihkardan/gorest/tree/master.svg?style=shield&circle-token=f5ef2ac0205609236599e34dba8aa957e3e12912

**How-To**

To run webserver, you need to run `go run restful.go` command. It wil start webserver, please go to `localhost:8080`.

I have created tabs on main UI:
 * Events - shows events saved in Cassandra
 * Monitor - shows number of request according to their response timestamp
 * Responses - shows saved responses with their response times

There are 3 end points:
 * GET /api/events - lists all saved events
 * GET /api/requests - lists all saved requests
 * POST /api/endpoint - consumes incoming events

**Notes**

1) I have used Gin web framework: `https://github.com/gin-gonic/gin`

2) I assumed coming requests will be in JSON format. Here is a example request sample I tested the project:

  		curl -X POST -H "Cache-Control: no-cache" -d '{
          "apiKey": "test-api-key2",
          "userID": "'"$userId"'",
          "timestamp": '$time'
 	    }' "http://localhost:8080/api/endpoint"

  Note: There is "run.sh" script for creating sample http requests for testing purpose.

3) I have defined some sample api keys in apikey.txt file. While starting webserver, I load contents of that file, used those predefined api keys during validation of users.

4) I have use glide as dependency manager (glide.yaml), to install dependencies run `glide install` command. To install glide refer here : `https://github.com/Masterminds/glide`

5) I have used jade template engine to create html pages: `http://jade-lang.com/`

6) All front end related codes under web/ folder.

7) For database I have use Cassandra 2.X version. There is a file under root directory of project which contains
table schema of Cassandra. To create those tables, you need to run `cqlsh -f db.cql` command.
