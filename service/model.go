package service

import (
	"slices"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/greedy_game/targeting_engine/domain"
)

type Model interface {
	GetCampaigns(app, country, os string) ([]domain.Campaign, error)
}

type model struct {
	db orm.Ormer
}

func NewDatabaseModel(db orm.Ormer) Model {
	return &model{
		db: db,
	}
}

func (m *model) GetCampaigns(app, country, os string) ([]domain.Campaign, error) {
	var campaignRules []*CampaignRule
	query := `select * from campaigns_rule where FIND_IN_SET('` + app + `', include_app) > 0 and FIND_IN_SET('` + country + `', include_country) > 0 and FIND_IN_SET('` + os + `', include_os) > 0`
	_, err := m.db.Raw(query).QueryRows(&campaignRules)
	if err != nil {
		return nil, err
	}

	//valid campaigns
	var validCampaigns []string
	// Check if any rule matches
	for _, rule := range campaignRules {
		if !m.ExcludeRule(*rule, app, country, os) {
			validCampaigns = append(validCampaigns, rule.Cid)
		}
	}

	// Return empty if no valid campaigns
	if len(validCampaigns) == 0 {
		return []domain.Campaign{}, nil
	}

	campString := "'" + strings.Join(validCampaigns, "','") + "'"
	query = `select * from campaigns where cid in (` + campString + `) and status = 'ACTIVE' order by id desc`

	var campaigns []*Campaign
	_, err = m.db.Raw(query).QueryRows(&campaigns)
	if err != nil {
		return nil, err
	}

	campaignsDomain := []domain.Campaign{}
	for _, campaign := range campaigns {
		campaignsDomain = append(campaignsDomain, domain.Campaign{
			Id:    campaign.Cid,
			Image: campaign.Image,
			CTA:   campaign.CTA,
		})
	}
	return campaignsDomain, nil
}

func (m *model) ExcludeRule(rule CampaignRule, app, country, os string) bool {
	// Check excludes
	if rule.ExcludeApp != "" {
		excludeApps := strings.Split(rule.ExcludeApp, ",")
		if slices.Contains(excludeApps, app) {
			return true
		}
	}
	if rule.ExcludeCountry != "" {
		excludeCountries := strings.Split(rule.ExcludeCountry, ",")
		if slices.Contains(excludeCountries, country) {
			return true
		}
	}
	if rule.ExcludeOS != "" {
		excludeOS := strings.Split(rule.ExcludeOS, ",")
		if slices.Contains(excludeOS, os) {
			return true
		}
	}
	return false
}
