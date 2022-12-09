package core

import (
	"github.com/fvbock/endless"
	"time"
	"wgin/initialize"
)

func RunServer() {
	routers := initialize.Routers()
	endServer := endless.NewServer("localhost:8000", routers)
	endServer.ReadHeaderTimeout = time.Second * 10
	endServer.WriteTimeout = time.Second * 10
	endServer.MaxHeaderBytes = 1 << 20
	println(endServer.ListenAndServe().Error())
}
