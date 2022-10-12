package ed

import (
	"time"
)

// EventType is the type of any event, used as its unique identifier.
type EventType string

// Event is an internal representation of an event, returned when the Aggregate
// uses NewEvent to create a new event. The events loaded from the db is
// represented by each DBs internal event type, implementing Event.
type Event struct {
	EventID     string
	EventType   string
	Data        []byte
	Timestamp   time.Time
	AggregateID string
	Version     int64
	Metadata    []byte
}
