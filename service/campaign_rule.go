package service

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type CampaignRule struct {
	Id             int       `orm:"column(id);auto"`
	Cid            string    `orm:"column(cid);size(255)"`
	IncludeOS      string    `orm:"column(include_os);size(255);index"`
	IncludeCountry string    `orm:"column(include_country);size(255);index"`
	IncludeApp     string    `orm:"column(include_app);size(255);index"`
	ExcludeOS      string    `orm:"column(exclude_os);size(255);index"`
	ExcludeCountry string    `orm:"column(exclude_country);size(255);index"`
	ExcludeApp     string    `orm:"column(exclude_app);size(255);index"`
	CreatedAt      time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt      time.Time `orm:"column(update_at);auto_now;type(datetime)"`
}

func (t *CampaignRule) TableName() string {
	return "campaigns_rule"
}

func init() {
	orm.RegisterModel(new(CampaignRule))
}
