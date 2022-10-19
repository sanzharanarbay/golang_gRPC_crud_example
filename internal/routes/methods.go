package routes

import (
	serializers "github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/serializers"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/service-servers"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitMethods() {
	srv := service_servers.CategoryServiceServer{}
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	serializers.RegisterCategoryServiceServer(grpcServer, &srv)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
