package labreserved

import (
	"os"
	"strconv"

	"github.com/go-openapi/swag"
	"github.com/tompscanlan/labreserved/restapi/operations"
)

// set some defaults
var BlobEndpoint = "http://blobs.vmwaredevops.appspot.com/api/v1/blobs"
var BlobID = 7357

func init() {
	// allow defaults to be overridden from ENV
	env := os.Getenv("BLOB_ID")
	if env != "" {
		i, err := strconv.Atoi(env)
		if err != nil {
			panic(err)
		}
		BlobID = i
	}
	env = os.Getenv("BLOB_ENDPOINT")
	if env != "" {
		BlobEndpoint = env
	}

}

// Also allow params to override defaults and ENV
func AddFlags(api *operations.LabreservedAPI) {
	var BlobFlags struct {
		Endpoint func(string) `short:"b" long:"blob-endpoint" description:"endpoint to blob storage" default:"http://blobs.vmwaredevops.appspot.com/api/v1/blobs"`
		Id       func(int)    `short:"i" long:"blob-id" description:"id of blob to set" optional:"false"`
	}

	BlobFlags.Id = func(id int) {
		BlobID = id
		return
	}

	BlobFlags.Endpoint = func(endpoint string) {
		BlobEndpoint = endpoint
		return
	}
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }

	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{"blob", "blob storage configuration", &BlobFlags},
	}
}
