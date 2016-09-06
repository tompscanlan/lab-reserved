package labreserved

import (
	"github.com/go-openapi/swag"
	"github.com/tompscanlan/labreserved/restapi/operations"
)

var BlobEndpoint = "http://blobs.vmwaredevops.appspot.com/api/v1/blobs"
var BlobID = 10

func AddFlags(api *operations.LabreservedAPI) {
	var BlobFlags struct {
		Endpoint func(string) `short:"b" long:"blob-endpoint" description:"endpoint to blob storage" default:"http://blobs.vmwaredevops.appspot.com/api/v1/blobs"`
		Id       func(int)    `short:"i" long:"blob-id" description:"id of blob to set" default:"9999" optional:"false"`
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
