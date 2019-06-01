// user-service/handler.go
package main

import (
	"errors"
	pb "github.com/pengxianghu/shipper/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
	micro "github.com/micro/go-micro"
)

const topic = "user.created"

type service struct {
	repo         Repository
	tokenService Authable
	Publisher micro.Publisher
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	var reqP string = req.Password
	log.Printf("input user: %+v", req)
	user, err := srv.repo.GetByEmailAndPassword(req)
	log.Printf("gotuser: %+v", user)
	if err != nil {
		return err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqP)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	// Generates a hashed version of our password
	log.Println("--------in handler--------")
	log.Printf("req user1: %+v", req)
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPass)
	if err := srv.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	log.Printf("req user2: %+v", req)
	
	if err := srv.Publisher.Publish(ctx, req); err != nil {
        return err
    }
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {

	// Decode token
	claims, err := srv.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	log.Println(claims)

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true

	return nil
}