// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// AmazonVpc is a generated struct representing the Sophos AmazonVpc Endpoint
// GET /api/nodes/amazon_vpc
type AmazonVpc struct {
	AutoPfrule  int64    `json:"auto_pfrule"`
	Connections []string `json:"connections"`
	Networks    []string `json:"networks"`
	Status      int64    `json:"status"`
}

var defsAmazonVpc = map[string]sophos.RestObject{
	"AmazonVpcGroup":      &AmazonVpcGroup{},
	"AmazonVpcTunnel":     &AmazonVpcTunnel{},
	"AmazonVpcConnection": &AmazonVpcConnection{},
}

// RestObjects implements the sophos.Node interface and returns a map of AmazonVpc's Objects
func (AmazonVpc) RestObjects() map[string]sophos.RestObject {
	return defsAmazonVpc
}

// GetPath implements sophos.RestGetter
func (*AmazonVpc) GetPath() string { return "/api/nodes/amazon_vpc" }

// RefRequired implements sophos.RestGetter
func (*AmazonVpc) RefRequired() (string, bool) { return "", false }

var defAmazonVpc = &sophos.Definition{Description: "amazon_vpc", Name: "amazon_vpc", Link: "/api/definitions/amazon_vpc", Swag: map[string]sophos.MethodMap{"/objects/amazon_vpc/connection/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object connection", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"amazon_vpc/connection"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available connection types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"amazon_vpc/connection"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object connection types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "amazon_vpc/connection that will be changes", Type: "", Required: true}}, Tags: []string{"amazon_vpc/connection"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object connection", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "amazon_vpc/connection that will be updated", Type: "", Required: true}}, Tags: []string{"amazon_vpc/connection"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}}, "/objects/amazon_vpc/connection/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"amazon_vpc/connection"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/amazon_vpc/group/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"amazon_vpc/group"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"amazon_vpc/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "amazon_vpc/group that will be changes", Type: "", Required: true}}, Tags: []string{"amazon_vpc/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "amazon_vpc/group that will be updated", Type: "", Required: true}}, Tags: []string{"amazon_vpc/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}}, "/objects/amazon_vpc/tunnel/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"amazon_vpc/tunnel"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/amazon_vpc/connection/": {"get": sophos.MethodDescriptions{Description: "Returns all available amazon_vpc/connection objects", Parameters: []sophos.Parameter(nil), Tags: []string{"amazon_vpc/connection"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}}}, "post": sophos.MethodDescriptions{Description: "Create a new amazon_vpc/connection object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "amazon_vpc/connection that will be created", Type: "", Required: true}}, Tags: []string{"amazon_vpc/connection"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/amazon_vpc/group/": {"get": sophos.MethodDescriptions{Description: "Returns all available amazon_vpc/group objects", Parameters: []sophos.Parameter(nil), Tags: []string{"amazon_vpc/group"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "post": sophos.MethodDescriptions{Description: "Create a new amazon_vpc/group object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "amazon_vpc/group that will be created", Type: "", Required: true}}, Tags: []string{"amazon_vpc/group"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}}}}, "/objects/amazon_vpc/group/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"amazon_vpc/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/amazon_vpc/tunnel/": {"get": sophos.MethodDescriptions{Description: "Returns all available amazon_vpc/tunnel objects", Parameters: []sophos.Parameter(nil), Tags: []string{"amazon_vpc/tunnel"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new amazon_vpc/tunnel object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "amazon_vpc/tunnel that will be created", Type: "", Required: true}}, Tags: []string{"amazon_vpc/tunnel"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}}}}, "/objects/amazon_vpc/tunnel/{ref}": {"put": sophos.MethodDescriptions{Description: "Creates or updates the complete object tunnel", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "amazon_vpc/tunnel that will be updated", Type: "", Required: true}}, Tags: []string{"amazon_vpc/tunnel"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object tunnel", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"amazon_vpc/tunnel"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available tunnel types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"amazon_vpc/tunnel"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object tunnel types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "amazon_vpc/tunnel that will be changes", Type: "", Required: true}}, Tags: []string{"amazon_vpc/tunnel"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}}}}

// Definition returns the /api/definitions struct of AmazonVpc
func (AmazonVpc) Definition() sophos.Definition { return *defAmazonVpc }

// ApiRoutes returns all known AmazonVpc Paths
func (AmazonVpc) ApiRoutes() []string {
	return []string{
		"/api/objects/amazon_vpc/connection/",
		"/api/objects/amazon_vpc/connection/{ref}",
		"/api/objects/amazon_vpc/connection/{ref}/usedby",
		"/api/objects/amazon_vpc/group/",
		"/api/objects/amazon_vpc/group/{ref}",
		"/api/objects/amazon_vpc/group/{ref}/usedby",
		"/api/objects/amazon_vpc/tunnel/",
		"/api/objects/amazon_vpc/tunnel/{ref}",
		"/api/objects/amazon_vpc/tunnel/{ref}/usedby",
	}
}

