package examples

import (
	"github.com/bulutistan/go-sophos/api/v1.3.0/objects"
	"github.com/bulutistan/go-sophos/facades"
	"github.com/bulutistan/go-sophos/facades/nats"
	"log"
)

func delete() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nat := new(nats.PacketFilterNats).Init(conn)
	nat.FindByReference("REF_PacNatAnyFromTest").Delete()

}

func update() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nat := new(nats.PacketFilterNats).Init(conn)
	nat.FindByReference("REF_PacNatAnyFromTest").SetStatus(true).SetName("test").SetComment("okokokok").Update()

}

func getAll() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nats := new(nats.PacketFilterNats).Init(conn).FetchAll()

	log.Print(nats)

}

func add() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nat := new(nats.PacketFilterNats).Init(conn)

	err = nat.SendOne(&objects.PacketfilterNat{
		Locked:                "",
		Reference:             "",
		ObjectType:            "packetfilter/nat",
		AutoPfIn:              "",
		AutoPfrule:            false,
		Comment:               "deneme 123 bu bir deneme",
		Destination:           "REF_NetworkAny",
		DestinationNatAddress: "",
		DestinationNatService: "",
		Group:                 "",
		Ipsec:                 false,
		Log:                   false,
		Mode:                  "snat",
		Name:                  "Any from deneme to deneme33112",
		Service:               "REF_ServiceAny",
		Source:                "REF_NetHosDemo",
		SourceNatAddress:      "REF_NetIntDemoAddr",
		SourceNatService:      "",
		Status:                true,
	}, 1)

	log.Fatal(err)

}
