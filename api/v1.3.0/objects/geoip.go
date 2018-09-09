package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Geoip is a generated struct representing the Sophos Geoip Endpoint
// GET /api/nodes/geoip
type Geoip struct {
	CountriesDst []string      `json:"countries_dst"`
	CountriesSrc []string      `json:"countries_src"`
	Exceptions   []interface{} `json:"exceptions"`
	Log          string        `json:"log"`
	Status       int64         `json:"status"`
}

var defsGeoip = map[string]sophos.RestObject{
	"GeoipDstexception": &GeoipDstexception{},
	"GeoipGeoipgroup":   &GeoipGeoipgroup{},
	"GeoipGroup":        &GeoipGroup{},
	"GeoipSrcexception": &GeoipSrcexception{},
}

// RestObjects implements the sophos.Node interface and returns a map of Geoip's Objects
func (Geoip) RestObjects() map[string]sophos.RestObject { return defsGeoip }

// GetPath implements sophos.RestGetter
func (*Geoip) GetPath() string { return "/api/nodes/geoip" }

// RefRequired implements sophos.RestGetter
func (*Geoip) RefRequired() (string, bool) { return "", false }

var defGeoip = &sophos.Definition{Description: "geoip", Name: "geoip", Link: "/api/definitions/geoip"}

// Definition returns the /api/definitions struct of Geoip
func (Geoip) Definition() sophos.Definition { return *defGeoip }

// ApiRoutes returns all known Geoip Paths
func (Geoip) ApiRoutes() []string {
	return []string{
		"/api/objects/geoip/dstexception/",
		"/api/objects/geoip/dstexception/{ref}",
		"/api/objects/geoip/dstexception/{ref}/usedby",
		"/api/objects/geoip/geoipgroup/",
		"/api/objects/geoip/geoipgroup/{ref}",
		"/api/objects/geoip/geoipgroup/{ref}/usedby",
		"/api/objects/geoip/group/",
		"/api/objects/geoip/group/{ref}",
		"/api/objects/geoip/group/{ref}/usedby",
		"/api/objects/geoip/srcexception/",
		"/api/objects/geoip/srcexception/{ref}",
		"/api/objects/geoip/srcexception/{ref}/usedby",
	}
}

// References returns the Geoip's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Geoip) References() []string {
	return []string{
		"REF_GeoipDstexception",
		"REF_GeoipGeoipgroup",
		"REF_GeoipGroup",
		"REF_GeoipSrcexception",
	}
}

// GeoipDstexception is an Sophos Endpoint subType and implements sophos.RestObject
type GeoipDstexception []interface{}

// GetPath implements sophos.RestObject and returns the GeoipDstexception GET path
// Returns all available geoip/dstexception objects
func (*GeoipDstexception) GetPath() string { return "/api/objects/geoip/dstexception/" }

// RefRequired implements sophos.RestObject
func (*GeoipDstexception) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the GeoipDstexception DELETE path
// Creates or updates the complete object dstexception
func (*GeoipDstexception) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/dstexception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the GeoipDstexception PATCH path
// Changes to parts of the object dstexception types
func (*GeoipDstexception) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/dstexception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the GeoipDstexception POST path
// Create a new geoip/dstexception object
func (*GeoipDstexception) PostPath() string {
	return "/api/objects/geoip/dstexception/"
}

// PutPath implements sophos.RestObject and returns the GeoipDstexception PUT path
// Creates or updates the complete object dstexception
func (*GeoipDstexception) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/dstexception/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*GeoipDstexception) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/dstexception/%s/usedby", ref)
}

// GeoipGeoipgroups is an Sophos Endpoint subType and implements sophos.RestObject
type GeoipGeoipgroups []GeoipGeoipgroup

// GeoipGeoipgroup is a generated Sophos object
type GeoipGeoipgroup struct {
	Locked    string   `json:"_locked"`
	Reference string   `json:"_ref"`
	_type     string   `json:"_type"`
	Comment   string   `json:"comment"`
	Countries []string `json:"countries"`
	Name      string   `json:"name"`
}

