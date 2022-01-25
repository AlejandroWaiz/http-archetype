package resttemplate

import (
	"time"

	resty "github.com/go-resty/resty/v2"
)

type RestTemplate struct {
	*resty.Client
}

type IRestTemplate interface {
	GetTransaction()
	PostTransaction()
}

func (rt *RestTemplate) GetTransaction()  {}
func (rt *RestTemplate) PostTransaction() {}

func GetIRestTemplate() IRestTemplate {

	c := RestTemplate{resty.New()}
	c.SetDebug(false).SetRetryCount(0).SetTimeout(30 * time.Second)
	return &c
}
