package main

import (
	"github.com/bulutistan/go-sophos/facades"
	_ "github.com/bulutistan/go-sophos/facades/packetfilter"
	"github.com/bulutistan/go-sophos/utils"
	"log"
)

func main() {
	getEnv := new(utils.GetEnviroment).Init()
	conn, err := new(facades.Connection).Init(getEnv)

	log.Print(conn)
	log.Print(err)
}
