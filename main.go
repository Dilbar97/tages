package main

import (
	"context"
	"fmt"
	"net"
	"tages/internal/server/repository"
	"tages/internal/server/service"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"

	"tages/internal/config"
	"tages/internal/server"
	"tages/internal/storage"
	pb "tages/protocol/media"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.GetConfigs()

	connPostgres, err := pgxpool.Connect(ctx, conf.DB)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer connPostgres.Close()

	mediaRepo := repository.NewMediaRepo(connPostgres)
	mediaSvc := service.NewMediaSvc(mediaRepo)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GrpcPort))
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to listen: %v", err.Error()))
	}

	mediaSrv := upload.NewMediaServer(storage.New(conf.FileLocation), mediaSvc)
	grpcSrv := grpc.NewServer()

	pb.RegisterUploadServiceServer(grpcSrv, &mediaSrv)
	if err = grpcSrv.Serve(listen); err != nil {
		fmt.Println(fmt.Sprintf("failed to serve grpc: %v", err.Error()))
	}
}
