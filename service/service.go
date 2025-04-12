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
}

func NewService(log *log.Logger) Service {
	return &deliveryService{
		Logger: log,
	}
}

func (s *deliveryService) GetDeliveryStatus(ctx context.Context) (*domain.ResponseModel, error) {

	app := cast.ToString(ctx.Value("app"))
	country := cast.ToString(ctx.Value("country"))
	os := cast.ToString(ctx.Value("os"))

	s.Logger.Println("app", app)
	s.Logger.Println("country", country)
	s.Logger.Println("os", os)

	return &domain.ResponseModel{
		Code: 200,
		Msg:  "Success",
		Model: domain.Campaign{
			Id:     "test-campaign",
			Name:   "Test Campaign",
			Image:  "http://example.com/image.jpg",
			CTA:    "Click Here",
			Status: "ACTIVE",
		},
	}, nil
}
