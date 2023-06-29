package services

import (
	"aqrus/Microservice/gateway-api/initializers"
	"aqrus/Microservice/gateway-api/pb/products"
	"aqrus/Microservice/gateway-api/services"
	"aqrus/Microservice/gateway-api/user"
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	initializers.ConnectToBD()
	RunGRPCServer()
	initializers.DB.AutoMigrate(&user.User{})
	user.RunServer(initializers.DB)
}

func RunGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// Register your service implementation with the gRPC server here
	products.RegisterProductServiceServer(s, &services.ProductServer{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:50051")
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = products.RegisterProductServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	if err := gwServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


