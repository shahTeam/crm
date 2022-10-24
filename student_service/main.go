package main

import (
	"fmt"
	"log"
	"net"
	"student/config"
	"student/domain/group"
	"student/domain/student"
	"student/repository"
	"student/server"
	"student/service"

	"github.com/shahTeam/crmconnect/id"
	"github.com/shahTeam/crmprotos/studentpb"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("error with config", err)
	}
	repo, err := repository.NewPostgres(cfg.Config)
	if err != nil {
		panic(err)
	}
	groupFactory := group.NewFactory(id.Generater{})
	studentFactory := student.NewFactory(id.Generater{})
	svc := service.New(repo, studentFactory,groupFactory)
	svr := server.New(svc, studentFactory, groupFactory)

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(grpcServer, svr)
	fmt.Println("Server starting at:", lis.Addr().String())

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}