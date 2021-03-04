package examples

import (
	"github.com/bulutistan/go-sophos/api/v1.3.0/objects"
	"github.com/bulutistan/go-sophos/facades"
	"github.com/bulutistan/go-sophos/facades/packetfilter"
	"log"
)

func pf_delete() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	pfr := new(packetfilter.PacketFilterRules).Init(conn)
	pfr.FindByReference("REF_PacNatAnyFromTest").Delete()

}

func pf_update() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	pfr := new(packetfilter.PacketFilterRules).Init(conn)
	pfr.FindByReference("REF_PacNatAnyFromTest").SetStatus(true).SetName("test").SetComment("okokokok").Update()

}

func pf_getAll() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	pfrs := new(packetfilter.PacketFilterRules).Init(conn).FetchAll()

	log.Print(pfrs)

}

func pf_add() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	pfr := new(packetfilter.PacketFilterRules).Init(conn)

	err = pfr.SendOne(&objects.PacketfilterPacketfilter{
		Locked:             "",
		Reference:          "",
		ObjectType:         "packetfilter/rules",
		Action:             "",
		Auto:               false,
		AutoType:           "",
		Comment:            "bu bir denemedir",
		Destinations:       nil,
		Direction:          "in",
		Group:              "",
		Interface:          "",
		Log:                false,
		Name:               "deneme 123",
		Services:           nil,
		SourceMacAddresses: "",
		Sources:            nil,
		Status:             true,
		Time:               "",
	}, 1)

	log.Fatal(err)

}
