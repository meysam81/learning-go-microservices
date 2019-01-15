package main

import (
	"context"
	"encoding/json"
	microClient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	pb "gotut/consignment-service/proto/consignment"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	_ = cmd.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microClient.DefaultClient)
	//log.Printf("client: %v" , client)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.TODO(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Created: %t", r.Created)
	if r.Created {
		log.Printf("Consignment: %t", r.Consignment)
	}

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not get all: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}

}
