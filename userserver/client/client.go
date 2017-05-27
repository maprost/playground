package client

import "github.com/maprost/restclient"

func CreateUser(basePath string, user UserDto) restclient.Result {
	return restclient.Post(basePath + "/user").AddJsonBody(&user).Send()
}
