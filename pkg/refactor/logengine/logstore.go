package logengine

import (
	"context"
	"time"
)

// LogStore manages storing and querying LogEntries.
type LogStore interface {
	// appends a log entry to the logs. Passed keysAnValues will be associated with the log entry.
	Append(ctx context.Context, level string, msg string, keysAndValues map[string]interface{}) error
	// returns a limited number of log-entries that have matching associated fields with the provided keysAndValues pairs
	// starting a given offset. For no offset or unlimited log-entries in the result set the value to -1.
	// - To query server-logs pass: "recipientType", "server" via keysAndValues
	// - level SHOULD be passed as a string. Valid values are "debug", "info", "error", "panic".
	// - This method will search for any of followings keys and query all matching logs:
	// level, workflow_id, namespace_logs, log_instance_call_path, root_instance_id, mirror_activity_id
	// Any other not mentioned passed key value pair will be ignored.
	// Returned log-entries will have same or higher level as the passed one.
	// - Passing a log_instance_call_path will return all logs which have a callpath with the prefix as the passed log_instance_call_path value.
	// when passing log_instance_call_path the root_instance_id SHOULD be passed to optimize the performance of the query.
	Get(ctx context.Context, keysAndValues map[string]interface{}, limit, offset int) ([]*LogEntry, error)
}

// Represents an individual log entry.
type LogEntry struct {
	// the timestamp of the log-entry.
	T   time.Time
	Msg string
	// Fields contains metadata of the log-entry.
	Fields map[string]interface{}
}
