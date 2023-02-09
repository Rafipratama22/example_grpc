package container

import (
	"proto_example/app/config/mongoconfig"
	"proto_example/app/repository"
	"proto_example/app/usecase"
	"proto_example/app/model/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Container struct {
	GrpcService *grpc.Server
}

func SetupContainer() Container {
	clientMongo := mongoconfig.SetupMongo()
	repository := repository.NewRepo(clientMongo.Client)
	usecase := usecase.NewUsecase(repository)
	grpcServer := grpc.NewServer()
	pb.RegisterPostServiceServer(grpcServer, usecase)
	reflection.Register(grpcServer)

	return Container{
		GrpcService: grpcServer,
	}
}