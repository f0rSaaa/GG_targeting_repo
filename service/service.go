package service

import (
	"context"
	"log"

	"github.com/greedy_game/targeting_engine/domain"
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

	app := cast.ToString(ctx.Value("app"))
	country := cast.ToString(ctx.Value("country"))
	os := cast.ToString(ctx.Value("os"))

	campaigns, err := s.Model.GetCampaigns(app, country, os)
	if err != nil {
		return nil, err
	}

	campaignsResp := []domain.CampaignResp{}
	for _, campaign := range campaigns {
		campaignsResp = append(campaignsResp, *campaign.ToCampaignResp())
	}

	return &domain.ResponseModel{
		Code:  200,
		Msg:   "Success",
		Model: campaignsResp,
	}, nil
}
