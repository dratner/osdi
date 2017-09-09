package osdi

type FundraisingPage struct {
	CommonFields
	OriginSystem      string        `json:"origin_system"`
	Name              string        `json:"name"`
	Title             string        `json:"title"`
	Description       string        `json:"description"`
	Summary           string        `json:"summary"`
	BrowserUrl        string        `json:"browser_url"`
	AdministrativeUrl string        `json:"administrative_url"`
	FeaturedImageUrl  string        `json:"featured_image_url"`
	TotalDonations    int64         `json:"total_donations"`
	TotalAmount       float64       `json:"total_amount"`
	Currency          string        `json:"currency"`
	ShareUrl          string        `json:"share_url"`
	TotalShares       int64         `json:"total_shares"`
	ShareOptions      []ShareOption `json:"share_options"`
	Links             Links         `json:"_links"`
}

type ShareOption struct {
	FacebookShare FacebookShare `json:"facebook_share"`
	TwitterShare  TwitterShare  `json:"twitter_share"`
	EmailShare    EmailShare    `json:"email"shared"`
}

type FacebookShare struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	TotalShares int64  `json:"total_shares"`
}

type TwitterShare struct {
	Message     string `json:"message"`
	TotalShares int64  `json:"total_shares"`
}

type EmailShare struct {
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	TotalShares int64  `json:"total_shares"`
}
