package service

import (
	"context"
	"log"

	"github.com/greedy_game/targeting_engine/domain"
	"github.com/greedy_game/targeting_engine/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
)

type Service interface {
	GetDeliveryStatus(ctx context.Context) (*domain.ResponseModel, error)
}

type deliveryService struct {
	Logger *log.Logger
	Model  Model
}

func NewService(log *log.Logger, model Model) Service {
	return &deliveryService{
		Logger: log,
		Model:  model,
	}
}

func (s *deliveryService) GetDeliveryStatus(ctx context.Context) (*domain.ResponseModel, error) {
	timer := prometheus.NewTimer(metrics.RequestDuration.WithLabelValues("/v1/delivery"))
	defer timer.ObserveDuration()

	app := cast.ToString(ctx.Value("app"))
	country := cast.ToString(ctx.Value("country"))
	os := cast.ToString(ctx.Value("os"))

	campaigns, err := s.Model.GetCampaigns(app, country, os)
	if err != nil {
		metrics.RequestTotal.WithLabelValues("error").Inc()
		return nil, err
	}

	if len(campaigns) == 0 {
		metrics.RequestTotal.WithLabelValues("no_campaigns").Inc()
		metrics.CampaignsReturned.Observe(0)
		return &domain.ResponseModel{
			Code:  200,
			Msg:   "No campaigns found",
			Model: []domain.CampaignResp{},
		}, nil
	}

	campaignsResp := []domain.CampaignResp{}
	for _, campaign := range campaigns {
		campaignsResp = append(campaignsResp, *campaign.ToCampaignResp())
	}

	metrics.RequestTotal.WithLabelValues("success").Inc()
	metrics.CampaignsReturned.Observe(float64(len(campaignsResp)))

	return &domain.ResponseModel{
		Code:  200,
		Msg:   "Success",
		Model: campaignsResp,
	}, nil
}
