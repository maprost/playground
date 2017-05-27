package client

import "github.com/maprost/restclient"

func Ping(basePath string) restclient.Result {
	return restclient.Get(basePath + "/ping").Send()
}
