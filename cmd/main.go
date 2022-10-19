package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/routes"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	fmt.Println("gRPC Server started at " + port + "...")
	routes.InitMethods()
}
