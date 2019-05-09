// consignment-service/main.go
package main

import (

	// 导入 protobuf
	"fmt"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/pengxianghu/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/pengxianghu/shipper/vessel-service/proto/vessel"
	"os"
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
