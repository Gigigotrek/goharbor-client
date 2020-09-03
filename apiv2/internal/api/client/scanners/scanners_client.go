// Code generated by go-swagger; DO NOT EDIT.

package scanners

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new scanners API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scanners API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteScannersRegistrationID(params *DeleteScannersRegistrationIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteScannersRegistrationIDOK, error)

	PatchScannersRegistrationID(params *PatchScannersRegistrationIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchScannersRegistrationIDOK, error)

	PostScanners(params *PostScannersParams, authInfo runtime.ClientAuthInfoWriter) (*PostScannersCreated, error)

	PutProjectsProjectIDScanner(params *PutProjectsProjectIDScannerParams, authInfo runtime.ClientAuthInfoWriter) (*PutProjectsProjectIDScannerOK, error)

	PutScannersRegistrationID(params *PutScannersRegistrationIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutScannersRegistrationIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteScannersRegistrationID deletes a scanner registration

  Deletes the specified scanner registration.

*/
func (a *Client) DeleteScannersRegistrationID(params *DeleteScannersRegistrationIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteScannersRegistrationIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteScannersRegistrationIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteScannersRegistrationID",
		Method:             "DELETE",
		PathPattern:        "/scanners/{registration_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteScannersRegistrationIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteScannersRegistrationIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteScannersRegistrationID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchScannersRegistrationID sets system default scanner registration

  Set the specified scanner registration as the system default one.

*/
func (a *Client) PatchScannersRegistrationID(params *PatchScannersRegistrationIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchScannersRegistrationIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchScannersRegistrationIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchScannersRegistrationID",
		Method:             "PATCH",
		PathPattern:        "/scanners/{registration_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchScannersRegistrationIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchScannersRegistrationIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchScannersRegistrationID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostScanners creates a scanner registration

  Creats a new scanner registration with the given data.

*/
func (a *Client) PostScanners(params *PostScannersParams, authInfo runtime.ClientAuthInfoWriter) (*PostScannersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostScannersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostScanners",
		Method:             "POST",
		PathPattern:        "/scanners",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostScannersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostScannersCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostScanners: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutProjectsProjectIDScanner configures scanner for the specified project

  Set one of the system configured scanner registration as the indepndent scanner of the specified project.
*/
func (a *Client) PutProjectsProjectIDScanner(params *PutProjectsProjectIDScannerParams, authInfo runtime.ClientAuthInfoWriter) (*PutProjectsProjectIDScannerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutProjectsProjectIDScannerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutProjectsProjectIDScanner",
		Method:             "PUT",
		PathPattern:        "/projects/{project_id}/scanner",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutProjectsProjectIDScannerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutProjectsProjectIDScannerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutProjectsProjectIDScanner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutScannersRegistrationID updates a scanner registration

  Updates the specified scanner registration.

*/
func (a *Client) PutScannersRegistrationID(params *PutScannersRegistrationIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutScannersRegistrationIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutScannersRegistrationIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutScannersRegistrationID",
		Method:             "PUT",
		PathPattern:        "/scanners/{registration_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutScannersRegistrationIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutScannersRegistrationIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutScannersRegistrationID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
