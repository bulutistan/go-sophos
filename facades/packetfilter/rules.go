package packetfilter

import (
	"github.com/bulutistan/go-sophos/api/v1.3.0/objects"
	"github.com/bulutistan/go-sophos/facades"
	"github.com/bulutistan/go-sophos/sophos"
)

type PacketFilterRules struct {
	connection facades.Connection
	nt         objects.PacketfilterPacketfilter
	nts        objects.PacketfilterPacketfilters
}

func (s PacketFilterRules) Init(c facades.Connection) PacketFilterRules {
	s.connection = c

	s.nt = objects.PacketfilterPacketfilter{}

	return s
}

func (s PacketFilterRules) Make(nt objects.PacketfilterPacketfilter) PacketFilterRules {
	s.nt = nt

	return s
}

func (s PacketFilterRules) FetchAll() PacketFilterRules {
	_ = s.connection.Client.GetObject(&s.nts)

	return s
}

func (s PacketFilterRules) GetByReference(reference string) objects.PacketfilterPacketfilter {
	var nt objects.PacketfilterPacketfilter
	nt.Reference = reference

	_ = s.connection.Client.GetObject(&nt)

	return nt
}

func (s PacketFilterRules) FindByReference(reference string) PacketFilterRules {
	s.nt.Reference = reference

	_ = s.connection.Client.GetObject(&s.nt)

	return s
}

func (s PacketFilterRules) SetStatus(status bool) PacketFilterRules {
	s.nt.Status = status
	return s
}

func (s PacketFilterRules) SetAuto(isAuto bool, autoType string) PacketFilterRules {
	s.nt.Auto = isAuto
	s.nt.AutoType = autoType
	return s
}

func (s PacketFilterRules) SetComment(comment string) PacketFilterRules {
	s.nt.Comment = comment
	return s
}

func (s PacketFilterRules) SetDirection(direction string) PacketFilterRules {
	s.nt.Direction = direction
	return s
}

func (s PacketFilterRules) SetName(name string) PacketFilterRules {
	s.nt.Name = name
	return s
}

func (s PacketFilterRules) SetAction(action string) PacketFilterRules {
	s.nt.Action = action
	return s
}

func (s PacketFilterRules) SetLog(isLogging bool) PacketFilterRules {
	s.nt.Log = isLogging
	return s
}

func (s PacketFilterRules) SetDestinations(destinations []string) PacketFilterRules {
	s.nt.Destinations = destinations
	return s
}

func (s PacketFilterRules) SetServices(services []string) PacketFilterRules {
	s.nt.Services = services
	return s
}

func (s PacketFilterRules) SetSources(sources []string) PacketFilterRules {
	s.nt.Sources = sources
	return s
}

func (s PacketFilterRules) SetSourceMacAddresses(macAddresses string) PacketFilterRules {
	s.nt.SourceMacAddresses = macAddresses
	return s
}

func (s PacketFilterRules) SetInterface(interfaceIs string) PacketFilterRules {
	s.nt.Interface = interfaceIs
	return s
}

func (s PacketFilterRules) SetGroup(group string) PacketFilterRules {
	s.nt.Group = group
	return s
}

func (s PacketFilterRules) SendOne(nt *objects.PacketfilterPacketfilter, position int) error {
	return s.connection.Client.PostObject(nt,
		sophos.WithRestdInsert("packetfilter.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s PacketFilterRules) Send(position int) error {
	return s.connection.Client.PostObject(&s.nt,
		sophos.WithRestdInsert("packetfilter.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s PacketFilterRules) UpdateWitPosition(position int) error {
	return s.connection.Client.PutObject(&s.nt,
		sophos.WithRestdInsert("packetfilter.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s PacketFilterRules) Update() error {
	return s.connection.Client.PutObject(&s.nt,
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s PacketFilterRules) Delete() error {
	return s.connection.Client.DeleteObject(&s.nt,
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}
