CREATE TABLE foo (
 id Int64,
 event_time DateTime
)
Engine=MergeTree()
PARTITION BY toYYYYMMDD(event_time)
ORDER BY id;