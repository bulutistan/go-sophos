package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// UserPreferences is a generated struct representing the Sophos UserPreferences Endpoint
// GET /api/nodes/user_preferences
type UserPreferences struct {
	UserPreferencesGroup    UserPreferencesGroup    `json:"user_preferences_group"`
	UserPreferencesWebadmin UserPreferencesWebadmin `json:"user_preferences_webadmin"`
}

var defsUserPreferences = map[string]sophos.RestObject{
	"UserPreferencesGroup":    &UserPreferencesGroup{},
	"UserPreferencesWebadmin": &UserPreferencesWebadmin{},
}

// RestObjects implements the sophos.Node interface and returns a map of UserPreferences's Objects
func (UserPreferences) RestObjects() map[string]sophos.RestObject { return defsUserPreferences }

// GetPath implements sophos.RestGetter
func (*UserPreferences) GetPath() string { return "/api/nodes/user_preferences" }

// RefRequired implements sophos.RestGetter
func (*UserPreferences) RefRequired() (string, bool) { return "", false }

var defUserPreferences = &sophos.Definition{Description: "user_preferences", Name: "user_preferences", Link: "/api/definitions/user_preferences"}

// Definition returns the /api/definitions struct of UserPreferences
func (UserPreferences) Definition() sophos.Definition { return *defUserPreferences }

// ApiRoutes returns all known UserPreferences Paths
func (UserPreferences) ApiRoutes() []string {
	return []string{
		"/api/objects/user_preferences/group/",
		"/api/objects/user_preferences/group/{ref}",
		"/api/objects/user_preferences/group/{ref}/usedby",
		"/api/objects/user_preferences/webadmin/",
		"/api/objects/user_preferences/webadmin/{ref}",
		"/api/objects/user_preferences/webadmin/{ref}/usedby",
	}
}

// References returns the UserPreferences's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (UserPreferences) References() []string {
	return []string{
		"REF_UserPreferencesGroup",
		"REF_UserPreferencesWebadmin",
	}
}

// UserPreferencesGroup is an Sophos Endpoint subType and implements sophos.RestObject
type UserPreferencesGroup []interface{}

// GetPath implements sophos.RestObject and returns the UserPreferencesGroup GET path
// Returns all available user_preferences/group objects
func (*UserPreferencesGroup) GetPath() string { return "/api/objects/user_preferences/group/" }

// RefRequired implements sophos.RestObject
func (*UserPreferencesGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the UserPreferencesGroup DELETE path
// Creates or updates the complete object group
func (*UserPreferencesGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/user_preferences/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the UserPreferencesGroup PATCH path
// Changes to parts of the object group types
func (*UserPreferencesGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/user_preferences/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the UserPreferencesGroup POST path
// Create a new user_preferences/group object
func (*UserPreferencesGroup) PostPath() string {
	return "/api/objects/user_preferences/group/"
}

// PutPath implements sophos.RestObject and returns the UserPreferencesGroup PUT path
// Creates or updates the complete object group
func (*UserPreferencesGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/user_preferences/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*UserPreferencesGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/user_preferences/group/%s/usedby", ref)
}

// UserPreferencesWebadmins is an Sophos Endpoint subType and implements sophos.RestObject
type UserPreferencesWebadmins []UserPreferencesWebadmin

// UserPreferencesWebadmin is a generated Sophos object
type UserPreferencesWebadmin struct {
	Locked               string   `json:"_locked"`
	Reference            string   `json:"_ref"`
	_type                string   `json:"_type"`
	BrowserTitle         string   `json:"browser_title"`
	Comment              string   `json:"comment"`
	DashboardAutogroup   bool     `json:"dashboard_autogroup"`
	DashboardLeftcolumn  []string `json:"dashboard_leftcolumn"`
	DashboardRightcolumn []string `json:"dashboard_rightcolumn"`
	ItemsPerPage         string   `json:"items_per_page"`
	Language             string   `json:"language"`
	MarketingWindow      bool     `json:"marketing_window"`
	Name                 string   `json:"name"`
	Shortcuts            struct {
		Aaa       []string `json:"aaa"`
		Epp       []string `json:"epp"`
		Interface []string `json:"interface"`
		Network   []string `json:"network"`
		Searchbox []string `json:"searchbox"`
		Service   []string `json:"service"`
	} `json:"shortcuts"`
	SkipTermsOfUse bool `json:"skip_terms_of_use"`
}

// GetPath implements sophos.RestObject and returns the UserPreferencesWebadmins GET path
// Returns all available user_preferences/webadmin objects
func (*UserPreferencesWebadmins) GetPath() string { return "/api/objects/user_preferences/webadmin/" }

// RefRequired implements sophos.RestObject
func (*UserPreferencesWebadmins) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the UserPreferencesWebadmins GET path
// Returns all available webadmin types
func (u *UserPreferencesWebadmin) GetPath() string {
	return fmt.Sprintf("/api/objects/user_preferences/webadmin/%s", u.Reference)
}

// RefRequired implements sophos.RestObject
func (u *UserPreferencesWebadmin) RefRequired() (string, bool) { return u.Reference, true }

// DeletePath implements sophos.RestObject and returns the UserPreferencesWebadmin DELETE path
// Creates or updates the complete object webadmin
func (*UserPreferencesWebadmin) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/user_preferences/webadmin/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the UserPreferencesWebadmin PATCH path
// Changes to parts of the object webadmin types
func (*UserPreferencesWebadmin) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/user_preferences/webadmin/%s", ref)
}

// PostPath implements sophos.RestObject and returns the UserPreferencesWebadmin POST path
// Create a new user_preferences/webadmin object
func (*UserPreferencesWebadmin) PostPath() string {
	return "/api/objects/user_preferences/webadmin/"
}

// PutPath implements sophos.RestObject and returns the UserPreferencesWebadmin PUT path
// Creates or updates the complete object webadmin
func (*UserPreferencesWebadmin) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/user_preferences/webadmin/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*UserPreferencesWebadmin) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/user_preferences/webadmin/%s/usedby", ref)
}

// GetType implements sophos.Object
func (u *UserPreferencesWebadmin) GetType() string { return u._type }