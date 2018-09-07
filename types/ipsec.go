// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Ipsec is a generated struct representing the Sophos Ipsec Endpoint
// GET /api/nodes/ipsec
type Ipsec struct {
	Advanced struct {
		CrlAutoFetching       int64         `json:"crl_auto_fetching"`
		CrlStrictPolicy       int64         `json:"crl_strict_policy"`
		DeadPeerDetection     int64         `json:"dead_peer_detection"`
		IkeDebug              []interface{} `json:"ike_debug"`
		IkePort               int64         `json:"ike_port"`
		Metric                int64         `json:"metric"`
		NatTraversal          int64         `json:"nat_traversal"`
		NatTraversalKeepalive int64         `json:"nat_traversal_keepalive"`
		ProbePsk              int64         `json:"probe_psk"`
		PskVpnID              string        `json:"psk_vpn_id"`
		PskVpnIDType          string        `json:"psk_vpn_id_type"`
	} `json:"advanced"`
	Connections []string `json:"connections"`
	LocalRsa    string   `json:"local_rsa"`
	LocalX509   string   `json:"local_x509"`
	Status      int64    `json:"status"`
}

var defsIpsec = map[string]sophos.RestObject{
	"IpsecPolicy":        &IpsecPolicy{},
	"IpsecGroup":         &IpsecGroup{},
	"IpsecRemoteGateway": &IpsecRemoteGateway{},
}

// RestObjects implements the sophos.Node interface and returns a map of Ipsec's Objects
func (Ipsec) RestObjects() map[string]sophos.RestObject {
	return defsIpsec
}

// GetPath implements sophos.RestGetter
func (*Ipsec) GetPath() string { return "/api/nodes/ipsec" }

// RefRequired implements sophos.RestGetter
func (*Ipsec) RefRequired() (string, bool) { return "", false }

