// consignment-service/handler.go

package main

import (
	pb "github.com/pengxianghu/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/pengxianghu/shipper/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"log"
)

// Service 应该实现我们在 protobuf 定义中定义的所有方法，检查生成的代码本身中是否有确切的签名方法等可以帮助你确认该 Service 是否实现了 protobuf 的所有定义。
type service struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

// CreateConsignment 是 service 的一个方法，该方法将 gRPC server 控制的 context 和 request 作为参数
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	repo := s.GetRepo()
	defer repo.Close()

	// 使用 consignment weight 和 容器数量 作为 capacity value 生成一个客户端实例
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	//将从 vessel service 获得的 Id 设为 VesselId
	req.VesselId = vesselResponse.Vessel.Id

	// 保存 consignment
	err = repo.Create(req)
	if err != nil {
		return err
	}

	//返回匹配到我们在 protobuf 里定义的 `Response` message
	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
