package grpcEngine

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"gitlab.com/RezaOptic/go-utils/logger"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	grpcServer *grpc.Server
)


func Run(Port string) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	var opts []grpc.ServerOption
	// Add unary interceptor
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.UnaryServerInterceptor(logger.ZLogger),
	))

	// Add stream interceptor (added as an example here)
	opts = append(opts, grpc_middleware.WithStreamServerChain(
		grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.StreamServerInterceptor(logger.ZLogger),
	))

	// create a server instance
	grpcServer = grpc.NewServer(opts...)

	// register api service
	registerAPI()

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			grpcServer.GracefulStop()
		}
	}()

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func registerAPI() {
}
