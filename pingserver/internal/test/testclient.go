package test

import (
	"github.com/maprost/assertion"
	"github.com/maprost/restclient"
	"github.com/maprost/restclient/rctest"

	"github.com/maprost/playground/pingserver/clientlib"
)

func Ping200(assert assertion.Assert) {
	Ping(assert, rctest.Status200())
}

func Ping(assert assertion.Assert, expected restclient.Result) {
	result := clientlib.Ping(getTestServerUrl())
	rctest.AssertResult(assert, result, expected)
}
