package server

import (
	"log"
	"net"
	"proto_example/app/interface/container"
)

func StartServer(container container.Container) {
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	err = container.GrpcService.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
