package transport

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/greedy_game/targeting_engine/domain"
	"github.com/greedy_game/targeting_engine/endpoint"
	"github.com/greedy_game/targeting_engine/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPHandler returns an HTTP handler for the delivery service
func NewHTTPHandler(svc service.Service) http.Handler {
	r := mux.NewRouter()

	// Create endpoint
	getDeliveryEndpoint := endpoint.MakeGetDeliveryStatusEndpoint(svc)

	// Create v1 subrouter
	v1 := r.PathPrefix("/v1").Subrouter()

	// GET /v1/delivery
	v1.Methods("GET").Path("/delivery").Handler(kithttp.NewServer(
		getDeliveryEndpoint,
		decodeGetDeliveryRequest,
		encodeResponse,
	))

	// Add metrics endpoint
	r.Handle("/metrics", promhttp.Handler())

	return r
}

func decodeGetDeliveryRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return domain.DeliveryRequest{
		App:     r.URL.Query().Get("app"),
		Country: r.URL.Query().Get("country"),
		OS:      r.URL.Query().Get("os"),
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
