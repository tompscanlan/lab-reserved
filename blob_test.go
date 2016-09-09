package labreserved

import (
	"log"
	"testing"
)

var testBlobId = 7357

func TestBlobSet(t *testing.T) {

	if testing.Short() {
		return
	}

	testStr := `{
		"server11": {
			"description": "from booo!",
			"name": "server11",
			"reservations": [
			{
				"begin": "2016-09-06T00:09:04.032-04:00",
				"hoursheld": 3,
				"username": "tom"
			},
			{
				"begin": "2016-09-06T00:09:04.032-04:00",
				"hoursheld": 3,
				"username": "tom"
			}
			]
		},
		"server12": {
			"description": "from booo!",
			"name": "server12",
			"reservations": [
			{
				"begin": "2016-09-06T00:09:04.032-04:00",
				"hoursheld": 3,
				"username": "tom"
			}
			]
		},
		"server13": {
			"description": "from booo!",
			"name": "server13"
		},
		"server9": {
			"description": "from jasmine",
			"name": "server9"
		}
	} `

	err := PostBlob(testBlobId, testStr)
	if err != nil {
		t.Error(err)
	}

	blob, err := GetBlob(testBlobId)
	if err != nil {
		t.Error(err)
	}
	if testing.Verbose() {
		log.Println("Got decoded blob content: ", string(blob))
	}

	if blob != testStr {
		t.Errorf("blob(%s) doesn't match (%s)", blob, testStr)
	}

}
