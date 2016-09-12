package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/tompscanlan/labreserved/models"
)

// NewPostUserParams creates a new PostUserParams object
// with the default values initialized.
func NewPostUserParams() *PostUserParams {
	var ()
	return &PostUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserParamsWithTimeout creates a new PostUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUserParamsWithTimeout(timeout time.Duration) *PostUserParams {
	var ()
	return &PostUserParams{

		timeout: timeout,
	}
}

/*PostUserParams contains all the parameters to send to the API endpoint
for the post user operation typically these are written to a http.Request
*/
type PostUserParams struct {

	/*Adduser
	  representation of the user to add

	*/
	Adduser *models.User

	timeout time.Duration
}

// WithAdduser adds the adduser to the post user params
func (o *PostUserParams) WithAdduser(Adduser *models.User) *PostUserParams {
	o.Adduser = Adduser
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Adduser == nil {
		o.Adduser = new(models.User)
	}

	if err := r.SetBodyParam(o.Adduser); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}