// GetPath implements sophos.RestObject and returns the GeoipGeoipgroups GET path
// Returns all available geoip/geoipgroup objects
func (*GeoipGeoipgroups) GetPath() string { return "/api/objects/geoip/geoipgroup/" }

// RefRequired implements sophos.RestObject
func (*GeoipGeoipgroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the GeoipGeoipgroups GET path
// Returns all available geoipgroup types
func (g *GeoipGeoipgroup) GetPath() string {
	return fmt.Sprintf("/api/objects/geoip/geoipgroup/%s", g.Reference)
}

// RefRequired implements sophos.RestObject
func (g *GeoipGeoipgroup) RefRequired() (string, bool) { return g.Reference, true }

// DeletePath implements sophos.RestObject and returns the GeoipGeoipgroup DELETE path
// Creates or updates the complete object geoipgroup
func (*GeoipGeoipgroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/geoipgroup/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the GeoipGeoipgroup PATCH path
// Changes to parts of the object geoipgroup types
func (*GeoipGeoipgroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/geoipgroup/%s", ref)
}

// PostPath implements sophos.RestObject and returns the GeoipGeoipgroup POST path
// Create a new geoip/geoipgroup object
func (*GeoipGeoipgroup) PostPath() string {
	return "/api/objects/geoip/geoipgroup/"
}

// PutPath implements sophos.RestObject and returns the GeoipGeoipgroup PUT path
// Creates or updates the complete object geoipgroup
func (*GeoipGeoipgroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/geoipgroup/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*GeoipGeoipgroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/geoipgroup/%s/usedby", ref)
}

// GetType implements sophos.Object
func (g *GeoipGeoipgroup) GetType() string { return g._type }

// GeoipGroup is an Sophos Endpoint subType and implements sophos.RestObject
type GeoipGroup []interface{}

// GetPath implements sophos.RestObject and returns the GeoipGroup GET path
// Returns all available geoip/group objects
func (*GeoipGroup) GetPath() string { return "/api/objects/geoip/group/" }

// RefRequired implements sophos.RestObject
func (*GeoipGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the GeoipGroup DELETE path
// Creates or updates the complete object group
func (*GeoipGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the GeoipGroup PATCH path
// Changes to parts of the object group types
func (*GeoipGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the GeoipGroup POST path
// Create a new geoip/group object
func (*GeoipGroup) PostPath() string {
	return "/api/objects/geoip/group/"
}

// PutPath implements sophos.RestObject and returns the GeoipGroup PUT path
// Creates or updates the complete object group
func (*GeoipGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*GeoipGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/group/%s/usedby", ref)
}

// GeoipSrcexception is an Sophos Endpoint subType and implements sophos.RestObject
type GeoipSrcexception []interface{}

// GetPath implements sophos.RestObject and returns the GeoipSrcexception GET path
// Returns all available geoip/srcexception objects
func (*GeoipSrcexception) GetPath() string { return "/api/objects/geoip/srcexception/" }

// RefRequired implements sophos.RestObject
func (*GeoipSrcexception) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the GeoipSrcexception DELETE path
// Creates or updates the complete object srcexception
func (*GeoipSrcexception) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/srcexception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the GeoipSrcexception PATCH path
// Changes to parts of the object srcexception types
func (*GeoipSrcexception) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/srcexception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the GeoipSrcexception POST path
// Create a new geoip/srcexception object
func (*GeoipSrcexception) PostPath() string {
	return "/api/objects/geoip/srcexception/"
}

// PutPath implements sophos.RestObject and returns the GeoipSrcexception PUT path
// Creates or updates the complete object srcexception
func (*GeoipSrcexception) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/srcexception/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*GeoipSrcexception) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/srcexception/%s/usedby", ref)
}