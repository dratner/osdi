package osdi

import (
	"time"
)

type Canvass struct {
	CommonFields
	OriginSystem string     `json:"origin_system"`
	ActionDate   *time.Time `json:"action_date"`
	ContactType  string     `json:"contact_type"`
	InputType    string     `json:"input_type"`
	Success      bool       `json:"success"`
	StatusCode   string     `json:"status_code"`
	Links        Links      `json:"_links"`
}
