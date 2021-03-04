package network

import (
	"github.com/bulutistan/go-sophos/api/v1.3.0/objects"
	"github.com/bulutistan/go-sophos/facades"
	"github.com/bulutistan/go-sophos/sophos"
)

type Host struct {
	connection facades.Connection
	nt         objects.NetworkHost
	nts        objects.NetworkHosts
}

func (s Host) Init(c facades.Connection) Host {
	s.connection = c

	s.nt = objects.NetworkHost{}

	return s
}

func (s Host) Make(nt objects.NetworkHost) Host {
	s.nt = nt

	return s
}

func (s Host) FetchAll() Host {
	_ = s.connection.Client.GetObject(&s.nts)

	return s
}

func (s Host) GetByReference(reference string) objects.NetworkHost {
	var nt objects.NetworkHost
	nt.Reference = reference

	_ = s.connection.Client.GetObject(&nt)

	return nt
}

func (s Host) FindByReference(reference string) Host {
	s.nt.Reference = reference

	_ = s.connection.Client.GetObject(&s.nt)

	return s
}

func (s Host) SetObjectType(objectType string) Host {
	s.nt.ObjectType = objectType
	return s
}

func (s Host) SetAddress(address string) Host {
	s.nt.Address = address
	return s
}

func (s Host) SetAddress6(address6 string) Host {
	s.nt.Address6 = address6
	return s
}

func (s Host) SetComment(comment string) Host {
	s.nt.Comment = comment
	return s
}

func (s Host) SetHostnames(hostnames []string) Host {
	s.nt.Hostnames = hostnames
	return s
}

func (s Host) SetInterface(interfaceIs string) Host {
	s.nt.Interface = interfaceIs
	return s
}

func (s Host) SetName(name string) Host {
	s.nt.Name = name
	return s
}

func (s Host) SetResolved(isResolved bool) Host {
	s.nt.Resolved = isResolved
	return s
}

func (s Host) SetResolved6(isResolved6 bool) Host {
	s.nt.Resolved6 = isResolved6
	return s
}

func (s Host) SetReverseDNS(isReverse bool) Host {
	s.nt.ReverseDNS = isReverse
	return s
}

func (s Host) SendOne(nt *objects.NetworkHost) error {
	return s.connection.Client.PostObject(nt,
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s Host) Send() error {
	return s.connection.Client.PostObject(&s.nt,
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s Host) Update() error {
	return s.connection.Client.PutObject(&s.nt,
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s Host) Delete() error {
	return s.connection.Client.DeleteObject(&s.nt,
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}
