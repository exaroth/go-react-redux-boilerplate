package main

import (
	"fmt"
	"net/http"
)

func main() {
	config, err := getConfig()
	if err != nil {
		panic(err)
	}
	initRoutes(config)
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Hostname, config.Port), nil)
}
