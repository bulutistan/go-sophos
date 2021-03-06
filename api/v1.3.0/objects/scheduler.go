package objects

import (
	"fmt"

	"github.com/bulutistan/go-sophos/sophos"
)

// Scheduler is a generated struct representing the Sophos Scheduler Endpoint
// GET /api/nodes/scheduler
type Scheduler struct {
	SchedulerGroup       SchedulerGroup       `json:"scheduler_group"`
	SchedulerLoadbalance SchedulerLoadbalance `json:"scheduler_loadbalance"`
	SchedulerRule        SchedulerRule        `json:"scheduler_rule"`
}

var _ sophos.Endpoint = &Scheduler{}

var defsScheduler = map[string]sophos.RestObject{
	"SchedulerGroup":       &SchedulerGroup{},
	"SchedulerLoadbalance": &SchedulerLoadbalance{},
	"SchedulerRule":        &SchedulerRule{},
}

// RestObjects implements the sophos.Node interface and returns a map of Scheduler's Objects
func (Scheduler) RestObjects() map[string]sophos.RestObject { return defsScheduler }

// GetPath implements sophos.RestGetter
func (*Scheduler) GetPath() string { return "/api/nodes/scheduler" }

// RefRequired implements sophos.RestGetter
func (*Scheduler) RefRequired() (string, bool) { return "", false }

var defScheduler = &sophos.Definition{Description: "scheduler", Name: "scheduler", Link: "/api/definitions/scheduler"}

// Definition returns the /api/definitions struct of Scheduler
func (Scheduler) Definition() sophos.Definition { return *defScheduler }

// ApiRoutes returns all known Scheduler Paths
func (Scheduler) ApiRoutes() []string {
	return []string{
		"/api/objects/scheduler/group/",
		"/api/objects/scheduler/group/{ref}",
		"/api/objects/scheduler/group/{ref}/usedby",
		"/api/objects/scheduler/loadbalance/",
		"/api/objects/scheduler/loadbalance/{ref}",
		"/api/objects/scheduler/loadbalance/{ref}/usedby",
		"/api/objects/scheduler/rule/",
		"/api/objects/scheduler/rule/{ref}",
		"/api/objects/scheduler/rule/{ref}/usedby",
	}
}

// References returns the Scheduler's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Scheduler) References() []string {
	return []string{
		"REF_SchedulerGroup",
		"REF_SchedulerLoadbalance",
		"REF_SchedulerRule",
	}
}

// SchedulerGroups is an Sophos Endpoint subType and implements sophos.RestObject
type SchedulerGroups []SchedulerGroup

// SchedulerGroup represents a UTM group
type SchedulerGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &SchedulerGroup{}

// GetPath implements sophos.RestObject and returns the SchedulerGroups GET path
// Returns all available scheduler/group objects
func (*SchedulerGroups) GetPath() string { return "/api/objects/scheduler/group/" }

