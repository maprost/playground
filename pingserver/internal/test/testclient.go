package test

import (
	"github.com/maprost/assertion"
	"github.com/maprost/restclient"
	"github.com/maprost/restclient/rctest"

	"github.com/maprost/playground/pingserver/client"
)

func Ping204(assert assertion.Assert) {
	Ping(assert, rctest.Status204())
}

func Ping(assert assertion.Assert, expected restclient.Result) {
	result := client.Ping(getTestServerUrl())
	rctest.AssertResult(assert, result, expected)
}