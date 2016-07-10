CREATE KEYSPACE Peak
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 };

CREATE TABLE Peak.events (
  api_key text,
  user_id text,
  timestamp timestamp,
  PRIMARY KEY (api_key, user_id, timestamp)
)
WITH CLUSTERING ORDER BY (user_id ASC, timestamp ASC);
//
CREATE TABLE Peak.responses (
  api_key text,
  response_time int,
	timestamp timestamp,
  PRIMARY KEY (api_key, timestamp)
)
WITH CLUSTERING ORDER BY (timestamp ASC);