// RefRequired implements sophos.RestObject
func (*SchedulerGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SchedulerGroups GET path
// Returns all available group types
func (s *SchedulerGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/scheduler/group/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *SchedulerGroup) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SchedulerGroup DELETE path
// Creates or updates the complete object group
func (*SchedulerGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SchedulerGroup PATCH path
// Changes to parts of the object group types
func (*SchedulerGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SchedulerGroup POST path
// Create a new scheduler/group object
func (*SchedulerGroup) PostPath() string {
	return "/api/objects/scheduler/group/"
}

// PutPath implements sophos.RestObject and returns the SchedulerGroup PUT path
// Creates or updates the complete object group
func (*SchedulerGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SchedulerGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/group/%s/usedby", ref)
}

// SchedulerLoadbalances is an Sophos Endpoint subType and implements sophos.RestObject
type SchedulerLoadbalances []SchedulerLoadbalance

// SchedulerLoadbalance is a generated Sophos object
type SchedulerLoadbalance struct {
	Locked          string        `json:"_locked"`
	Reference       string        `json:"_ref"`
	ObjectType      string        `json:"_type"`
	Algorithm       string        `json:"algorithm"`
	CheckData       string        `json:"check_data"`
	CheckHosts      []interface{} `json:"check_hosts"`
	CheckInterval   int64         `json:"check_interval"`
	CheckPort       int64         `json:"check_port"`
	CheckTimeout    int64         `json:"check_timeout"`
	CheckType       string        `json:"check_type"`
	Comment         string        `json:"comment"`
	Name            string        `json:"name"`
	Persistence     bool          `json:"persistence"`
	PersistenceHash string        `json:"persistence_hash"`
	PersistenceSize int64         `json:"persistence_size"`
	PersistenceTime int64         `json:"persistence_time"`
	Weight          struct{}      `json:"weight"`
}

var _ sophos.RestGetter = &SchedulerLoadbalance{}

// GetPath implements sophos.RestObject and returns the SchedulerLoadbalances GET path
// Returns all available scheduler/loadbalance objects
func (*SchedulerLoadbalances) GetPath() string { return "/api/objects/scheduler/loadbalance/" }

// RefRequired implements sophos.RestObject
func (*SchedulerLoadbalances) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SchedulerLoadbalances GET path
// Returns all available loadbalance types
func (s *SchedulerLoadbalance) GetPath() string {
	return fmt.Sprintf("/api/objects/scheduler/loadbalance/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *SchedulerLoadbalance) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SchedulerLoadbalance DELETE path
// Creates or updates the complete object loadbalance
func (*SchedulerLoadbalance) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/loadbalance/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SchedulerLoadbalance PATCH path
// Changes to parts of the object loadbalance types
func (*SchedulerLoadbalance) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/loadbalance/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SchedulerLoadbalance POST path
// Create a new scheduler/loadbalance object
func (*SchedulerLoadbalance) PostPath() string {
	return "/api/objects/scheduler/loadbalance/"
}

// PutPath implements sophos.RestObject and returns the SchedulerLoadbalance PUT path
// Creates or updates the complete object loadbalance
func (*SchedulerLoadbalance) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/loadbalance/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SchedulerLoadbalance) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/loadbalance/%s/usedby", ref)
}

// GetType implements sophos.Object
func (s *SchedulerLoadbalance) GetType() string { return s.ObjectType }

// SchedulerRules is an Sophos Endpoint subType and implements sophos.RestObject
type SchedulerRules []SchedulerRule

// SchedulerRule is a generated Sophos object
type SchedulerRule struct {
	Locked          string `json:"_locked"`
	Reference       string `json:"_ref"`
	ObjectType      string `json:"_type"`
	Comment         string `json:"comment"`
	Destination     string `json:"destination"`
	Interface       string `json:"interface"`
	InterfaceGroup  string `json:"interface_group"`
	Name            string `json:"name"`
	Persistence     bool   `json:"persistence"`
	PersistenceHash string `json:"persistence_hash"`
	Service         string `json:"service"`
	SkipOnError     bool   `json:"skip_on_error"`
	Source          string `json:"source"`
	Status          bool   `json:"status"`
}

var _ sophos.RestGetter = &SchedulerRule{}

// GetPath implements sophos.RestObject and returns the SchedulerRules GET path
// Returns all available scheduler/rule objects
func (*SchedulerRules) GetPath() string { return "/api/objects/scheduler/rule/" }

// RefRequired implements sophos.RestObject
func (*SchedulerRules) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SchedulerRules GET path
// Returns all available rule types
func (s *SchedulerRule) GetPath() string {
	return fmt.Sprintf("/api/objects/scheduler/rule/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *SchedulerRule) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SchedulerRule DELETE path
// Creates or updates the complete object rule
func (*SchedulerRule) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/rule/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SchedulerRule PATCH path
// Changes to parts of the object rule types
func (*SchedulerRule) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/rule/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SchedulerRule POST path
// Create a new scheduler/rule object
func (*SchedulerRule) PostPath() string {
	return "/api/objects/scheduler/rule/"
}

// PutPath implements sophos.RestObject and returns the SchedulerRule PUT path
// Creates or updates the complete object rule
func (*SchedulerRule) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/rule/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SchedulerRule) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/scheduler/rule/%s/usedby", ref)
}

// GetType implements sophos.Object
func (s *SchedulerRule) GetType() string { return s.ObjectType }
