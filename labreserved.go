package labreserved

import (
	"fmt"
	"log"

	"github.com/tompscanlan/labreserved/models"
)

var AllItems models.Items
var AllUsers models.Users

func init() {
	log.Println("init labreserved")

	// read inventoy from the blob, and initialize our
	// lab equipment from that list

	log.Printf("loading inventory from blob #%d", BlobID)
	inventory, err := GetBlob(BlobID)
	if err != nil {
		panic(fmt.Sprintf("Failed to load blob: %s", err))
	}

	// make a new empty inventory
	AllItems = models.NewItems()

	// if inventory from blob exists, load up form there
	if inventory != "" {
		log.Printf("got non-nil inventory, starting with that")
		err = AllItems.LoadJSON(inventory)

		if err != nil {
			panic(fmt.Sprintf("Failed to load inventory: %s", err))
		}
	}

	AllUsers = models.NewUsers()

	// potentiall load test data here
	//	AllItems["server1"] = models.NewItem("server1", "testing")
}
