package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Monitor is a Monitor model.
type Monitor struct {
	Hello string
}

// MonitorRepo is a Greater repo.
type MonitorRepo interface {
	Save(context.Context, *Monitor) (*Monitor, error)
}

// MonitorUsecase is a Greeter usecase.
type MonitorUsecase struct {
	repo MonitorRepo
	log  *log.Helper
}

// NewMonitorUsecase new a Greeter usecase.
func NewMonitorUsecase(repo MonitorRepo, logger log.Logger) *MonitorUsecase {
	return &MonitorUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *MonitorUsecase) CheckHealth(ctx context.Context, g *Monitor) (*Monitor, error) {
	uc.log.WithContext(ctx).Infof("CheckHealth: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
