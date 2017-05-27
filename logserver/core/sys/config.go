package sys

import (
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/pqx/pqdep"
)

func InitConfig() {
	gxcfg.InitConfig("config.gox", gxcfg.DatabaseAccessPort)
}

