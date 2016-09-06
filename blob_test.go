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
  		"id": "xxxx",
  		"version": "v1.0",
  		"name": "Our team #10 blob",
  		"content": "Wooot",
  		"tag": ""
		}`

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
