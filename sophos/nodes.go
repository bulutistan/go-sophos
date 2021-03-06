package sophos

// An Endpoint represents a UTM endpoint which is defined by its definitions endpoint
// GET /api/definitions/aws
type Endpoint interface {
	// Definition returns the endpoint Definition
	Definition() Definition
	// RestObjects returns a map of the Node's editable Objects
	RestObjects() map[string]RestObject
	// ApiRoutes returns all known Node paths
	ApiRoutes() []string
}

// A Node is a resource within Sophos UTM that you can update but cannot create, delete, or move.
// Nodes are organized hierarchically as in a filesystem. An example of a node is the “Shell Access”
// service module within Sophos UTM (identified as the “ssh” node). To enable the Shell Access service,
// you can set the ssh.status leaf node.
//
// Unlike filesystems which use a slash “/” to separate the different nodes, Sophos UTM RESTful API uses
// a dot “.” to separate different nodes. Nodes reference objects, for example “Shell Access” and have a
// leaf node of allowed_networks which is an array of references to network objects.
type Node interface {
	Get(c ClientInterface, oo ...Option) error
	Update(c ClientInterface, oo ...Option) error
}

// An Object resides in collections, which you can create, change, or delete. The collections of objects are
// predefined into classes and types. A class describes the general purpose of the objects whereas the type
// describes the required data for the object. Some objects need to be referenced by a node, to work properly
// (e.g. packetfilter/packetfilter and packetfilter/nat).
//
// For example, one of most used class is “network”. The network class describes objects like “host” for real hosts
// (e.g. 192.168.0.1) or “network” for subnets (e.g. 192.168.0.0/24). The collection of objects is always expressed
// with a class and a type, e.g. “network/host” or “network/network”.
type Object interface {
	// GetType returns the objects Type
	GetType() string
}

// UsedBy represents which Objects and Nodes use this RestObject
type UsedBy struct {
	Nodes   []Reference
	Objects []Reference
}

// RestObject is an interface for REST objects
type RestObject interface {
	RestGetter
	DeletePath(ref string) string // DeletePath returns the DELETE path of the Object
	PatchPath(ref string) string  // PatchPath returns the PATCH path of the Object
	PostPath() string             // PostPath returns the POST path of the Object
	PutPath(ref string) string    // GetPath returns the PUT path of the Object
	UsedByPath(ref string) string // UsedByPath returns the usedby URL path to query for UsedBy data
}

// RestGetter is an interface ensuring a Reference is passed when required
// Additionally used to GET pluralized Objects like objects.AmazonVpcConnections
type RestGetter interface {
	GetPath() string             // GetPath returns the GET path of the Object, will sometimes require Reference
	RefRequired() (string, bool) // RefRequired returns true if Ref required
}
