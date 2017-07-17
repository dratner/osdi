package osdi

import (
	"encoding/json"
	"testing"
    "log"
)

func TestItem(t *testing.T) {

	var iJson = []byte(`{
    "identifiers": [
        "osdi_sample_system:d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3",
        "foreign_system:1"
    ],
    "origin_system": "OSDI Sample System",
    "created_date": "2014-03-20T21:04:31Z",
    "modified_date": "2014-03-20T21:04:31Z",
    "item_type": "osdi:person",
    "_links": {
        "self": {
            "href": "https://osdi-sample-system.org/api/v1/lists/c945d6fe-929e-11e3-a2e9-12313d316c29/items/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3"
        },
        "osdi:list": {
            "href": "https://osdi-sample-system.org/api/v1/lists/c945d6fe-929e-11e3-a2e9-12313d316c29"
        },
        "osdi:person": {
            "href": "https://osdi-sample-system.org/api/v1/people/65345d7d-cd24-466a-a698-4a7686ef684f"
        }
    }
}`)

	log.Println("Test Item")

	var err error
	i := new(Item)
	err = json.Unmarshal(iJson, i)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestItems(t *testing.T) {

	var iJson = []byte(`{
    "total_pages": 10,
    "per_page": 25,
    "page": 1,
    "total_records": 250,
    "_links": {
        "next": {
            "href": "https://osdi-sample-system.org/api/v1/lists/c945d6fe-929e-11e3-a2e9-12313d316c29/items?page=2"
        },
        "osdi:items": [
            {
                "href": "https://osdi-sample-system.org/api/v1/lists/c945d6fe-929e-11e3-a2e9-12313d316c29/items/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3"
            },
            {
                "href": "https://osdi-sample-system.org/api/v1/lists/c945d6fe-929e-11e3-a2e9-12313d316c29/items/1efc3644-af25-4253-90b8-a0baf12dbd1e"
            }
        ],
        "curies": [
            {
                "name": "osdi",
                "href": "https://osdi-sample-system.org/docs/v1/{rel}",
                "templated": true
            }
        ],
        "self": {
            "href": "https://osdi-sample-system.org/api/v1/lists/c945d6fe-929e-11e3-a2e9-12313d316c29/items"
        }
    },
    "_embedded": {
        "osdi:items": [
            {
                "identifiers": [
                    "osdi_sample_system:d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3",
                    "foreign_system:1"
                ],
                "origin_system": "OSDI Sample System",
                "created_date": "2014-03-20T21:04:31Z",
                "modified_date": "2014-03-20T21:04:31Z",
                "item_type": "osdi:person",
                "_links": {
                    "self": {
                        "href": "https://osdi-sample-system.org/api/v1/lists/c945d6fe-929e-11e3-a2e9-12313d316c29/items/d91b4b2e-ae0e-4cd3-9ed7-d0ec501b0bc3"
                    },
                    "osdi:list": {
                        "href": "https://osdi-sample-system.org/api/v1/lists/c945d6fe-929e-11e3-a2e9-12313d316c29"
                    },
                    "osdi:person": {
                        "href": "https://osdi-sample-system.org/api/v1/people/65345d7d-cd24-466a-a698-4a7686ef684f"
                    }
                }
            },
            {
                "identifiers": [
                    "osdi_sample_system:1efc3644-af25-4253-90b8-a0baf12dbd1e"
                ],
                "origin_system": "OSDI Sample System",
                "created_date": "2014-03-20T20:44:13Z",
                "modified_date": "2014-03-20T20:44:13Z",
                "item_type": "osdi:event",
                "_links": {
                    "self": {
                        "href": "https://osdi-sample-system.org/api/v1/lists/c945d6fe-929e-11e3-a2e9-12313d316c29/items/1efc3644-af25-4253-90b8-a0baf12dbd1e"
                    },
                    "osdi:list": {
                        "href": "https://osdi-sample-system.org/api/v1/lists/c945d6fe-929e-11e3-a2e9-12313d316c29"
                    },
                    "osdi:event": {
                        "href": "https://osdi-sample-system.org/api/v1/events/adb951cb-51f9-420e-b7e6-de953195ec86"
                    }
                }
            }
        ]
    }
}`)

	log.Println("Test Items")

	var err error
	i := new(Collection)
	err = json.Unmarshal(iJson, i)

	if err != nil {
		t.Fatal(err.Error())
	}
}
