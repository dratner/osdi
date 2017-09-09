package osdi

type List struct {
	CommonFields
	OriginSystem      string `json:"origin_system"`
	Name              string `json:"name"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	Summary           string `json:"summary"`
	BrowserUrl        string `json:"browser_url"`
	AdministrativeUrl string `json:"AdministrativeUrl"`
	TotalItems        int64  `json:"total_items"`
	Links             Links  `json:"_links"`
}
