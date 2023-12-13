package main

import (
	"context"
	"net"
	"net/http"

	contracts "github.com/atrop1ne/logger-service-api/gen/go/contracts/v1"
	serverimplement "github.com/atrop1ne/logger-service/serverImplement"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	restPort = ":8081"
	grpcPort = ":8082"
)

func main() {
	grpcServer := grpc.NewServer()
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		panic(err)
	}

	contracts.RegisterLogsServer(
		grpcServer,
		&serverimplement.GRPCServer{UnimplementedLogsServer: contracts.UnimplementedLogsServer{}},
	)

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = contracts.RegisterLogsHandlerFromEndpoint(context.Background(), mux, grpcPort, opts)
	if err != nil {
		panic(err)
	}

	g := new(errgroup.Group)
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		return http.ListenAndServe(restPort, mux)
	})

	err = g.Wait()
	if err != nil {
		panic(err)
	}
}
