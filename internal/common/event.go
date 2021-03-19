package common

import "encoding/json"

// Event
// Incoming request from AWS AppSync.
type Event struct {
	Arguments json.RawMessage `json:"arguments"`
}