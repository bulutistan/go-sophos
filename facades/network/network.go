package network

import (
	"encoding/json"
	"github.com/bulutistan/go-sophos/api/v1.3.0/objects"
	"github.com/bulutistan/go-sophos/facades"
	"github.com/bulutistan/go-sophos/sophos"
)

type Network struct {
	connection facades.Connection
	nt         objects.NetworkNetwork
	nts        objects.NetworkNetworks
}

func (s Network) Init(c facades.Connection) Network {
	s.connection = c

	s.nt = objects.NetworkNetwork{}

	return s
}

func (s Network) Make(nt objects.NetworkNetwork) Network {
	s.nt = nt

	return s
}

func (s Network) FetchAll() Network {
	_ = s.connection.Client.GetObject(&s.nts)

	return s
}

func (s Network) GetByReference(reference string) objects.NetworkNetwork {
	var nt objects.NetworkNetwork
	nt.Reference = reference

	_ = s.connection.Client.GetObject(&nt)

	return nt
}

func (s Network) FindByReference(reference string) Network {
	s.nt.Reference = reference

	_ = s.connection.Client.GetObject(&s.nt)

	return s
}

func (s Network) SetObjectType(objectType string) Network {
	s.nt.ObjectType = objectType
	return s
}

func (s Network) SetAddress(address string) Network {
	s.nt.Address = address
	return s
}

func (s Network) SetAddress6(address6 string) Network {
	s.nt.Address6 = address6
	return s
}

func (s Network) SetComment(comment string) Network {
	s.nt.Comment = comment
	return s
}

func (s Network) SetInterface(interfaceIs string) Network {
	s.nt.Interface = interfaceIs
	return s
}

func (s Network) SetName(name string) Network {
	s.nt.Name = name
	return s
}

func (s Network) SetNetmask(netmask string) Network {
	s.nt.Netmask = json.Number(netmask)
	return s
}

func (s Network) SetNetmask6(netmask6 string) Network {
	s.nt.Netmask6 = json.Number(netmask6)
	return s
}

func (s Network) SetResolved(isResolved bool) Network {
	s.nt.Resolved = isResolved
	return s
}

func (s Network) SetResolved6(isResolved6 bool) Network {
	s.nt.Resolved6 = isResolved6
	return s
}

func (s Network) SendOne(nt *objects.NetworkNetwork, position int) error {
	return s.connection.Client.PostObject(nt,
		sophos.WithRestdInsert("network.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s Network) Send(position int) error {
	return s.connection.Client.PostObject(&s.nt,
		sophos.WithRestdInsert("network.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s Network) UpdateWithPosition(position int) error {
	return s.connection.Client.PutObject(&s.nt,
		sophos.WithRestdInsert("network.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s Network) Update() error {
	return s.connection.Client.PutObject(&s.nt,
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s Network) Delete() error {
	return s.connection.Client.DeleteObject(&s.nt,
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}
