package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/greedy_game/targeting_engine/domain"
	"github.com/greedy_game/targeting_engine/service"
)

func MakeGetDeliveryStatusEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.DeliveryRequest)

		// Validate request
		if req.App == "" || req.Country == "" || req.OS == "" {
			return domain.DeliveryResponse{
				Code: 400,
				Msg:  "missing app/country/os param",
			}, nil
		}

		//pass all the values to the service using context
		ctx = context.WithValue(ctx, "app", req.App)
		ctx = context.WithValue(ctx, "country", req.Country)
		ctx = context.WithValue(ctx, "os", req.OS)

		resp, err := svc.GetDeliveryStatus(ctx)
		if err != nil {
			return domain.DeliveryResponse{
				Code: 500,
				Msg:  "Internal server error, Please try again later",
				Err:  err.Error(),
			}, nil
		}

		return domain.DeliveryResponse{
			Code:  resp.Code,
			Msg:   resp.Msg,
			Model: resp.Model,
		}, nil
	}
}
