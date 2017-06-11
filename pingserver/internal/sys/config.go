package sys

import (
	"github.com/maprost/gox/gxcfg"
)

func InitConfig() {
	gxcfg.InitConfig("local.gx")
}

func GetPort() string {
	return gxcfg.GetConfig().Port
}