var defIpsec = &sophos.Definition{Description: "ipsec", Name: "ipsec", Link: "/api/definitions/ipsec", Swag: map[string]sophos.MethodMap{"/objects/ipsec/group/": {"get": sophos.MethodDescriptions{Description: "Returns all available ipsec/group objects", Parameters: []sophos.Parameter(nil), Tags: []string{"ipsec/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "post": sophos.MethodDescriptions{Description: "Create a new ipsec/group object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ipsec/group that will be created", Type: "", Required: true}}, Tags: []string{"ipsec/group"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}}, "/objects/ipsec/group/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"ipsec/group"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 204: {Description: "OK"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ipsec/group"}, Responses: map[int]struct{ Description string }{404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ipsec/group that will be changes", Type: "", Required: true}}, Tags: []string{"ipsec/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ipsec/group that will be updated", Type: "", Required: true}}, Tags: []string{"ipsec/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}, "/objects/ipsec/group/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ipsec/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/ipsec/policy/": {"post": sophos.MethodDescriptions{Description: "Create a new ipsec/policy object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ipsec/policy that will be created", Type: "", Required: true}}, Tags: []string{"ipsec/policy"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available ipsec/policy objects", Parameters: []sophos.Parameter(nil), Tags: []string{"ipsec/policy"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/ipsec/policy/{ref}": {"put": sophos.MethodDescriptions{Description: "Creates or updates the complete object policy", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ipsec/policy that will be updated", Type: "", Required: true}}, Tags: []string{"ipsec/policy"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object policy", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"ipsec/policy"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available policy types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ipsec/policy"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object policy types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ipsec/policy that will be changes", Type: "", Required: true}}, Tags: []string{"ipsec/policy"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}}, "/objects/ipsec/policy/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ipsec/policy"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/ipsec/remote_gateway/": {"get": sophos.MethodDescriptions{Description: "Returns all available ipsec/remote_gateway objects", Parameters: []sophos.Parameter(nil), Tags: []string{"ipsec/remote_gateway"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new ipsec/remote_gateway object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ipsec/remote_gateway that will be created", Type: "", Required: true}}, Tags: []string{"ipsec/remote_gateway"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}}, "/objects/ipsec/remote_gateway/{ref}": {"put": sophos.MethodDescriptions{Description: "Creates or updates the complete object remote_gateway", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ipsec/remote_gateway that will be updated", Type: "", Required: true}}, Tags: []string{"ipsec/remote_gateway"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object remote_gateway", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"ipsec/remote_gateway"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available remote_gateway types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ipsec/remote_gateway"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object remote_gateway types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ipsec/remote_gateway that will be changes", Type: "", Required: true}}, Tags: []string{"ipsec/remote_gateway"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}}, "/objects/ipsec/remote_gateway/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ipsec/remote_gateway"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 200: {Description: "OK"}, 401: {Description: "Unauthorized"}}}}}}

// Definition returns the /api/definitions struct of Ipsec
func (Ipsec) Definition() sophos.Definition { return *defIpsec }

// ApiRoutes returns all known Ipsec Paths
func (Ipsec) ApiRoutes() []string {
	return []string{
		"/api/objects/ipsec/group/",
		"/api/objects/ipsec/group/{ref}",
		"/api/objects/ipsec/group/{ref}/usedby",
		"/api/objects/ipsec/policy/",
		"/api/objects/ipsec/policy/{ref}",
		"/api/objects/ipsec/policy/{ref}/usedby",
		"/api/objects/ipsec/remote_gateway/",
		"/api/objects/ipsec/remote_gateway/{ref}",
		"/api/objects/ipsec/remote_gateway/{ref}/usedby",
	}
}

// References returns the Ipsec's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Ipsec) References() []string {
	return []string{
		"REF_IpsecGroup",
		"REF_IpsecPolicy",
		"REF_IpsecRemoteGateway",
	}
}

// IpsecPolicy is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecPolicys []IpsecPolicy
type IpsecPolicy struct {
	Locked            string `json:"_locked"`
	Reference         string `json:"_ref"`
	_type             string `json:"_type"`
	Comment           string `json:"comment"`
	IkeAuthAlg        string `json:"ike_auth_alg"`
	IkeDhGroup        string `json:"ike_dh_group"`
	IkeEncAlg         string `json:"ike_enc_alg"`
	IkeSaLifetime     int64  `json:"ike_sa_lifetime"`
	IpsecAuthAlg      string `json:"ipsec_auth_alg"`
	IpsecCompression  bool   `json:"ipsec_compression"`
	IpsecEncAlg       string `json:"ipsec_enc_alg"`
	IpsecPfsGroup     string `json:"ipsec_pfs_group"`
	IpsecSaLifetime   int64  `json:"ipsec_sa_lifetime"`
	IpsecStrictPolicy bool   `json:"ipsec_strict_policy"`
	Name              string `json:"name"`
}

// GetPath implements sophos.RestObject and returns the IpsecPolicys GET path
// Returns all available ipsec/policy objects
func (*IpsecPolicys) GetPath() string { return "/api/objects/ipsec/policy/" }

// RefRequired implements sophos.RestObject
func (*IpsecPolicys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecPolicys GET path
// Returns all available policy types
func (i *IpsecPolicy) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec/policy/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecPolicy) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecPolicy DELETE path
// Creates or updates the complete object policy
func (*IpsecPolicy) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/policy/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecPolicy PATCH path
// Changes to parts of the object policy types
func (*IpsecPolicy) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/policy/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecPolicy POST path
// Create a new ipsec/policy object
func (*IpsecPolicy) PostPath() string {
	return "/api/objects/ipsec/policy/"
}

// PutPath implements sophos.RestObject and returns the IpsecPolicy PUT path
// Creates or updates the complete object policy
func (*IpsecPolicy) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/policy/%s", ref)
}

// Type implements sophos.Object
func (i *IpsecPolicy) GetType() string { return i._type }

// IpsecGroup is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecGroup []interface{}

// GetPath implements sophos.RestObject and returns the IpsecGroup GET path
// Returns all available ipsec/group objects
func (*IpsecGroup) GetPath() string { return "/api/objects/ipsec/group/" }

// RefRequired implements sophos.RestObject
func (*IpsecGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the IpsecGroup DELETE path
// Creates or updates the complete object group
func (*IpsecGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecGroup PATCH path
// Changes to parts of the object group types
func (*IpsecGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecGroup POST path
// Create a new ipsec/group object
func (*IpsecGroup) PostPath() string {
	return "/api/objects/ipsec/group/"
}

// PutPath implements sophos.RestObject and returns the IpsecGroup PUT path
// Creates or updates the complete object group
func (*IpsecGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/group/%s", ref)
}

// IpsecRemoteGateway is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecRemoteGateways []IpsecRemoteGateway
type IpsecRemoteGateway struct {
	Locked         string   `json:"_locked"`
	Reference      string   `json:"_ref"`
	_type          string   `json:"_type"`
	Authentication string   `json:"authentication"`
	Comment        string   `json:"comment"`
	Ecn            bool     `json:"ecn"`
	Host           string   `json:"host"`
	Name           string   `json:"name"`
	Networks       []string `json:"networks"`
	PmtuDiscovery  bool     `json:"pmtu_discovery"`
	Xauth          bool     `json:"xauth"`
	XauthPassword  string   `json:"xauth_password"`
	XauthUsername  string   `json:"xauth_username"`
}

// GetPath implements sophos.RestObject and returns the IpsecRemoteGateways GET path
// Returns all available ipsec/remote_gateway objects
func (*IpsecRemoteGateways) GetPath() string { return "/api/objects/ipsec/remote_gateway/" }

// RefRequired implements sophos.RestObject
func (*IpsecRemoteGateways) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecRemoteGateways GET path
// Returns all available remote_gateway types
func (i *IpsecRemoteGateway) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec/remote_gateway/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecRemoteGateway) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecRemoteGateway DELETE path
// Creates or updates the complete object remote_gateway
func (*IpsecRemoteGateway) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/remote_gateway/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecRemoteGateway PATCH path
// Changes to parts of the object remote_gateway types
func (*IpsecRemoteGateway) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/remote_gateway/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecRemoteGateway POST path
// Create a new ipsec/remote_gateway object
func (*IpsecRemoteGateway) PostPath() string {
	return "/api/objects/ipsec/remote_gateway/"
}

// PutPath implements sophos.RestObject and returns the IpsecRemoteGateway PUT path
// Creates or updates the complete object remote_gateway
func (*IpsecRemoteGateway) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/remote_gateway/%s", ref)
}

// Type implements sophos.Object
func (i *IpsecRemoteGateway) GetType() string { return i._type }