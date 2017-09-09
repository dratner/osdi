package osdi

import (
	"time"
)

type Donation struct {
	CommonFields
	OriginSystem         string        `json:"origin_system"`
	ActionDate           *time.Time    `json:"action_date"`
	Amount               float64       `json:"amount"`
	CreditedAmount       float64       `json:"credited_amount"`
	CreditedDate         *time.Time    `json:"credited_date"`
	Currency             string        `json:"currency"`
	Recipients           []Recipient   `json:"recipients"`
	Payment              *Payment      `json:"payment"`
	SubscriptionInstance string        `json:"subscription_instance"`
	Voided               bool          `json:"voided"`
	VoidedDate           *time.Time    `json:"voided_date"`
	Url                  string        `json:"url"`
	ReferrerData         *ReferrerData `json:"referrer_data"`
	Links                Links         `json:"_links"`
}

type Recipient struct {
	DisplayName string  `json:"display_name"`
	LegalName   string  `json:"legal_name"`
	Amount      float64 `json:"amount"`
}

type Payment struct {
	Method              string `json:"method"`
	ReferenceNumber     string `json:"reference_number"`
	AuthorizationStored bool   `json:"authorization_stored"`
}

type ReferrerData struct {
	Source   string `json:"source"`
	Referrer string `json:"referrer"`
	Website  string `json:"website"`
	Url      string `json:"url"`
}
