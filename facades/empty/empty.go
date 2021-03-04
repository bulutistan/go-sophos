package empty

import (
	"github.com/bulutistan/go-sophos/api/v1.3.0/objects"
	"github.com/bulutistan/go-sophos/facades"
	"github.com/bulutistan/go-sophos/sophos"
)

type Empty struct {
	connection facades.Connection
	nt         objects.PacketfilterPacketfilter
	nts        objects.PacketfilterPacketfilters
}

func (s Empty) Init(c facades.Connection) Empty {
	s.connection = c

	s.nt = objects.PacketfilterPacketfilter{}

	return s
}

func (s Empty) Make(nt objects.PacketfilterPacketfilter) Empty {
	s.nt = nt

	return s
}

func (s Empty) FetchAll() Empty {
	_ = s.connection.Client.GetObject(&s.nts)

	return s
}

func (s Empty) GetByReference(reference string) objects.PacketfilterPacketfilter {
	var nt objects.PacketfilterPacketfilter
	nt.Reference = reference

	_ = s.connection.Client.GetObject(&nt)

	return nt
}

func (s Empty) FindByReference(reference string) Empty {
	s.nt.Reference = reference

	_ = s.connection.Client.GetObject(&s.nt)

	return s
}

// TODO Here..........

func (s Empty) Send(position int) error {
	return s.connection.Client.PostObject(&s.nt,
		sophos.WithRestdInsert("network.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s Empty) Update(position int) error {
	return s.connection.Client.PutObject(&s.nt,
		sophos.WithRestdInsert("network.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}

func (s Empty) Delete(position int) error {
	return s.connection.Client.DeleteObject(&s.nt,
		sophos.WithRestdInsert("network.rules", position),
		sophos.WithRestdLockOverride,
		sophos.AutoResolveErrsMode,
		sophos.WithSessionClose)
}
