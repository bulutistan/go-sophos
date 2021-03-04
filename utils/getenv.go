package utils

import (
	"os"
	"strings"
)

type GetEnviroment struct {
	envEndpoint string
	envUserName string
	envPassword string
	envInsecure string

	EnvEndpoint string
	EnvUserName string
	EnvPassword string
	EnvInsecure string
}

func (g GetEnviroment) Init() GetEnviroment {
	g.envEndpoint = "SOPHOS_ENDPOINT"
	g.envUserName = "SOPHOS_USERNAME"
	g.envPassword = "SOPHOS_PASSWORD"
	g.envInsecure = "SOPHOS_INSECURE"

	g.EnvEndpoint = os.Getenv(g.envEndpoint)
	g.EnvUserName = os.Getenv(g.envUserName)
	g.EnvPassword = os.Getenv(g.envPassword)
	g.EnvInsecure = os.Getenv(g.envInsecure)

	return g
}

// getEnvString returns string from environment variable.
func (GetEnviroment) getEnvString(v string, def string) string {
	r := os.Getenv(v)
	if r == "" {
		return def
	}

	return r
}

// getEnvBool returns boolean from environment variable.
func (GetEnviroment) getEnvBool(v string, def bool) bool {
	r := os.Getenv(v)
	if r == "" {
		return def
	}

	switch strings.ToLower(r[0:1]) {
	case "t", "y", "1":
		return true
	}

	return false
}
