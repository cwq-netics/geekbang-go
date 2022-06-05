package data

import (
	"context"

	"module4/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type monitorRepo struct {
	data *Data
	log  *log.Helper
}

// NewMonitorRepo .
func NewMonitorRepo(data *Data, logger log.Logger) biz.MonitorRepo {
	return &monitorRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *monitorRepo) Save(ctx context.Context, g *biz.Monitor) (*biz.Monitor, error) {
	return g, nil
}
