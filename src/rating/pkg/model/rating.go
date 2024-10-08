package model

// RecordID defines a record id. Together with RecordType identifies unique records across all types.
type RecordID string

// RecordType defines a record type. Together with RecordID identifies unique records across all types.
type RecordType string

// RecordTypeMovie Existing record types.
const (
	RecordTypeMovie = RecordType("movie")
)

// UserID defines a user id.
type UserID string

// RatingValue defines a value of a rating record.
type RatingValue int

// Rating defines an individual rating created by a user for some record.
type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}

// RatingEvent defines an event containing raring information.
type RatingEvent struct {
	UserID     UserID          `json:"user_id"`
	RecordID   RecordID        `json:"record_id"`
	RecordType RecordType      `json:"record_type"`
	Value      RatingValue     `json:"value"`
	EventType  RatingEventType `json:"event_type"`
}

// RatingEventType defines the type of rating event.
type RatingEventType string

// Rating event types.
const (
	RatingEventTypePut    = "put"
	RatingEventTypeDelete = "delete"
)
