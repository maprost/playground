package data

import (
	"github.com/maprost/pqx"

	"github.com/maprost/playground/userserver/internal/sys"
)

func OpenDB() {
	err := pqx.OpenDatabaseConnection(sys.GetDBConfig())
	if err != nil {
		panic(err)
	}
}
