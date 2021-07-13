package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"net/http"
	practics "practics/api/api"
	service "practics/internal/api"
	"practics/internal/kafka"
	"practics/internal/pkg"
	"practics/internal/repository"
)

func run() error {
	ctx := context.Background()
	db := repository.NewDBBalancer(ctx)
	dao := repository.NewDAO(*db)
	uService := pkg.NewUsersService(dao)
	iService := pkg.NewItemsService(dao)
	pService := pkg.NewPurchasesService(dao)
	service := service.NewPracticsApiService(uService, iService, pService)

	gw := runtime.NewServeMux()

	err := practics.RegisterPurchaseApiServiceHandlerServer(ctx, gw, service)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	practics.RegisterPurchaseApiServiceServer(grpcServer, service)



	var group errgroup.Group

	group.Go(func() error {
		return grpcServer.Serve(listener)
	})


	group.Go(func() error {
		return http.ListenAndServe(":2662", gw)
	})

	group.Go(func() error {
		return kafka.NewConsumer(ctx)
	})

	return group.Wait()
}


func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

