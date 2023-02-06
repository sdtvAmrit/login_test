package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"net/http"

	pb "service_1/proto"
	"google.golang.org/grpc"
	"github.com/hashicorp/consul/api"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

const (
	address = "printer:50052"
)

// AddHealthCheck adds a health check endpoint to the gRPC server.
// func (s *server) AddHealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
// 	return &pb.HealthCheckResponse{Status: pb.HealthCheckResponse_HEALTHY}, nil
// }
  

func main() {
	registerToConsul()
	config := api.DefaultConfig()
	config.Address = "consul:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	services, _, err := client.Health().Service("printer", "", true, nil)
	if err != nil {
		log.Fatal(err)
	}
	var printerAddr string
	for _, service := range services {
		printerAddr = fmt.Sprintf("%s:%d", service.Service.Address, service.Service.Port)
	}
	log.Println("Discovered Address: %s", printerAddr)
	conn, err := grpc.Dial(printerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPrinterClient(conn)

	// Start the gRPC server
	// lis, err := net.Listen("tcp", "0.0.0.0:50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	s := grpc.NewServer()
	pb.RegisterAdderServer(s, &server{c: c})
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	mux := runtime.NewServeMux()
	if err := pb.RegisterAdderHandlerFromEndpoint(context.Background(), mux, "localhost:50051", []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		fmt.Println(err)
	}

	// Add health check endpoint to the REST gateway
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.Handle("/", mux)

	if err := http.ListenAndServe(":50051", nil); err != nil {
		fmt.Println(err)
	}
}

// server is the implementation of the Adder service
type server struct {
	c pb.PrinterClient
}

// Add takes two numbers and returns the sum
func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	// Compute the sum of the two numbers
	sum := in.A + in.B

	// Forward the sum to the printer microservice
	log.Println("Do the sum now")
	_, err := s.c.Print(ctx, &pb.PrintRequest{Message: fmt.Sprintf("The sum of %d and %d is %d", in.A, in.B, sum)})
	if err != nil {
		return nil, fmt.Errorf("could not print sum: %v", err)
	}

	return &pb.AddResponse{Sum: sum}, nil
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
        Name: "adder",
        ID:   "adder-1",
        Address: "adder",
        Port: 50051,
        Check: &api.AgentServiceCheck{
            HTTP:     "http://adder:50051/health",
            Interval: "5s",
            Timeout:  "3s",
        },
    })
    if err != nil {
        log.Fatalf("Error registering Adder service with Consul: %v", err)
    }
}