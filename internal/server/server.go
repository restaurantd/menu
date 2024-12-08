package server

import (
	"context"
	"menu/internal/data"
	"menu/internal/loger"
	"menu/internal/proto/pb"
)

type Server struct {
	pb.UnimplementedBookingServer
	DB     data.DB
	Logger loger.Loger
}

func (s Server) GetList(_ context.Context, req *pb.GetListReq) (*pb.RespListPage, error) {

}

func (s Server) GetProduct(_ context.Context, req *pb.GetProductReq) (*pb.Product, error) {

}
