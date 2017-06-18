package sys

import (
	"flag"
	"github.com/maprost/gox/gxarg"
	"github.com/maprost/gox/gxcfg"
)

var configFile string

func init() {
	gxarg.ConfigFileVar(&configFile)
	flag.Parse()
}

func InitConfig() {
	gxcfg.InitConfig(configFile, true)
}

func GetPort() string {
	return gxcfg.GetConfig().Port
}