// References returns the AmazonVpc's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (AmazonVpc) References() []string {
	return []string{
		"REF_AmazonVpcConnection",
		"REF_AmazonVpcGroup",
		"REF_AmazonVpcTunnel",
	}
}

// AmazonVpcGroup is an Sophos Endpoint subType and implements sophos.RestObject
type AmazonVpcGroup []interface{}

// GetPath implements sophos.RestObject and returns the AmazonVpcGroup GET path
// Returns all available amazon_vpc/group objects
func (*AmazonVpcGroup) GetPath() string { return "/api/objects/amazon_vpc/group/" }

// RefRequired implements sophos.RestObject
func (*AmazonVpcGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AmazonVpcGroup DELETE path
// Creates or updates the complete object group
func (*AmazonVpcGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AmazonVpcGroup PATCH path
// Changes to parts of the object group types
func (*AmazonVpcGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AmazonVpcGroup POST path
// Create a new amazon_vpc/group object
func (*AmazonVpcGroup) PostPath() string {
	return "/api/objects/amazon_vpc/group/"
}

// PutPath implements sophos.RestObject and returns the AmazonVpcGroup PUT path
// Creates or updates the complete object group
func (*AmazonVpcGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/group/%s", ref)
}

// AmazonVpcTunnel is an Sophos Endpoint subType and implements sophos.RestObject
type AmazonVpcTunnels []AmazonVpcTunnel
type AmazonVpcTunnel struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Address   string `json:"address"`
	Bgp       string `json:"bgp"`
	Comment   string `json:"comment"`
	Ipsec     string `json:"ipsec"`
	Name      string `json:"name"`
	Netmask   int64  `json:"netmask"`
}

// GetPath implements sophos.RestObject and returns the AmazonVpcTunnels GET path
// Returns all available amazon_vpc/tunnel objects
func (*AmazonVpcTunnels) GetPath() string { return "/api/objects/amazon_vpc/tunnel/" }

// RefRequired implements sophos.RestObject
func (*AmazonVpcTunnels) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AmazonVpcTunnels GET path
// Returns all available tunnel types
func (a *AmazonVpcTunnel) GetPath() string {
	return fmt.Sprintf("/api/objects/amazon_vpc/tunnel/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AmazonVpcTunnel) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the AmazonVpcTunnel DELETE path
// Creates or updates the complete object tunnel
func (*AmazonVpcTunnel) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/tunnel/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AmazonVpcTunnel PATCH path
// Changes to parts of the object tunnel types
func (*AmazonVpcTunnel) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/tunnel/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AmazonVpcTunnel POST path
// Create a new amazon_vpc/tunnel object
func (*AmazonVpcTunnel) PostPath() string {
	return "/api/objects/amazon_vpc/tunnel/"
}

// PutPath implements sophos.RestObject and returns the AmazonVpcTunnel PUT path
// Creates or updates the complete object tunnel
func (*AmazonVpcTunnel) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/tunnel/%s", ref)
}

// Type implements sophos.Object
func (a *AmazonVpcTunnel) GetType() string { return a._type }

// AmazonVpcConnection is an Sophos Endpoint subType and implements sophos.RestObject
type AmazonVpcConnections []AmazonVpcConnection
type AmazonVpcConnection struct {
	Locked     string   `json:"_locked"`
	Reference  string   `json:"_ref"`
	_type      string   `json:"_type"`
	Comment    string   `json:"comment"`
	Dev        string   `json:"dev"`
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Region     string   `json:"region"`
	Status     bool     `json:"status"`
	Tunnel     []string `json:"tunnel"`
	VpcGateway string   `json:"vpc_gateway"`
	VpcID      string   `json:"vpc_id"`
	VpcNetmask int64    `json:"vpc_netmask"`
	VpcNetwork string   `json:"vpc_network"`
}

// GetPath implements sophos.RestObject and returns the AmazonVpcConnections GET path
// Returns all available amazon_vpc/connection objects
func (*AmazonVpcConnections) GetPath() string { return "/api/objects/amazon_vpc/connection/" }

// RefRequired implements sophos.RestObject
func (*AmazonVpcConnections) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AmazonVpcConnections GET path
// Returns all available connection types
func (a *AmazonVpcConnection) GetPath() string {
	return fmt.Sprintf("/api/objects/amazon_vpc/connection/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AmazonVpcConnection) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the AmazonVpcConnection DELETE path
// Creates or updates the complete object connection
func (*AmazonVpcConnection) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/connection/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AmazonVpcConnection PATCH path
// Changes to parts of the object connection types
func (*AmazonVpcConnection) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/connection/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AmazonVpcConnection POST path
// Create a new amazon_vpc/connection object
func (*AmazonVpcConnection) PostPath() string {
	return "/api/objects/amazon_vpc/connection/"
}

// PutPath implements sophos.RestObject and returns the AmazonVpcConnection PUT path
// Creates or updates the complete object connection
func (*AmazonVpcConnection) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/connection/%s", ref)
}

// Type implements sophos.Object
func (a *AmazonVpcConnection) GetType() string { return a._type }