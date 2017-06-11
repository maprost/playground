package main

import (
	"fmt"
	"github.com/maprost/playground/pingserver/internal"
	"github.com/maprost/playground/pingserver/internal/sys"
	"net/http"
)

func main() {
	sys.InitConfig()

	router := internal.InitServer()
	fmt.Println("Port:" + sys.GetPort())
	http.ListenAndServe(":"+sys.GetPort(), router)
}
