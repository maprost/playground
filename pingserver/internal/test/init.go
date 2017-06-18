package test

import (
	"flag"
	"github.com/maprost/assertion"
	"github.com/maprost/gox/gxarg"
	"github.com/maprost/gox/gxcfg"
	"net/http/httptest"

	"github.com/maprost/playground/pingserver/internal"
)

var initServer = false
var server *httptest.Server = nil
var configFile string

func init() {
	gxarg.ConfigFileVar(&configFile)
	flag.Parse()
}

func InitRestTest(t assertion.TestEnvironment) assertion.Assert {
	assert := assertion.New(t)

	if initServer == false {
		gxcfg.InitConfig(configFile, true)

		router := internal.InitServer()
		server = httptest.NewServer(router)
		initServer = true
	}

	return assert
}

func getTestServerUrl() string {
	if server == nil {
		return "ServerNotInitialized"
	}
	return server.URL
}
