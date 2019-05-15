// consignment-cli/cli.go

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	pb "github.com/pengxianghu/shipper/consignment-service/proto/consignment"
	"golang.org/x/net/context"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
	defaultToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiMzQ5MjIxNzAtNzQwNi00MWRlLWIyNzctYWNmN2NiNmNlZmVlIiwibmFtZSI6Imh1cHgiLCJjb21wYW55IjoiQUJDIiwiZW1haWwiOiJodUBnbWFpbC5jb20iLCJwYXNzd29yZCI6IiQyYSQxMCQ2Zk9JQ2plNldDUmFmbEZKc1hrRlAuSjFyWUxnVzQyRXlRRXRGbk5xL0p6UGhZeG84QUVwVyJ9LCJleHAiOjE1NTgxMDU1MjQsImlzcyI6ImdvLm1pY3JvLnNydi51c2VyIn0.jPG83TpVq6NSy8WKy5ETL2uUCJGrgHqkUeDPUCzx2OY"
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

	// create a new client
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	file := defaultFilename
	token := defaultToken

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
		return 
	}

	// Create a new context which contains our given token.
	// This same context will be passed into both the calls we make
	// to our consignment-service.
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// create a new consignment
	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
		return 
	}
	log.Printf("Created: %t", r.Created)

	// get all consignments
	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
		return 
	}
	for k, v := range getAll.Consignments {
		log.Printf("consignment %d: %v", k, v)
	}
}
