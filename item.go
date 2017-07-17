package osdi

type Collection struct {
	TotalPages   int64                  `json:"total_pages"`
	PerPage      int64                  `json:"per_page"`
	Page         int64                  `json:"page"`
	TotalRecords int64                  `json:"total_records"`
	Links        LinkMap                `json:"_links"`
	Embedded     map[string]interface{} `json:"_embedded"`
}

type Item struct {
	CommonFields
	OriginSystem string  `json:"origin_system"`
	ItemType     string  `json:"item_type"`
	Links        LinkMap `json:"_links"`
}
