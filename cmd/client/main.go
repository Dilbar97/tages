package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"tages/internal/config"

	"tages/internal/client"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.GetConfigs()

	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatalln("Missing file path")
	}

	conn, err := grpc.Dial(fmt.Sprintf(":%d", conf.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := upload.NewClient(conn)
	name, err := client.Upload(ctx, flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(name)
}
