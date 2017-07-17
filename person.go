package osdi

import (
	"time"
)

type CommonFields struct {
	Identifiers  []string   `json:"identifiers"`
	CreatedDate  *time.Time `json:"created_date"`
	ModifiedDate *time.Time `json:"modified_date"`
}

type Birthdate struct {
	Month int `json:"month"`
	Day   int `json:"day"`
	Year  int `json:"year"`
}

type Party struct {
	Identification   string     `json:"identification"`
	LastVerifiedDate *time.Time `json:"last_verified_date"`
	Active           bool       `json:"active"`
}

type LocationHash struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Accuracy  string  `json:"accuract"` // Enum
}

type EmployerAddress struct {
	Venue        string       `json:"venue"`
	AddressLines []string     `json:"address_lines"`
	Locality     string       `json:"locality"`
	Region       string       `json:"region"`
	PostalCode   string       `json:"postal_code"`
	Country      string       `json:"country"`
	Language     string       `json:"string"`
	Location     LocationHash `json:"location"`
	Status       string       `json:"status"` // Enum
}

type PostalAddress struct {
	Primary     bool   `json:"primary"`
	AddressType string `json:"address_type"` // Enum
	EmployerAddress
	LastVerifiedDate *time.Time `json:"last_verified_date"`
	Occupation       string     `json:"occupation"`
}

type EmailAddress struct {
	Primary     bool   `json:"primary"`
	Address     string `json:"address"`
	AddressType string `json:"address_type"` //Flexenum
	Status      string `json:"status"`       // Enum
}

type PhoneNumber struct {
	Primary     bool   `json:"primary"` // Documentation says string ... error?
	Number      string `json:"number"`
	Extension   string `json:"extension"`
	Description string `json:"description"`
	NumberType  string `json:"number_type"` //Flexenum
	Operator    string `json:"operator"`
	Country     string `json:"country"`
	SMSCapable  bool   `json:"sms_capable"`
	DoNotCall   bool   `json:"do_not_call"`
}

type Profile struct {
	Provider string `json:"provider"`
	Id       string `json:"id"`
	URL      string `json:"url"`
	Handle   string `json:"handle"`
}

type Person struct {
	CommonFields
	FamilyName          string            `json:"family_name"`
	GivenName           string            `json:"given_name"`
	AdditionalName      string            `json:"additional_name"`
	HonorificPrefix     string            `json:"honorific_prefix"`
	HonorificSuffix     string            `json:"honorific_suffix"`
	Gender              string            `json:"gender"` // Enum
	GenderIdentity      string            `json:"gender_identity"`
	PartyIdentification string            `json:"party_identification"`
	Parties             []Party           `json:"parties"`
	Source              string            `json:"source"`
	Ethnicities         []string          `json:"ethnicities"`
	LanguagesSpoken     []string          `json:"languages_spoken"`
	PreferredLanguage   string            `json:"preferred_language"`
	Birthdate           Birthdate         `json:"birthdate"`
	Employer            string            `json:"employer"`
	EmployerAddress     EmployerAddress   `json:"employer_address"`
	PostalAddresses     []PostalAddress   `json:"postal_addresses"`
	EmailAddresses      []EmailAddress    `json:"email_addresses"`
	PhoneNumbers        []PhoneNumber     `json:"phone_numbers"`
	CustomFields        map[string]string `json:"custom_fields"`
	Links               LinkMap           `json:"_links"`
}
