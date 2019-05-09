// consignment-service/repository.go
package main

import (
	pb "github.com/pengxianghu/shipper/consignment-service/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type ConsignmentRepository struct {
	session *mgo.Session
}

// 创建一个新的 consignment（委托）
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

// 获取所有的consignments
func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {

	var consignments []*pb.Consignment

	//Find()通常需要一个参数，但如果我们想要返回所有的结果就将参数设为 nil
	//然后将所有的 consignments 作为参数传递给.All（）函数，
	//.All（）函数将所有的 consignments 作为查询的结果返回
	//在这里还可以调用 One（）方法来返回一个单一的consignment

	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}

// Close() 负责在每个查询运行结束后关闭数据库session。
// Mgo 在启动时创建一个主 session ，主 session 会为每个请求创建一个新的 session。
// 这意味着每个请求都有自己的数据库 session，这样的机制会使得会话更安全、高效。
// 更底层来讲，每个 session 中都有自己独立的数据库 socket 和错误处理机制。
// 使用一个主数据库 socket 意味着其余请求必须等待主 session 优先使用 cpu 资源。
// I.e方法使得我们拒绝锁机制而允许多个请求同时处理。这一点很棒！但是...这意味着我们必须确保每个 session 在完成时关闭掉，于此同时你可能会建立大量的连接，以至于达到连接限制。这一点尤其需要注意！！
func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(consignmentCollection)
}
