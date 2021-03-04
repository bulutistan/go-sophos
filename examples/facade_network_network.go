package examples

import (
	"github.com/bulutistan/go-sophos/api/v1.3.0/objects"
	"github.com/bulutistan/go-sophos/facades"
	"github.com/bulutistan/go-sophos/facades/network"
	"log"
)

func nt_delete() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nat := new(network.Network).Init(conn)
	nat.FindByReference("REF_PacNatAnyFromTest").Delete()

}

func nt_update() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nat := new(network.Network).Init(conn)
	nat.FindByReference("REF_PacNatAnyFromTest").SetName("test").SetComment("okokokok").Update()

}

func nt_getAll() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nats := new(network.Network).Init(conn).FetchAll()

	log.Print(nats)

}

func nt_add() {
	conn, err := new(facades.Connection).Init()

	if err != nil {
		log.Fatal(err)
	}

	nat := new(network.Network).Init(conn)

	err = nat.SendOne(&objects.NetworkNetwork{
		Locked:     "",
		Reference:  "",
		ObjectType: "",
		Address:    "10.0.0.0",
		Address6:   "",
		Comment:    "deneme 10-0-0-0",
		Interface:  "",
		Name:       "deneme 000",
		Netmask:    "24",
		Netmask6:   "",
		Resolved:   false,
		Resolved6:  false,
	}, 1)

	log.Fatal(err)

}
