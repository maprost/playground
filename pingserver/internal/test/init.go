package test

import (
	"github.com/maprost/assertion"
	"github.com/maprost/playground/pingserver/internal"
	"net/http/httptest"
)

var initServer = false
var server *httptest.Server = nil

func InitRestTest(t assertion.TestEnvironment) assertion.Assert {
	assert := assertion.New(t)

	if initServer == false {
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
