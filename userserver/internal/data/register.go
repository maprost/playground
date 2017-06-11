package data

import "github.com/maprost/pqx"

func Register() {
	err := pqx.Register(UserAccount{})
	if err != nil {
		panic(err)
	}
}
