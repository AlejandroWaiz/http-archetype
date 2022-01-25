package resttemplate

import (
	"time"

	resty "github.com/go-resty/resty/v2"
)

var (
	urls map[string]string = make(map[string]string)
)

type RestTemplate struct {
	*resty.Client
}

type IRestTemplate interface {
	GetTransaction(headers map[string]string, transactionID string) (*resty.Response, error)
	PostTransaction(headers map[string]string, body interface{}) (*resty.Response, error)
}

//GetTransaction will make a request to given endpoint with given headers to a database adapter with a mock database
//and get mock transactions.
func (rt *RestTemplate) GetTransaction(headers map[string]string, transactionID string) (*resty.Response, error) {
	return rt.R().
		SetHeaders(headers).
		Get(urls["http-adapter-url"] + transactionID)
}

//PostTransaction will bring to you the transaction that you ask it with given ID.
func (rt *RestTemplate) PostTransaction(headers map[string]string, body interface{}) (*resty.Response, error) {
	return rt.R().
		SetHeaders(headers).
		SetBody(body).
		Post(urls["http-adapter-url"])
}

func GetIRestTemplate() IRestTemplate {

	c := RestTemplate{resty.New()}
	c.SetDebug(false).SetRetryCount(0).SetTimeout(30 * time.Second)
	return &c
}
