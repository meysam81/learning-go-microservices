package main

import (
	pb "../consignment-service/proto/consignment"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
)

const (
	ADDRESS          = "localhost:8000"
	DEFAULT_FILENAME = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not establish connection: %v", err)
	}

	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)

	file := DEFAULT_FILENAME
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Created: %t", r.Created)
	if r.Created {
		log.Printf("Consignment: %t", r.Consignment)
	}

	getAll, err := client.GetConsignment(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not get all: %v", err)
	}

	log.Printf("Created: %t", r.Created)
	for _, v := range getAll.Consignments {
		log.Println(v)
	}

}
