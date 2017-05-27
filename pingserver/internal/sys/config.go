package sys

import (
	"github.com/maprost/gox/gxcfg"
)

func InitConfig() {
	gxcfg.InitConfig("config.gox", gxcfg.DatabaseAccessPort)
}

func GetPort() string {
	return gxcfg.GetConfig().Port
}
