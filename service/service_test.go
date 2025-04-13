package service

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/greedy_game/targeting_engine/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockModel is a mock implementation of the Model interface
type MockModel struct {
	mock.Mock
}

func (m *MockModel) GetCampaigns(app, country, os string) ([]domain.Campaign, error) {
	args := m.Called(app, country, os)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Campaign), args.Error(1)
}

func TestGetDeliveryStatus(t *testing.T) {
	tests := []struct {
		name          string
		app           string
		country       string
		os            string
		mockCampaigns []domain.Campaign
		mockError     error
		expectedCode  int
		expectedMsg   string
		expectError   bool
	}{
		{
			name:    "successful_campaign_fetch",
			app:     "testapp",
			country: "US",
			os:      "android",
			mockCampaigns: []domain.Campaign{
				{Id: "camp1", Image: "image1.jpg", CTA: "Click here"},
				{Id: "camp2", Image: "image2.jpg", CTA: "Download now"},
			},
			mockError:    nil,
			expectedCode: 200,
			expectedMsg:  "Success",
			expectError:  false,
		},
		{
			name:          "no_campaigns_found",
			app:           "testapp",
			country:       "UK",
			os:            "ios",
			mockCampaigns: []domain.Campaign{},
			mockError:     nil,
			expectedCode:  200,
			expectedMsg:   "No campaigns found",
			expectError:   false,
		},
		{
			name:          "database_error",
			app:           "testapp",
			country:       "US",
			os:            "android",
			mockCampaigns: nil,
			mockError:     assert.AnError,
			expectedCode:  0,
			expectedMsg:   "",
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock model
			mockModel := new(MockModel)
			mockModel.On("GetCampaigns", tt.app, tt.country, tt.os).Return(tt.mockCampaigns, tt.mockError)

			// Create service with mock
			logger := log.New(os.Stdout, "", log.LstdFlags)
			svc := NewService(logger, mockModel)

			// Create context with test values
			ctx := context.Background()
			ctx = context.WithValue(ctx, "app", tt.app)
			ctx = context.WithValue(ctx, "country", tt.country)
			ctx = context.WithValue(ctx, "os", tt.os)

			// Call service method
			resp, err := svc.GetDeliveryStatus(ctx)

			// Assert results
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectedCode, resp.Code)
				assert.Equal(t, tt.expectedMsg, resp.Msg)

				if tt.mockCampaigns != nil && len(tt.mockCampaigns) > 0 {
					campaigns := resp.Model.([]domain.CampaignResp)
					assert.Equal(t, len(tt.mockCampaigns), len(campaigns))
					for i, camp := range campaigns {
						assert.Equal(t, tt.mockCampaigns[i].Id, camp.Id)
						assert.Equal(t, tt.mockCampaigns[i].Image, camp.Image)
						assert.Equal(t, tt.mockCampaigns[i].CTA, camp.CTA)
					}
				}
			}

			// Verify mock expectations
			mockModel.AssertExpectations(t)
		})
	}
}
