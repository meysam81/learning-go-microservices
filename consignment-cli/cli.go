package main

import (
	"context"
	"encoding/json"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	pb "github.com/meysam81/learning-go-microservices/consignment-service/proto/consignment"
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
	cmd.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

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

	for _, v := range getAll.Consignments {
		log.Println(v)
	}

}
