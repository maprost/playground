package test

import (
	"github.com/maprost/assertion"
	"github.com/maprost/pqx"

	"github.com/maprost/playground/userserver/internal/data"
)

var initDB = false

func InitUnitTest(t assertion.TestEnvironment) assertion.Assert {
	assert := assertion.New(t)

	if initDB == false {
		data.OpenDB()
		initDB = true
	}

	return assert
}

func CloseDB() {
	initDB = false

	err := pqx.DB.Close()
	panic(err)
}
