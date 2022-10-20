package routes

import (
	serializers "github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/serializers"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/service-servers"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitMethods() {
	srvCategory := service_servers.CategoryServiceServer{}
	srvProduct := service_servers.ProductServiceServer{}
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	serializers.RegisterCategoryServiceServer(grpcServer, &srvCategory)
	serializers.RegisterProductServiceServer(grpcServer, &srvProduct)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
