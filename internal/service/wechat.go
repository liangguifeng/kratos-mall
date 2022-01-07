package service

import (
	"context"

	pb "kratos-mall/api/pay/v1"
)

type WechatService struct {
	pb.UnimplementedWechatServer
}

func NewWechatService() *WechatService {
	return &WechatService{}
}

func (s *WechatService) H5(ctx context.Context, req *pb.H5Request) (*pb.H5Response, error) {
	return &pb.H5Response{}, nil
}
