package main

import (
	"proto_example/app/interface/container"
	"proto_example/app/interface/server"
)

func main() {
	server.StartServer(container.SetupContainer())
}