package osdi

import (
	"encoding/json"
	"testing"
)

func TestPerson(t *testing.T) {

	var iJson = []byte(`{
    "identifiers": [
        "osdi_sample_system:d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3",
        "foreign_system:1"
    ],
    "origin_system": "OSDI Sample System",
    "created_date": "2014-03-20T21:04:31Z",
    "modified_date": "2014-03-20T21:04:31Z",
    "given_name": "John",
    "family_name": "Smith",
    "honorific_prefix": "Mr.",
    "honorific_suffix": "Ph.D",
    "additional_name": "Scott",
    "gender": "Male",
    "gender_identity": "Male",
    "party_identification": "Democratic",   
    "parties": [
      {
        "identification": "Democratic",
        "last_verified_date": "2014-03-20T21:04:31Z",
        "active": true
      }
    ],
    "source": "october_canvass",
    "birthdate": {
        "month": 6,
        "day": 2,
        "year": 1973
    },
    "ethnicities": [
        "African American"
    ],
    "languages_spoken": [
        "en",
        "fr"
    ],
    "employer": "Acme Corp",
    "employer_address": {
        "venue": "Bull Hall",
        "address_lines": [
            "123 Acme Street",
            "Suite 400"
        ],
        "locality": "New Yorkhaven",
        "region": "NY",
        "postal_code": "10001",
        "country": "US",
        "language": "en",
        "location": {
            "latitude": 38.9382,
            "longitude": -77.3349,
            "accuracy": "Rooftop"
        },
        "status": "Verified"
    },
    "occupation": "Accountant",
    "postal_addresses": [
        {
            "primary": true,
            "address_type": "Home",
            "address_lines": [
                "1900 Pennsylvania Ave"
            ],
            "locality": "Washington",
            "region": "DC",
            "postal_code": "20009",
            "country": "US",
            "language": "en",
            "location": {
                "latitude": 38.919,
                "longitude": -77.0379,
                "accuracy": "Rooftop"
            }
        }
    ],
    "email_addresses": [
        {
            "primary": true,
            "address": "johnsmith@mail.com",
            "address_type": "personal",
            "status": "subscribed"
        }
    ],
    "phone_numbers": [
        {
            "primary": true,
            "number": "11234567890",
            "extension": "432",
            "description": "Worksite line",
            "number_type": "Work",
            "operator": "ATT",
            "country": "US",
            "sms_capable": false,
            "do_not_call": true
        }
    ],
    "profiles": [
        {
            "provider": "Facebook",
            "id": "john.doe.1234",
            "url": "https://facebook.com/john.doe"
        },
        {
            "provider": "Twitter",
            "id": "eds34d8j2kddfd45",
            "url": "https://twitter.com/johndoe",
            "handle": "johndoe"
        }
    ],
    "custom_fields": {
        "is_volunteer": "true",
        "most_important_issue": "Equal pay",
        "union_member": "true"
    }, 
    "_links": {
        "self": {
            "href": "https://osdi-sample-system.org/api/v1/people/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3"
        },
        "osdi:answers": {
            "href": "https://osdi-sample-system.org/api/v1/people/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3/answers"
        },
        "osdi:attendance": {
            "href": "https://osdi-sample-system.org/api/v1/people/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3/attendance"
        },
        "osdi:signatures": {
            "href": "https://osdi-sample-system.org/api/v1/people/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3/signatures"
        },
        "osdi:submissions": {
            "href": "https://osdi-sample-system.org/api/v1/people/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3/submissions"
        },
        "osdi:donations": {
            "href": "https://osdi-sample-system.org/api/v1/people/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3/donations"
        },
        "osdi:outreaches": {
            "href": "https://osdi-sample-system.org/api/v1/people/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3/outreaches"
        },
        "osdi:taggings": {
            "href": "https://osdi-sample-system.org/api/v1/people/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3/taggings"
        },
        "osdi:items": {
            "href": "https://osdi-sample-system.org/api/v1/people/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3/items"
        },
        "osdi:record_canvass_helper": {
            "href": "https://osdi-sample-system.org/api/v1/people/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3/record_canvass_helper"
        }
    }
}`)

	var err error
	i := new(Person)
	err = json.Unmarshal(iJson, i)

	if err != nil {
		t.Fatal(err.Error())
	}

}
