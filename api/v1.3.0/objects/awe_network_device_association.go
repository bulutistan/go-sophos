package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// AweNetworkDeviceAssociation is a generated struct representing the Sophos AweNetworkDeviceAssociation Endpoint
// GET /api/nodes/awe_network_device_association
type AweNetworkDeviceAssociation struct {
	AweNetworkDeviceAssociationGroup    AweNetworkDeviceAssociationGroup    `json:"awe_network_device_association_group"`
	AweNetworkDeviceAssociationMeshRole AweNetworkDeviceAssociationMeshRole `json:"awe_network_device_association_mesh_role"`
}

var defsAweNetworkDeviceAssociation = map[string]sophos.RestObject{
	"AweNetworkDeviceAssociationGroup":    &AweNetworkDeviceAssociationGroup{},
	"AweNetworkDeviceAssociationMeshRole": &AweNetworkDeviceAssociationMeshRole{},
}

// RestObjects implements the sophos.Node interface and returns a map of AweNetworkDeviceAssociation's Objects
func (AweNetworkDeviceAssociation) RestObjects() map[string]sophos.RestObject {
	return defsAweNetworkDeviceAssociation
}

// GetPath implements sophos.RestGetter
func (*AweNetworkDeviceAssociation) GetPath() string {
	return "/api/nodes/awe_network_device_association"
}

// RefRequired implements sophos.RestGetter
func (*AweNetworkDeviceAssociation) RefRequired() (string, bool) { return "", false }

var defAweNetworkDeviceAssociation = &sophos.Definition{Description: "awe_network_device_association", Name: "awe_network_device_association", Link: "/api/definitions/awe_network_device_association"}

// Definition returns the /api/definitions struct of AweNetworkDeviceAssociation
func (AweNetworkDeviceAssociation) Definition() sophos.Definition {
	return *defAweNetworkDeviceAssociation
}

// ApiRoutes returns all known AweNetworkDeviceAssociation Paths
func (AweNetworkDeviceAssociation) ApiRoutes() []string {
	return []string{
		"/api/objects/awe_network_device_association/group/",
		"/api/objects/awe_network_device_association/group/{ref}",
		"/api/objects/awe_network_device_association/group/{ref}/usedby",
		"/api/objects/awe_network_device_association/mesh_role/",
		"/api/objects/awe_network_device_association/mesh_role/{ref}",
		"/api/objects/awe_network_device_association/mesh_role/{ref}/usedby",
	}
}

// References returns the AweNetworkDeviceAssociation's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (AweNetworkDeviceAssociation) References() []string {
	return []string{
		"REF_AweNetworkDeviceAssociationGroup",
		"REF_AweNetworkDeviceAssociationMeshRole",
	}
}

// AweNetworkDeviceAssociationGroup is an Sophos Endpoint subType and implements sophos.RestObject
type AweNetworkDeviceAssociationGroup []interface{}

// GetPath implements sophos.RestObject and returns the AweNetworkDeviceAssociationGroup GET path
// Returns all available awe_network_device_association/group objects
func (*AweNetworkDeviceAssociationGroup) GetPath() string {
	return "/api/objects/awe_network_device_association/group/"
}

// RefRequired implements sophos.RestObject
func (*AweNetworkDeviceAssociationGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AweNetworkDeviceAssociationGroup DELETE path
// Creates or updates the complete object group
func (*AweNetworkDeviceAssociationGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/awe_network_device_association/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AweNetworkDeviceAssociationGroup PATCH path
// Changes to parts of the object group types
func (*AweNetworkDeviceAssociationGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe_network_device_association/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AweNetworkDeviceAssociationGroup POST path
// Create a new awe_network_device_association/group object
func (*AweNetworkDeviceAssociationGroup) PostPath() string {
	return "/api/objects/awe_network_device_association/group/"
}

// PutPath implements sophos.RestObject and returns the AweNetworkDeviceAssociationGroup PUT path
// Creates or updates the complete object group
func (*AweNetworkDeviceAssociationGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe_network_device_association/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*AweNetworkDeviceAssociationGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe_network_device_association/group/%s/usedby", ref)
}

// AweNetworkDeviceAssociationMeshRole is an Sophos Endpoint subType and implements sophos.RestObject
type AweNetworkDeviceAssociationMeshRole []interface{}

// GetPath implements sophos.RestObject and returns the AweNetworkDeviceAssociationMeshRole GET path
// Returns all available awe_network_device_association/mesh_role objects
func (*AweNetworkDeviceAssociationMeshRole) GetPath() string {
	return "/api/objects/awe_network_device_association/mesh_role/"
}

// RefRequired implements sophos.RestObject
func (*AweNetworkDeviceAssociationMeshRole) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AweNetworkDeviceAssociationMeshRole DELETE path
// Creates or updates the complete object mesh_role
func (*AweNetworkDeviceAssociationMeshRole) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/awe_network_device_association/mesh_role/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AweNetworkDeviceAssociationMeshRole PATCH path
// Changes to parts of the object mesh_role types
func (*AweNetworkDeviceAssociationMeshRole) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe_network_device_association/mesh_role/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AweNetworkDeviceAssociationMeshRole POST path
// Create a new awe_network_device_association/mesh_role object
func (*AweNetworkDeviceAssociationMeshRole) PostPath() string {
	return "/api/objects/awe_network_device_association/mesh_role/"
}

// PutPath implements sophos.RestObject and returns the AweNetworkDeviceAssociationMeshRole PUT path
// Creates or updates the complete object mesh_role
func (*AweNetworkDeviceAssociationMeshRole) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe_network_device_association/mesh_role/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*AweNetworkDeviceAssociationMeshRole) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe_network_device_association/mesh_role/%s/usedby", ref)
}