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

//GetTransaction will make a request to given endpoint with given headers to a database adapter with a mock database
//and get mock transactions.
func (rt *RestTemplate) GetTransaction() {}

//PostTransaction will bring to you the transaction that you ask it with given ID.
func (rt *RestTemplate) PostTransaction() {}

func GetIRestTemplate() IRestTemplate {

	c := RestTemplate{resty.New()}
	c.SetDebug(false).SetRetryCount(0).SetTimeout(30 * time.Second)
	return &c
}
