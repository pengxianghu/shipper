// shippy-email-service
package main

import (
    "log"
    pb "github.com/pengxianghu/shipper/user-service/proto/user"
    micro "github.com/micro/go-micro"
	"context"
)

const topic = "user.created"

type Subscriber struct{}

func main() {
    srv := micro.NewService(
        micro.Name("go.micro.srv.email"),
        micro.Version("latest"),
    )

    srv.Init()
	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

    // Run the server
    if err := srv.Run(); err != nil {
        log.Println(err)
    }
}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
    log.Println("Picked up a new message")
    log.Println("Sending email to:", user.Name)
    return nil
}