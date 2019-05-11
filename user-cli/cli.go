package main

import (
	"log"

	pb "github.com/pengxianghu/shipper/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"github.com/micro/go-micro"
)


func main() {

	srv := micro.NewService(

        micro.Name("go.micro.srv.user-cli"),
        micro.Version("latest"),
    )

    // Init will parse the command line flags.
    srv.Init()

    client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

    name := "hupx"
    email := "hpx@163.com"
    password := "p"
    company := "ali"

    r, err := client.Create(context.TODO(), &pb.User{
        Name:     name,
        Email:    email,
        Password: password,
        Company:  company,
    })
    if err != nil {
        log.Fatalf("Could not create: %v", err)
    }
    log.Printf("Created: %s", r.User.Id)

    getAll, err := client.GetAll(context.Background(), &pb.Request{})
    if err != nil {
        log.Fatalf("Could not list users: %v", err)
    }
    for _, v := range getAll.Users {
        log.Println(v)
	}
	

}
