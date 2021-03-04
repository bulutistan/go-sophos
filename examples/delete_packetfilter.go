package examples

import (
	"fmt"
	"log"
	"os"

	types "github.com/bulutistan/go-sophos/api/v1.3.0/objects"
	"github.com/bulutistan/go-sophos/sophos"
)

func init() {
	if len(os.Args) == 3 {
		endpoint = os.Args[1]
		token = os.Args[2]
	}

	if endpoint == "" {
		endpoint = os.Getenv("ENDPOINT")
	}
	if token == "" {
		token = os.Getenv("TOKEN")
	}

	if endpoint == "" || token == "" {
		panic("need endpoint and token as args or from env ($ENDPOINT, $TOKEN)")
	}
}

func main() {
	client, err := sophos.New(endpoint, sophos.WithAPIToken(token))
	if err != nil {
		log.Fatal(err)
	}

	p := types.PacketfilterPacketfilter{
		Reference: "REF_PacPacAnyFromAnyTo2",
	}

	err = client.DeleteObject(&p, sophos.WithSessionClose, sophos.AutoResolveErrsMode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully deleted", p.Reference)
}
