// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// RemoteSyslog is a generated struct representing the Sophos RemoteSyslog Endpoint
// GET /api/nodes/remote_syslog
type RemoteSyslog struct {
	Buffer int64         `json:"buffer"`
	Logs   []interface{} `json:"logs"`
	Status int64         `json:"status"`
	Target []interface{} `json:"target"`
}

var defsRemoteSyslog = map[string]sophos.RestObject{
	"RemoteSyslogGroup":  &RemoteSyslogGroup{},
	"RemoteSyslogServer": &RemoteSyslogServer{},
}

// RestObjects implements the sophos.Node interface and returns a map of RemoteSyslog's Objects
func (RemoteSyslog) RestObjects() map[string]sophos.RestObject {
	return defsRemoteSyslog
}

// GetPath implements sophos.RestGetter
func (*RemoteSyslog) GetPath() string { return "/api/nodes/remote_syslog" }

// RefRequired implements sophos.RestGetter
func (*RemoteSyslog) RefRequired() (string, bool) { return "", false }

var defRemoteSyslog = &sophos.Definition{Description: "remote_syslog", Name: "remote_syslog", Link: "/api/definitions/remote_syslog", Swag: map[string]sophos.MethodMap{"/objects/remote_syslog/server/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"remote_syslog/server"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/remote_syslog/group/": {"get": sophos.MethodDescriptions{Description: "Returns all available remote_syslog/group objects", Parameters: []sophos.Parameter(nil), Tags: []string{"remote_syslog/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "post": sophos.MethodDescriptions{Description: "Create a new remote_syslog/group object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "remote_syslog/group that will be created", Type: "", Required: true}}, Tags: []string{"remote_syslog/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}}}}, "/objects/remote_syslog/group/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"remote_syslog/group"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"remote_syslog/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "remote_syslog/group that will be changes", Type: "", Required: true}}, Tags: []string{"remote_syslog/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "remote_syslog/group that will be updated", Type: "", Required: true}}, Tags: []string{"remote_syslog/group"}, Responses: map[int]struct{ Description string }{404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/remote_syslog/group/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"remote_syslog/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/remote_syslog/server/": {"get": sophos.MethodDescriptions{Description: "Returns all available remote_syslog/server objects", Parameters: []sophos.Parameter(nil), Tags: []string{"remote_syslog/server"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new remote_syslog/server object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "remote_syslog/server that will be created", Type: "", Required: true}}, Tags: []string{"remote_syslog/server"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}}}}, "/objects/remote_syslog/server/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object server", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"remote_syslog/server"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available server types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"remote_syslog/server"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object server types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "remote_syslog/server that will be changes", Type: "", Required: true}}, Tags: []string{"remote_syslog/server"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object server", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "remote_syslog/server that will be updated", Type: "", Required: true}}, Tags: []string{"remote_syslog/server"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}}}

// Definition returns the /api/definitions struct of RemoteSyslog
func (RemoteSyslog) Definition() sophos.Definition { return *defRemoteSyslog }

// ApiRoutes returns all known RemoteSyslog Paths
func (RemoteSyslog) ApiRoutes() []string {
	return []string{
		"/api/objects/remote_syslog/group/",
		"/api/objects/remote_syslog/group/{ref}",
		"/api/objects/remote_syslog/group/{ref}/usedby",
		"/api/objects/remote_syslog/server/",
		"/api/objects/remote_syslog/server/{ref}",
		"/api/objects/remote_syslog/server/{ref}/usedby",
	}
}

// References returns the RemoteSyslog's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (RemoteSyslog) References() []string {
	return []string{
		"REF_RemoteSyslogGroup",
		"REF_RemoteSyslogServer",
	}
}

// RemoteSyslogGroup is an Sophos Endpoint subType and implements sophos.RestObject
type RemoteSyslogGroup []interface{}

// GetPath implements sophos.RestObject and returns the RemoteSyslogGroup GET path
// Returns all available remote_syslog/group objects
func (*RemoteSyslogGroup) GetPath() string { return "/api/objects/remote_syslog/group/" }

// RefRequired implements sophos.RestObject
func (*RemoteSyslogGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the RemoteSyslogGroup DELETE path
// Creates or updates the complete object group
func (*RemoteSyslogGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the RemoteSyslogGroup PATCH path
// Changes to parts of the object group types
func (*RemoteSyslogGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the RemoteSyslogGroup POST path
// Create a new remote_syslog/group object
func (*RemoteSyslogGroup) PostPath() string {
	return "/api/objects/remote_syslog/group/"
}

// PutPath implements sophos.RestObject and returns the RemoteSyslogGroup PUT path
// Creates or updates the complete object group
func (*RemoteSyslogGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/group/%s", ref)
}

// RemoteSyslogServer is an Sophos Endpoint subType and implements sophos.RestObject
type RemoteSyslogServer []interface{}

// GetPath implements sophos.RestObject and returns the RemoteSyslogServer GET path
// Returns all available remote_syslog/server objects
func (*RemoteSyslogServer) GetPath() string { return "/api/objects/remote_syslog/server/" }

// RefRequired implements sophos.RestObject
func (*RemoteSyslogServer) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the RemoteSyslogServer DELETE path
// Creates or updates the complete object server
func (*RemoteSyslogServer) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/server/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the RemoteSyslogServer PATCH path
// Changes to parts of the object server types
func (*RemoteSyslogServer) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/server/%s", ref)
}

// PostPath implements sophos.RestObject and returns the RemoteSyslogServer POST path
// Create a new remote_syslog/server object
func (*RemoteSyslogServer) PostPath() string {
	return "/api/objects/remote_syslog/server/"
}

// PutPath implements sophos.RestObject and returns the RemoteSyslogServer PUT path
// Creates or updates the complete object server
func (*RemoteSyslogServer) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/server/%s", ref)
}