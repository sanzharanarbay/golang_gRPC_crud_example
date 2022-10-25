# Golang gRPC with the CRUD Methods
- docker-compose build --no-cache
- docker-compose up -d
- docker-compose ps
- docker-compose logs -f {serviceName} 


# Protoc commands (additional commands, if you want to generate new proto services)
- protoc proto\category.proto --go_out=plugins=grpc:serializers
- protoc proto\product.proto --go_out=plugins=grpc:serializers

