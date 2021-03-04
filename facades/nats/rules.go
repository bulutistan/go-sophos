package nats

import (
	"github.com/bulutistan/go-sophos/api/v1.3.0/objects"
	"github.com/bulutistan/go-sophos/facades"
	"github.com/bulutistan/go-sophos/sophos"
)

type PacketFilterNats struct {
	connection facades.Connection
	nt         objects.PacketfilterNat
	nts        objects.PacketfilterNats
}

func (s PacketFilterNats) Init(c facades.Connection) PacketFilterNats {
	s.connection = c

	s.nt = objects.PacketfilterNat{}

	return s
}

func (s PacketFilterNats) FetchAll() PacketFilterNats {
	_ = s.connection.Client.GetObject(&s.nts)

	return s
}

func (s PacketFilterNats) GetByReference(reference string) objects.PacketfilterNat {
	var nt objects.PacketfilterNat
	nt.Reference = reference

	_ = s.connection.Client.GetObject(&nt)

	return nt
}

func (s PacketFilterNats) FindByReference(reference string) PacketFilterNats {
	s.nt.Reference = reference

	_ = s.connection.Client.GetObject(&s.nt)

	return s
}

func (s PacketFilterNats) SetObjectType(objectType string) PacketFilterNats {
	s.nt.ObjectType = objectType
	return s
}

func (s PacketFilterNats) SetAutoPfIn(autoPfIn string) PacketFilterNats {
	s.nt.AutoPfIn = autoPfIn
	return s
}

func (s PacketFilterNats) SetAutoPfRule(autoPfRule bool) PacketFilterNats {
	s.nt.AutoPfrule = autoPfRule
	return s
}

func (s PacketFilterNats) SetComment(comment string) PacketFilterNats {
	s.nt.Comment = comment
	return s
}

func (s PacketFilterNats) SetDestination(destination string) PacketFilterNats {
	s.nt.Destination = destination
	return s
}

func (s PacketFilterNats) SetDestNatAddress(address string) PacketFilterNats {
	s.nt.DestinationNatAddress = address
	return s
}

func (s PacketFilterNats) SetDestNatService(service string) PacketFilterNats {
	s.nt.DestinationNatService = service
	return s
}

func (s PacketFilterNats) SetGroup(group string) PacketFilterNats {
	s.nt.Group = group
	return s
}

func (s PacketFilterNats) SetIpSec(isIpSec bool) PacketFilterNats {
	s.nt.Ipsec = isIpSec
	return s
}

func (s PacketFilterNats) SetLog(isLogEnable bool) PacketFilterNats {
	s.nt.Log = isLogEnable
	return s
}

func (s PacketFilterNats) SetMode(mode string) PacketFilterNats {
	s.nt.Mode = mode
	return s
}

func (s PacketFilterNats) SetName(name string) PacketFilterNats {
	s.nt.Name = name
	return s
}

func (s PacketFilterNats) SetService(service string) PacketFilterNats {
	s.nt.Service = service
	return s
}

func (s PacketFilterNats) SetSource(source string) PacketFilterNats {
	s.nt.Source = source
	return s
}

func (s PacketFilterNats) SetSrcNatAddress(address string) PacketFilterNats {
	s.nt.SourceNatAddress = address
	return s
}

func (s PacketFilterNats) SetSrcNatService(service string) PacketFilterNats {
	s.nt.SourceNatService = service
	return s
}

func (s PacketFilterNats) SetStatus(isStatus bool) PacketFilterNats {
	s.nt.Status = isStatus
	return s
}

func (s PacketFilterNats) SendOne(nt *objects.PacketfilterNat, position int) error {
	return s.connection.Client.PostObject(nt,
		sophos.WithRestdInsert("nat.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s PacketFilterNats) Send(position int) error {
	return s.connection.Client.PostObject(&s.nt,
		sophos.WithRestdInsert("nat.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s PacketFilterNats) UpdateWithPosition(position int) error {
	return s.connection.Client.PutObject(&s.nt,
		sophos.WithRestdInsert("nat.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s PacketFilterNats) Update() error {
	return s.connection.Client.PutObject(&s.nt,
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s PacketFilterNats) Delete() error {
	return s.connection.Client.DeleteObject(&s.nt,
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}
