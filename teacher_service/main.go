package main

import (
	"net"
	"teacher/config"
	"teacher/domain/subject"
	"teacher/domain/teacher"
	"teacher/pkg/idgen"
	"teacher/repository"
	"teacher/service"

	"teacher/server"

	"github.com/shahTeam/crmprotos/teacherpb"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	repo, err := repository.NewPostgres(cfg.Config)
	if err != nil {
		panic(err)
	}
	subjectFactory := subject.NewFactory(idgen.Generator{})
	teacherFactory := teacher.NewFactory(idgen.Generator{})

	svc := service.New(repo, subjectFactory, teacherFactory)
	server := server.New(svc, subjectFactory, teacherFactory)

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	teacherpb.RegisterTeacherServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}