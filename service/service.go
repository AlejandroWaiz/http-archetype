package service

import "github.com/AlejandroWaiz/http-archetype/resttemplate"

type Service struct {
	Rest resttemplate.IRestTemplate
}

type IService interface {
	GetTransactionService(header map[string]string, transactionID string) (int, []byte, error)
	PostTransactionService(headers map[string]string, body interface{}) (int, []byte, error)
}

func (s *Service) GetTransactionService(header map[string]string, transactionID string) (int, []byte, error) {
	response, err := s.Rest.GetTransaction(header, transactionID)
	return response.StatusCode(), response.Body(), err
}
func (s *Service) PostTransactionService(headers map[string]string, body interface{}) (int, []byte, error) {

	response, err := s.Rest.PostTransaction(headers, body)

	return response.StatusCode(), response.Body(), err

}

func GetIService(rest resttemplate.IRestTemplate) IService {

	return &Service{rest}

}
