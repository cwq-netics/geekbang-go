package service

import (
	"context"

	pb "module4/api/healthz/v1"
	"module4/internal/biz"
)

type MonitorService struct {
	pb.UnimplementedMonitorServer
	uc *biz.MonitorUsecase
}

func NewMonitorService(uc *biz.MonitorUsecase) *MonitorService {
	return &MonitorService{uc: uc}
}

func (s *MonitorService) CheckHealth(ctx context.Context, req *pb.CheckHealthRequest) (*pb.CheckHealthReply, error) {
	g, err := s.uc.CheckHealth(ctx, &biz.Monitor{Hello: req.Message})
	if err != nil {
		return nil, err
	}
	return &pb.CheckHealthReply{Message: "Hellp" + g.Hello}, nil
}
