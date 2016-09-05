package labreserved

import (
	"log"

	"github.com/tompscanlan/labreserved/models"
)

var AllItems models.Items
var AllUsers models.Users

func init() {
	log.Println("init labreserved")
	AllItems = models.NewItems()
	AllUsers = models.NewUsers()
	//	_ = AllItems
	//	_ = AllUsers

	AllItems["server1"] = models.NewItem("server1", "testing")
}
