#!/bin/bash

# sample request creater ....

for i in {0..10}
do

curl -X POST -H "Cache-Control: no-cache" -d '{
        "apiKey": "test-api-key2",
        "userID": "`echo $RANDOM | tr '[0-9]' '[a-zA-Z]'`",
        "timestamp": 1467331200
}' "http://localhost:8080/api/endpoint"

sleep 1

done
