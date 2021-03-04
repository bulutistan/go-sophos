package facades

import (
	"crypto/tls"
	"github.com/bulutistan/go-sophos/sophos"
	"github.com/bulutistan/go-sophos/utils"
	"log"
	"net/http"
)

type Connection struct {
	Client *sophos.Client
	GetEnv utils.GetEnviroment
}

func (s Connection) Init(getEnv utils.GetEnviroment) (Connection, error) {
	s.GetEnv = getEnv
	var err error

	if s.GetEnv.EnvUserName == "token" {
		s.Client, err = sophos.New(
			s.GetEnv.EnvEndpoint,
			sophos.WithAPIToken(s.GetEnv.EnvPassword),
		)
	} else {
		s.Client, err = sophos.New(
			s.GetEnv.EnvEndpoint,
			sophos.WithBasicAuth(
				s.GetEnv.EnvUserName,
				s.GetEnv.EnvPassword,
			),
		)
	}

	if err != nil {
		log.Fatal(err)
	} else {
		if s.GetEnv.EnvInsecure == "yes" {
			s.Client.SetDefaultHTTPClient(http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			})
		}
	}

	return s, err
}
