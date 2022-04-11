package grpc

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewClient() *grpc.ClientConn {
	fmt.Println("utils.grpc.client[NewClient]: ", viper.GetString("port.client"))
	grpcConn, err := grpc.Dial(":"+viper.GetString("port.client"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return grpcConn
}
