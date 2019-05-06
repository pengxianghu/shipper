// consignment-cli/cli.go

package main

import (  
    "encoding/json"
    "io/ioutil"
    "log"
    "os"

    pb "github.com/pengxianghu/shipper/consignment-service/proto/consignment"
    "golang.org/x/net/context"
    "github.com/micro/go-micro/cmd"
    microclient "github.com/micro/go-micro/client"
)

const (  
    address         = "localhost:50051"
    defaultFilename = "consignment.json"
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
    // 创建和服务器的一个连接
    cmd.Init()
    // Create new greeter client
    client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)
    // 和服务器通信，并打印出返回信息
    file := defaultFilename
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
	

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
    if err != nil {
        log.Fatalf("Could not list consignments: %v", err)
    }
    for _, v := range getAll.Consignments {
        log.Println(v)
    }
}