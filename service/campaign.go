package service

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Campaign struct {
	Id        int       `orm:"column(id);auto"`
	Cid       string    `orm:"column(cid);size(255);index"`
	Name      string    `orm:"column(cname);size(255);index"`
	Image     string    `orm:"column(image);size(255)"`
	CTA       string    `orm:"column(cta);size(255);unique"`
	Status    string    `orm:"column(status);size(255);index"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(update_at);auto_now;type(datetime)"`
}

func (t *Campaign) TableName() string {
	return "campaigns"
}

func init() {
	orm.RegisterModel(new(Campaign))
}
