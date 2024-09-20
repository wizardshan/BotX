package controller

import (
	"botx/controller/request"
	"botx/controller/response"
	"botx/pkg/sms"
	"github.com/gin-gonic/gin"
)

type Sms struct {
	smsClient *sms.Client
}

func NewSms(smsClient *sms.Client) *Sms {
	ctr := new(Sms)
	ctr.smsClient = smsClient
	return ctr
}

func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
	req := new(request.SmsCaptcha)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	domainMobile, err := req.Mapper()
	if err != nil {
		return nil, err
	}

	ctr.smsClient.Send(domainMobile.Value)

	return domainMobile, nil
}
