package domain

type ResponseModel struct {
	Code  int
	Msg   string
	Model interface{}
}

type Campaign struct {
	Id     string
	Name   string
	Image  string
	CTA    string
	Status string
}

type CampaignResp struct {
	Id    string `json:"cid"`
	Image string `json:"img"`
	CTA   string `json:"cta"`
}

type DeliveryRequest struct {
	App     string `json:"app"`
	Country string `json:"country"`
	OS      string `json:"os"`
}

type DeliveryResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"message"`
	Model interface{} `json:"model,omitempty"`
	Err   string      `json:"error,omitempty"`
}
