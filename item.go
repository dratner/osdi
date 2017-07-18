package osdi

type Item struct {
	CommonFields
	OriginSystem string `json:"origin_system"`
	ItemType     string `json:"item_type"`
	Links        Links  `json:"_links"`
}
