package service

import (
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"
)

type MockOrm struct {
	orm.Ormer
	mockQueryRows func(interface{}) (int64, error)
}

func (m *MockOrm) Raw(query string, args ...interface{}) orm.RawSeter {
	return &MockRawSeter{mockQueryRows: m.mockQueryRows}
}

type MockRawSeter struct {
	orm.RawSeter
	mockQueryRows func(interface{}) (int64, error)
}

func (m *MockRawSeter) QueryRows(container ...interface{}) (int64, error) {
	if len(container) > 0 {
		return m.mockQueryRows(container[0])
	}
	return 0, nil
}

func TestGetCampaigns(t *testing.T) {
	tests := []struct {
		name          string
		app           string
		country       string
		os            string
		mockRules     []*CampaignRule
		mockCampaigns []*Campaign
		expectedIds   []string
		expectError   bool
	}{
		{
			name:    "matching_campaigns",
			app:     "testapp",
			country: "US",
			os:      "android",
			mockRules: []*CampaignRule{
				{
					Cid:            "camp1",
					IncludeApp:     "testapp",
					IncludeCountry: "US",
					IncludeOS:      "android",
				},
			},
			mockCampaigns: []*Campaign{
				{
					Cid:    "camp1",
					Image:  "image1.jpg",
					CTA:    "Click here",
					Status: "ACTIVE",
				},
			},
			expectedIds: []string{"camp1"},
			expectError: false,
		},
		{
			name:    "excluded_campaign",
			app:     "testapp",
			country: "US",
			os:      "android",
			mockRules: []*CampaignRule{
				{
					Cid:            "camp1",
					IncludeApp:     "testapp",
					IncludeCountry: "US",
					IncludeOS:      "android",
					ExcludeApp:     "testapp",
				},
			},
			mockCampaigns: nil,
			expectedIds:   nil,
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock orm
			mockOrm := &MockOrm{
				mockQueryRows: func(container interface{}) (int64, error) {
					switch v := container.(type) {
					case *[]*CampaignRule:
						*v = tt.mockRules
						return int64(len(tt.mockRules)), nil
					case *[]*Campaign:
						*v = tt.mockCampaigns
						return int64(len(tt.mockCampaigns)), nil
					}
					return 0, nil
				},
			}

			// Create model with mock orm
			model := NewDatabaseModel(mockOrm)

			// Call model method
			campaigns, err := model.GetCampaigns(tt.app, tt.country, tt.os)

			// Assert results
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.expectedIds != nil {
					assert.Equal(t, len(tt.expectedIds), len(campaigns))
					for i, camp := range campaigns {
						assert.Equal(t, tt.expectedIds[i], camp.Id)
					}
				} else {
					assert.Empty(t, campaigns)
				}
			}
		})
	}
}

func TestExcludeRule(t *testing.T) {
	tests := []struct {
		name          string
		rule          CampaignRule
		app           string
		country       string
		os            string
		expectedMatch bool
	}{
		{
			name: "no_exclusions",
			rule: CampaignRule{
				Cid:            "camp1",
				IncludeApp:     "testapp",
				IncludeCountry: "US",
				IncludeOS:      "android",
			},
			app:           "testapp",
			country:       "US",
			os:            "android",
			expectedMatch: true,
		},
		{
			name: "excluded_app",
			rule: CampaignRule{
				Cid:            "camp1",
				IncludeApp:     "testapp",
				IncludeCountry: "US",
				IncludeOS:      "android",
				ExcludeApp:     "testapp",
			},
			app:           "testapp",
			country:       "US",
			os:            "android",
			expectedMatch: false,
		},
		{
			name: "excluded_country",
			rule: CampaignRule{
				Cid:            "camp1",
				IncludeApp:     "testapp",
				IncludeCountry: "US",
				IncludeOS:      "android",
				ExcludeCountry: "US",
			},
			app:           "testapp",
			country:       "US",
			os:            "android",
			expectedMatch: false,
		},
		{
			name: "excluded_os",
			rule: CampaignRule{
				Cid:            "camp1",
				IncludeApp:     "testapp",
				IncludeCountry: "US",
				IncludeOS:      "android",
				ExcludeOS:      "android",
			},
			app:           "testapp",
			country:       "US",
			os:            "android",
			expectedMatch: false,
		},
		{
			name: "multiple_exclusions",
			rule: CampaignRule{
				Cid:            "camp1",
				IncludeApp:     "testapp",
				IncludeCountry: "US",
				IncludeOS:      "android",
				ExcludeApp:     "testapp,otherapp",
				ExcludeCountry: "US,UK",
				ExcludeOS:      "android,ios",
			},
			app:           "testapp",
			country:       "US",
			os:            "android",
			expectedMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := &model{}
			result := !model.ExcludeRule(tt.rule, tt.app, tt.country, tt.os)
			assert.Equal(t, tt.expectedMatch, result)
		})
	}
}
