package sophos

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// DefaultHTTPClient is the default http.Client used. Caller can modify Client (e.g. to allow SkipInsecure)
var DefaultHTTPClient HTTPClient

// HttpClient is an interface which represents an http.Client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func init() { DefaultHTTPClient = &http.Client{} }

// ClientInterface represents a Sophos 9 REST API client
type ClientInterface interface {
	Get(path string, options ...Option) (*Response, error)
	Put(path string, body io.Reader, options ...Option) (*Response, error)
	Post(path string, body io.Reader, options ...Option) (*Response, error)
	Patch(path string, body io.Reader, options ...Option) (*Response, error)
	Delete(path string, options ...Option) (*Response, error)
}

// ObjectClient is an interface with syntactic sugar dealing with Objects
type ObjectClient interface {
	GetObject(o RestGetter, options ...Option) error
	PutObject(o RestObject, options ...Option) error
	PatchObject(o RestObject, options ...Option) error
	PostObject(o RestObject, options ...Option) error
	DeleteObject(o RestObject, options ...Option) error
	GetUsedBy(o RestObject, options ...Option) (*UsedBy, error)
	GetEndpointSwag(e Endpoint, options ...Option) (Swag, error)
}

// Client implements ClientInterface to provide a REST client
type Client struct {
	endpoint string
	apiKey   string
	opts     []Option
}

var ensureInterface Client
var _ ClientInterface = ensureInterface
var _ ObjectClient = ensureInterface

// ErrRefRequired is an error that is returned when the resource requires a Referencee to fetch data
var ErrRefRequired = errors.New("client: Reference is required")

// Re-set DefaultHTTPClient for https errors.
// 
// Ex. Verify SSL error
//
// client, _ := sophos.New(...)
//
// client.SetDefaultHTTPClient(http.Client{
//	Transport: &http.Transport{
//		TLSClientConfig: &tls.Config{
//			InsecureSkipVerify: true,
//		},
//	},
//})
//
func (c Client) SetDefaultHTTPClient(cl http.Client)  {
	DefaultHTTPClient = &cl
}

// New returns a new Client.
// The endpoint provided should point to the Sophos Gateway.
func New(endpoint string, opts ...Option) (*Client, error) {
	if endpoint == "" {
		return nil, errors.New("endpoint is required")
	}

	if !strings.HasPrefix(endpoint, "http:") && !strings.HasPrefix(endpoint, "https:") {
		endpoint = "https://" + endpoint
	}

	return &Client{
		endpoint: endpoint,
		opts:     opts,
	}, nil
}

// Do executes the call and returns a *Response
func (c Client) Do(method, path string, body io.Reader, options ...Option) (resp *Response, err error) {
	resp = &Response{Response: &http.Response{}}
	resp.Request, err = c.Request(method, path, body, options...)
	if err != nil {
		return
	}

	res, err := DefaultHTTPClient.Do(resp.Request)
	if err != nil {
		return
	}

	resp.Response = res
	if res.StatusCode >= 400 && res.StatusCode <= 422 {
		err = fmt.Errorf("client do: error from server: %s", res.Status)

		// check for Errors
		var ee Errors
		_ = resp.MarshalTo(&ee)
		if len(ee) > 0 {
			resp.Errors = &ee
			err = resp.Errors
		}

		return
	}

	if !(res.StatusCode >= 200 && res.StatusCode <= 204) {
		err = fmt.Errorf("client do: error from server: %s", res.Status)
	}

	return
}

// Delete executes a DELETE call
//
// You can use DELETE requests to destroy object resources that were created with
// POST or PUT requests.
// Confd may deny DELETE requests due to validation checks. To use confd auto
// correction, use the special headers described X-Restd-Err-Ack header.
//
// DELETE /api/objects/packetfilter/packetfilter/REF_PacPacAllowAnyFTPOut
func (c Client) Delete(path string, options ...Option) (*Response, error) {
	return c.Do(http.MethodDelete, path, nil, options...)
}

// Get requests are used to retrieve information. The GET request cannot modify the data from confd.
//
// Examples for GET calls:
// GET /api/nodes
// GET /api/nodes/webadmin.port
// GET /api/objects/network/network
// GET /api/objects/network/network/REF_NetNet100008
func (c Client) Get(path string, options ...Option) (*Response, error) {
	return c.Do(http.MethodGet, path, nil, options...)
}

// Put executes a PUT call
//
// You can use PUT and POST for creating new resources. POST will create a new
// resource with an auto generated REF_ string whereas PUT will create resource for the
// REF_ string. You can use PUT to update the same resource after creation. PUT and
// POST require that you set all mandatory attributes of an object or node. You may need
// to override changes by removing locks (see XRestdLockOverride header).
//
// PUT /api/nodes/packetfilter.rules
// POST /api/objects/packetfilter/packetfilter
// PUT /api/objects/packetfilter/packetfilter/REF_PacPacAllowAnyFTPOut
func (c Client) Put(path string, body io.Reader, options ...Option) (*Response, error) {
	r, err := c.Do(http.MethodPut, path, body, options...)
	if err != nil {
		return r, err
	}

	return r, nil
}

