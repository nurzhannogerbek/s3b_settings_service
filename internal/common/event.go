package common

import "encoding/json"

type Event struct {
	Arguments json.RawMessage `json:"arguments"`
}