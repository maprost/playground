package sys

import (
	"flag"
	"github.com/maprost/gox/gxarg"
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/pqx/pqdep"
	"log"
)

var cfgFile string

func init() {
	gxarg.ConfigFileVar(&cfgFile)
	flag.Parse()
}

func InitConfig() {
	gxcfg.InitConfig(cfgFile, true)
}

func GetDBConfig() pqdep.ConnectionInfo {
	log.Println("DB Size: ", len(gxcfg.GetConfig().Database))
	log.Println("DB Config: ", gxcfg.GetConfig().Database[0])
	return gxcfg.GetConfig().Database[0]
}

func GetPort() string {
	return gxcfg.GetConfig().Port
}
