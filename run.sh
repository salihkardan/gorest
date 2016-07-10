#!/bin/bash

# sample request creater ....
userId=$1
echo $userId

if [ -z "$userId" ]
  then
	userId="default"
        echo "Using default userId while sending requests..."
fi

for i in {0..10}
do
time=`echo $(($(date +'%s * 1000 + %-N / 1000000')))`
a=`echo $RANDOM | tr '[0-9]' '[a-zA-Z]'`

curl -X POST -H "Cache-Control: no-cache" -d '{
        "apiKey": "test-api-key2",
        "userID": "'"$userId"'",
        "timestamp": '$time' 
}' "http://localhost:8080/api/endpoint"

sleep 1

done
