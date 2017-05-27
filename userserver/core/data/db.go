package data

import (
	"github.com/maprost/pqx"

	"github.com/maprost/playground/userserver/core/sys"
)

func OpenDB() {
	err := pqx.OpenDatabaseConnection(sys.GetDBConfig())
	if err != nil {
		panic(err)
	}
}