// Post executes a POST call
//
// You can use PUT and POST for creating new resources. POST will create a new
// resource with an auto generated REF_ string whereas PUT will create resource for the
// REF_ string. You can use PUT to update the same resource after creation. PUT and
// POST require that you set all mandatory attributes of an object or node. You may need
// to override changes by removing locks (see chapter X-Restd-Lock-Override header).
//
// PUT /api/nodes/packetfilter.rules
// POST /api/objects/packetfilter/packetfilter
// PUT /api/objects/packetfilter/packetfilter/REF_PacPacAllowAnyFTPOut
func (c Client) Post(path string, body io.Reader, options ...Option) (*Response, error) {
	return c.Do(http.MethodPost, path, body, options...)
}

// Patch executes a PATCH call
//
// You can use PATCH requests to update fields on an existing object:
// PATCH /api/objects/packetfilter/packetfilter/REF_PacPacAllowAnyFTPOut
func (c Client) Patch(path string, body io.Reader, options ...Option) (*Response, error) {
	return c.Do(http.MethodPatch, path, body, options...)
}

// Version contains the UTMs version data
type Version struct {
	UTM   string `json:"utm"`
	Restd string `json:"restd"`
}

// Ping the gateway to retrieve its versioning
func (c Client) Ping(options ...Option) (v *Version, err error) {
	r, err := c.Get("/api/status/version", options...)
	if err != nil {
		return nil, fmt.Errorf("ping: error retrieving version from gateway: %s", err.Error())
	}

	v = &Version{}
	err = r.MarshalTo(v)

	return
}

// Endpoint returns the client's UTM endpoint
func (c Client) Endpoint() string {
	return c.endpoint
}

// Request generates a new *http.Request that is modified with the Client's Options (set on New)
// and with the provided Options.
func Request(method, path string, body io.Reader, options ...Option) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, fmt.Errorf("request: error generating new http.Request: %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	if err := evaluateOpts(req, options); err != nil {
		return req, fmt.Errorf("request: error evaluation Options: %s", err.Error())
	}

	return req, nil
}

// Request generates a new *http.Request that is modified with the Client's Options (set on New)
// and with the provided Options.
func (c *Client) Request(method, path string, body io.Reader, options ...Option) (*http.Request, error) {
	return Request(method, c.endpoint+path, body, append(c.opts, options...)...)
}

// GetObject implements TypeClient
func (c Client) GetObject(o RestGetter, options ...Option) error {
	if ref, required := o.RefRequired(); required && ref == "" {
		return ErrRefRequired
	}
	res, err := c.Get(o.GetPath(), options...)
	if err != nil {
		return err
	}
	err = res.MarshalTo(o)
	return err
}

// PostObject POSTs the RestObject
func (c Client) PostObject(o RestObject, options ...Option) error {
	byt, _ := json.Marshal(o)
	res, err := c.Post(o.PostPath(), bytes.NewReader(byt), options...)

	if res.StatusCode == http.StatusCreated {
		// Operation successful and created a new resource. The newly created
		// resource and its path and REF_ string are returned in the
		// Location header
		err = res.MarshalTo(o)
	}

	return err
}

// PatchObject PATCHes the RestObject
func (c Client) PatchObject(o RestObject, options ...Option) error {
	ref, required := o.RefRequired()
	if required && ref == "" {
		return ErrRefRequired
	}
	byt, _ := json.Marshal(o)
	_, err := c.Patch(o.PatchPath(ref), bytes.NewReader(byt), options...)
	return err
}

// PutObject PUTs the RestObject
func (c Client) PutObject(o RestObject, options ...Option) error {
	ref, required := o.RefRequired()
	if required && ref == "" {
		return ErrRefRequired
	}
	byt, _ := json.Marshal(o)
	_, err := c.Put(o.PutPath(ref), bytes.NewReader(byt), options...)
	return err
}

// DeleteObject DELETEs the RestObject
func (c Client) DeleteObject(o RestObject, options ...Option) error {
	ref, required := o.RefRequired()
	if required && ref == "" {
		return ErrRefRequired
	}
	_, err := c.Delete(o.DeletePath(ref), options...)
	return err
}

// GetEndpointSwag is syntactic sugar around an Endpoint Definition's GetSwag method.
func (c Client) GetEndpointSwag(e Endpoint, options ...Option) (Swag, error) {
	def := e.Definition()
	return def.GetSwag(&c)
}

// GetUsedBy GETs the Objects UsedBy data
func (c Client) GetUsedBy(o RestObject, options ...Option) (*UsedBy, error) {
	ref, required := o.RefRequired()
	if required && ref == "" {
		return nil, ErrRefRequired
	}
	var usedBy UsedBy
	resp, err := c.Get(o.UsedByPath(ref), options...)
	if err != nil {
		return nil, err
	}
	err = resp.MarshalTo(&usedBy)
	return &usedBy, err
}
