package main

import (
	"wgin/core"
)

//go:generate go mod tidy

// @title Operation API
// @version 1.0
// @description This is a TEST server.
// @Tags Operation API
// @host 127.0.0.1:8000
func main() {
	core.RunServer()
}
