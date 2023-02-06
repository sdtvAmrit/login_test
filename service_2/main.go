// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "printer/proto"
	"google.golang.org/grpc"
	"github.com/hashicorp/consul/api"
	
)

type server struct{}

func (s *server) Print(ctx context.Context, req *pb.PrintRequest) (*pb.PrintResponse, error) {
	log.Println("Returning written one reat")
	message := fmt.Sprintf("The sum is: %d", req.Sum)
	return &pb.PrintResponse{Message: message}, nil
}

func main() {
	registerToConsul()
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPrinterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func registerToConsul() {
	config := api.DefaultConfig()
    config.Address = os.Getenv("CONSUL_HTTP_ADDR")
    client, err := api.NewClient(config)
    if err != nil {
        log.Fatalf("Error creating Consul client: %v", err)
    }

    // Register the Adder service with Consul
    err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
        Name: "printer",
        ID:   "printer-1",
        Address: "printer",
        Port: 50052,
        // Check: &api.AgentServiceCheck{
        //     HTTP:     "http://printer:50052/health",
        //     Interval: "5s",
        //     Timeout:  "3s",
        // },
    })
    if err != nil {
        log.Fatalf("Error registering Printer service with Consul: %v", err)
    }
}
