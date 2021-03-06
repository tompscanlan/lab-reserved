package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/tompscanlan/labreserved/models"
)

// GetItemsReader is a Reader for the GetItems structure.
type GetItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetItemsOK creates a GetItemsOK with default headers values
func NewGetItemsOK() *GetItemsOK {
	return &GetItemsOK{}
}

/*GetItemsOK handles this case with default header values.

list all items in the lab
*/
type GetItemsOK struct {
	Payload models.Items
}

func (o *GetItemsOK) Error() string {
	return fmt.Sprintf("[GET /items][%d] getItemsOK  %+v", 200, o.Payload)
}

func (o *GetItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
