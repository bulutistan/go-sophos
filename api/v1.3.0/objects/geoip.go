package objects

import (
	"fmt"

	"github.com/bulutistan/go-sophos"
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

var _ sophos.Endpoint = &Geoip{}

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

// GeoipDstexceptions is an Sophos Endpoint subType and implements sophos.RestObject
type GeoipDstexceptions []GeoipDstexception

// GeoipDstexception represents a UTM incoming exceptions
type GeoipDstexception struct {
	Locked              string        `json:"_locked"`
	ObjectType          string        `json:"_type"`
	Reference           string        `json:"_ref"`
	DestinationNetworks []interface{} `json:"destination_networks"`
	Name                string        `json:"name"`
	Services            []interface{} `json:"services"`
	// Status default value is false
	Status    bool          `json:"status"`
	Comment   string        `json:"comment"`
	Countries []interface{} `json:"countries"`
}

var _ sophos.RestGetter = &GeoipDstexception{}

// GetPath implements sophos.RestObject and returns the GeoipDstexceptions GET path
// Returns all available geoip/dstexception objects
func (*GeoipDstexceptions) GetPath() string { return "/api/objects/geoip/dstexception/" }

// RefRequired implements sophos.RestObject
func (*GeoipDstexceptions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the GeoipDstexceptions GET path
// Returns all available dstexception types
func (g *GeoipDstexception) GetPath() string {
	return fmt.Sprintf("/api/objects/geoip/dstexception/%s", g.Reference)
}

// RefRequired implements sophos.RestObject
func (g *GeoipDstexception) RefRequired() (string, bool) { return g.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*GeoipDstexception) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/dstexception/%s/usedby", ref)
}

// GeoipGeoipgroups is an Sophos Endpoint subType and implements sophos.RestObject
type GeoipGeoipgroups []GeoipGeoipgroup

// GeoipGeoipgroup is a generated Sophos object
type GeoipGeoipgroup struct {
	Locked     string   `json:"_locked"`
	Reference  string   `json:"_ref"`
	ObjectType string   `json:"_type"`
	Comment    string   `json:"comment"`
	Countries  []string `json:"countries"`
	Name       string   `json:"name"`
}

var _ sophos.RestGetter = &GeoipGeoipgroup{}

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*GeoipGeoipgroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/geoipgroup/%s/usedby", ref)
}

// GetType implements sophos.Object
func (g *GeoipGeoipgroup) GetType() string { return g.ObjectType }

// GeoipGroups is an Sophos Endpoint subType and implements sophos.RestObject
type GeoipGroups []GeoipGroup

// GeoipGroup represents a UTM group
type GeoipGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &GeoipGroup{}

// GetPath implements sophos.RestObject and returns the GeoipGroups GET path
// Returns all available geoip/group objects
func (*GeoipGroups) GetPath() string { return "/api/objects/geoip/group/" }

// RefRequired implements sophos.RestObject
func (*GeoipGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the GeoipGroups GET path
// Returns all available group types
func (g *GeoipGroup) GetPath() string { return fmt.Sprintf("/api/objects/geoip/group/%s", g.Reference) }

// RefRequired implements sophos.RestObject
func (g *GeoipGroup) RefRequired() (string, bool) { return g.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*GeoipGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/group/%s/usedby", ref)
}

// GeoipSrcexceptions is an Sophos Endpoint subType and implements sophos.RestObject
type GeoipSrcexceptions []GeoipSrcexception

// GeoipSrcexception represents a UTM outgoing exceptions
type GeoipSrcexception struct {
	Locked         string        `json:"_locked"`
	ObjectType     string        `json:"_type"`
	Reference      string        `json:"_ref"`
	Comment        string        `json:"comment"`
	Countries      []interface{} `json:"countries"`
	Name           string        `json:"name"`
	Services       []interface{} `json:"services"`
	SourceNetworks []interface{} `json:"source_networks"`
	// Status default value is false
	Status bool `json:"status"`
}

var _ sophos.RestGetter = &GeoipSrcexception{}

// GetPath implements sophos.RestObject and returns the GeoipSrcexceptions GET path
// Returns all available geoip/srcexception objects
func (*GeoipSrcexceptions) GetPath() string { return "/api/objects/geoip/srcexception/" }

// RefRequired implements sophos.RestObject
func (*GeoipSrcexceptions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the GeoipSrcexceptions GET path
// Returns all available srcexception types
func (g *GeoipSrcexception) GetPath() string {
	return fmt.Sprintf("/api/objects/geoip/srcexception/%s", g.Reference)
}

// RefRequired implements sophos.RestObject
func (g *GeoipSrcexception) RefRequired() (string, bool) { return g.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*GeoipSrcexception) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/geoip/srcexception/%s/usedby", ref)
}
