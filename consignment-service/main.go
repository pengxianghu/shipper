// consignment-service/main.go
package main

import (

	// 导入 protobuf
	"fmt"
	"log"
	"errors"

	"golang.org/x/net/context"

	"github.com/micro/go-micro"
	pb "github.com/pengxianghu/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/pengxianghu/shipper/vessel-service/proto/vessel"
	"os"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/client"
	userService "github.com/pengxianghu/shipper/user-service/proto/user"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	// 从环境变量中导入 Database host
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	// Mgo 创建的主 session必须于 main() 函数结束之前关闭
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	//创建一个新service
	srv := micro.NewService(
		// name 必须和 protobuf 定义的包名匹配
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		// Our auth middleware
        micro.WrapHandler(AuthWrapper),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// Init()用于解析命令行参数
	srv.Init()

	// 注册调度器Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	// 启动server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
    return func(ctx context.Context, req server.Request, resp interface{}) error {
        meta, ok := metadata.FromContext(ctx)
        if !ok {
            return errors.New("no auth meta-data found in request")
        }

        // Note this is now uppercase (not entirely sure why this is...)
        token := meta["Token"]
        log.Println("Authenticating with token: ", token)

        // Auth here
        authClient := userService.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
        _, err := authClient.ValidateToken(context.Background(), &userService.Token{
            Token: token,
        })
        if err != nil {
            return err
        }
        err = fn(ctx, req, resp)
        return err
    }
}