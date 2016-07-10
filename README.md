How-To

To run webserver, you need to run `go run restful.go` command. It wil start webserver, please go to `localhost:8080`.

I have created tabs on main UI:
  1) Events - shows events saved in Cassandra
  2) Monitor - shows number of request according to their response timestamp
  3) Responses - shows saved responses with their response time

For database, I have chosen Cassandra and use 2.X version of it. There is a file under root directory of project which contains
table schema of Cassandra. To create those tables, you need to run `cqlsh -f db.cql` command.


Assumptions:

1) I assumed coming requests will be in JSON format. Here is a example request sample I tested the project:

  		curl -X POST -H "Cache-Control: no-cache" -d '{
          "apiKey": "test-api-key2",
          "userID": "'"$userId"'",
          "timestamp": '$time'
 	    }' "http://localhost:8080/api/endpoint"
