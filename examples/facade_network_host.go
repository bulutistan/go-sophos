package examples

import (
	"github.com/bulutistan/go-sophos/api/v1.3.0/objects"
	"github.com/bulutistan/go-sophos/facades"
	"github.com/bulutistan/go-sophos/facades/network"
	"log"
)

func nh_delete() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nat := new(network.Host).Init(conn)
	nat.FindByReference("REF_PacNatAnyFromTest").Delete()

}

func nh_update() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nat := new(network.Host).Init(conn)
	nat.FindByReference("REF_PacNatAnyFromTest").SetName("test").SetComment("okokokok").Update()

}

func nh_getAll() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nats := new(network.Host).Init(conn).FetchAll()

	log.Print(nats)

}

func nh_add() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nat := new(network.Host).Init(conn)

	err = nat.SendOne(&objects.NetworkHost{
		Locked:     "",
		Reference:  "", // not aggree it's generate from own
		ObjectType: "network/host",
		Address:    "10.0.0.1",
		Address6:   "",
		Comment:    "deneme 01",
		Duids:      nil,
		Hostnames:  nil,
		Interface:  "",
		Macs:       nil,
		Name:       "ok 001",
		Resolved:   true,
		Resolved6:  false,
		ReverseDNS: false,
	})

	log.Fatal(err)

}
