package test

import (
	"github.com/maprost/assertion"
	"github.com/maprost/pqx"

	"github.com/maprost/playground/userserver/internal/data"
	"github.com/maprost/playground/userserver/internal/sys"
)

var initDB = false

func InitUnitTest(t assertion.TestEnvironment) assertion.Assert {
	assert := assertion.New(t)

	if initDB == false {
		sys.InitConfig()
		data.OpenDB()
		data.Register()
		initDB = true
	}

	return assert
}

func CloseDB() {
	initDB = false

	err := pqx.DB.Close()
	panic(err)
}
