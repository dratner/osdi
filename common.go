package osdi

import (
	"time"
)

type CommonFields struct {
	Identifiers  []string   `json:"identifiers"`
	CreatedDate  *time.Time `json:"created_date"`
	ModifiedDate *time.Time `json:"modified_date"`
}

type Collection struct {
	TotalPages   int64                  `json:"total_pages"`
	PerPage      int64                  `json:"per_page"`
	Page         int64                  `json:"page"`
	TotalRecords int64                  `json:"total_records"`
	Links        Links                  `json:"_links"`
	Embedded     map[string]interface{} `json:"_embedded"`
}

type Links map[string]interface{}
