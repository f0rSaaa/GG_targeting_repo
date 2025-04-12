package service

import (
	"github.com/astaxie/beego/orm"
	"github.com/greedy_game/targeting_engine/domain"
)

type Model interface {
	GetCampaigns(app, country, os string) ([]domain.Campaign, error)
}

type model struct {
	db *orm.Ormer
}

func NewDatabseModel(db *orm.Ormer) Model {
	return &model{
		db: db,
	}
}

func (m *model) GetCampaigns(app string, country string, os string) ([]domain.Campaign, error) {
	return nil, nil
}
