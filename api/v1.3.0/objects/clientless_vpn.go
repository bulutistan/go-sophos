package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// ClientlessVpn is a generated struct representing the Sophos ClientlessVpn Endpoint
// GET /api/nodes/clientless_vpn
type ClientlessVpn struct {
	ClientlessVpnConnection ClientlessVpnConnection `json:"clientless_vpn_connection"`
	ClientlessVpnGroup      ClientlessVpnGroup      `json:"clientless_vpn_group"`
}

var defsClientlessVpn = map[string]sophos.RestObject{
	"ClientlessVpnConnection": &ClientlessVpnConnection{},
	"ClientlessVpnGroup":      &ClientlessVpnGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of ClientlessVpn's Objects
func (ClientlessVpn) RestObjects() map[string]sophos.RestObject { return defsClientlessVpn }

// GetPath implements sophos.RestGetter
func (*ClientlessVpn) GetPath() string { return "/api/nodes/clientless_vpn" }

// RefRequired implements sophos.RestGetter
func (*ClientlessVpn) RefRequired() (string, bool) { return "", false }

var defClientlessVpn = &sophos.Definition{Description: "clientless_vpn", Name: "clientless_vpn", Link: "/api/definitions/clientless_vpn"}

// Definition returns the /api/definitions struct of ClientlessVpn
func (ClientlessVpn) Definition() sophos.Definition { return *defClientlessVpn }

// ApiRoutes returns all known ClientlessVpn Paths
func (ClientlessVpn) ApiRoutes() []string {
	return []string{
		"/api/objects/clientless_vpn/connection/",
		"/api/objects/clientless_vpn/connection/{ref}",
		"/api/objects/clientless_vpn/connection/{ref}/usedby",
		"/api/objects/clientless_vpn/group/",
		"/api/objects/clientless_vpn/group/{ref}",
		"/api/objects/clientless_vpn/group/{ref}/usedby",
	}
}

// References returns the ClientlessVpn's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (ClientlessVpn) References() []string {
	return []string{
		"REF_ClientlessVpnConnection",
		"REF_ClientlessVpnGroup",
	}
}

// ClientlessVpnConnections is an Sophos Endpoint subType and implements sophos.RestObject
type ClientlessVpnConnections []ClientlessVpnConnection

// ClientlessVpnConnection is a generated Sophos object
type ClientlessVpnConnection struct {
	Locked        string        `json:"_locked"`
	Reference     string        `json:"_ref"`
	_type         string        `json:"_type"`
	AllowedUsers  []string      `json:"allowed_users"`
	AutoLogin     bool          `json:"auto_login"`
	Comment       string        `json:"comment"`
	Destination   string        `json:"destination"`
	HostKeyCert   string        `json:"host_key_cert"`
	Login         string        `json:"login"`
	Name          string        `json:"name"`
	Password      string        `json:"password"`
	PfExceptions  []interface{} `json:"pf_exceptions"`
	Port          int64         `json:"port"`
	PrivateKey    string        `json:"private_key"`
	RdpSecurity   string        `json:"rdp_security"`
	RecordSession bool          `json:"record_session"`
	Service       string        `json:"service"`
	ShareSession  bool          `json:"share_session"`
	Status        bool          `json:"status"`
	UID           int64         `json:"uid"`
	WebPath       string        `json:"web_path"`
}

// GetPath implements sophos.RestObject and returns the ClientlessVpnConnections GET path
// Returns all available clientless_vpn/connection objects
func (*ClientlessVpnConnections) GetPath() string { return "/api/objects/clientless_vpn/connection/" }

// RefRequired implements sophos.RestObject
func (*ClientlessVpnConnections) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ClientlessVpnConnections GET path
// Returns all available connection types
func (c *ClientlessVpnConnection) GetPath() string {
	return fmt.Sprintf("/api/objects/clientless_vpn/connection/%s", c.Reference)
}

// RefRequired implements sophos.RestObject
func (c *ClientlessVpnConnection) RefRequired() (string, bool) { return c.Reference, true }

// DeletePath implements sophos.RestObject and returns the ClientlessVpnConnection DELETE path
// Creates or updates the complete object connection
func (*ClientlessVpnConnection) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/connection/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ClientlessVpnConnection PATCH path
// Changes to parts of the object connection types
func (*ClientlessVpnConnection) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/connection/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ClientlessVpnConnection POST path
// Create a new clientless_vpn/connection object
func (*ClientlessVpnConnection) PostPath() string {
	return "/api/objects/clientless_vpn/connection/"
}

// PutPath implements sophos.RestObject and returns the ClientlessVpnConnection PUT path
// Creates or updates the complete object connection
func (*ClientlessVpnConnection) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/connection/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*ClientlessVpnConnection) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/connection/%s/usedby", ref)
}

// GetType implements sophos.Object
func (c *ClientlessVpnConnection) GetType() string { return c._type }

// ClientlessVpnGroup is an Sophos Endpoint subType and implements sophos.RestObject
type ClientlessVpnGroup []interface{}

// GetPath implements sophos.RestObject and returns the ClientlessVpnGroup GET path
// Returns all available clientless_vpn/group objects
func (*ClientlessVpnGroup) GetPath() string { return "/api/objects/clientless_vpn/group/" }

// RefRequired implements sophos.RestObject
func (*ClientlessVpnGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ClientlessVpnGroup DELETE path
// Creates or updates the complete object group
func (*ClientlessVpnGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ClientlessVpnGroup PATCH path
// Changes to parts of the object group types
func (*ClientlessVpnGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ClientlessVpnGroup POST path
// Create a new clientless_vpn/group object
func (*ClientlessVpnGroup) PostPath() string {
	return "/api/objects/clientless_vpn/group/"
}

// PutPath implements sophos.RestObject and returns the ClientlessVpnGroup PUT path
// Creates or updates the complete object group
func (*ClientlessVpnGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*ClientlessVpnGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/group/%s/usedby", ref)
}