package service

import (
	"context"
	"fmt"

	pb "smalls/api"
	"smalls/internal/dao"
	"github.com/bilibili/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.AsmallsServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// SayHello grpc smalls func.
func (s *Service) SayHello(ctx context.Context, req *pb.HelloReq) (reply *empty.Empty, err error) {
	reply = new(empty.Empty)
	fmt.Printf("hello %s", req.Name)
	return
}

// SayHelloURL bm smalls func.
func (s *Service) SayHelloURL(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {
	reply = &pb.HelloResp{
		Content: "hello " + req.Name,
	}
	fmt.Printf("hello url %s", req.Name)
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}

func (s *Service) Create(ctx context.Context, req *pb.Req) (reply *pb.Resp, err error) {
	reply = &pb.Resp{
		Content: "Create " + req.Name + reverseString(req.Name),
	}
	fmt.Printf("Create %s", req.Name)
	return
}

func (s *Service) Delete(ctx context.Context, req *pb.Req) (reply *pb.Resp, err error) {
	reply = &pb.Resp{
		Content: "Delete " + req.Name,
	}
	fmt.Printf("Delete %s", req.Name)
	return
}

//	fmt.Println(reverseString(s))
//	fmt.Println(reverseString(reverseString(s)))

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func (s *Service) Get(ctx context.Context, req *pb.Req) (reply *pb.Resp, err error) {
	reply = &pb.Resp{
		Content: "Get " + req.Name + reverseString(req.Name) + "is ok!",
	}
	fmt.Printf("Get %s", req.Name)
	return
}