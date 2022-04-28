package main

import (
	"fmt"
	"httpdatatime/server"

	"net/http"
)

var IP = ""
var PORT = "3000"

func main() {
	server.SetHandlers()
	http.ListenAndServe(fmt.Sprintf("%s:%s", IP, PORT), nil)

}
