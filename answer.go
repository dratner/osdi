package osdi

import (
	"time"
)

type Answer struct {
	CommonFields
	OriginSystem string     `json:"origin_system"`
	ActionDate   *time.Time `json:"action_date"`
	Value        string     `json:"value"`
	Responses    []string   `json:"responses"`
	Links        Links      `json:"_links"`
}
