package main

import (
	"github.com/fvbock/endless"
	"time"
	"wgin/initialize"
)

//go:generate go mod tidy

// @title Operation API
// @version 1.0
// @description This is a TEST server.
// @Tags Operation API
// @host 127.0.0.1:8000
func main() {
	routers := initialize.Routers()
	endServer := endless.NewServer("localhost:8000", routers)
	endServer.ReadHeaderTimeout = time.Second * 10
	endServer.WriteTimeout = time.Second * 10
	endServer.MaxHeaderBytes = 1 << 20
	println(endServer.ListenAndServe().Error())
}
