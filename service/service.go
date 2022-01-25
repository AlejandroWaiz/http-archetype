package service

import "github.com/AlejandroWaiz/http-archetype/resttemplate"

type Service struct {
	Rest resttemplate.IRestTemplate
}

type IService interface {
	GetTransactionService()
	PostTransactionService()
}

func (s *Service) GetTransactionService()  {}
func (s *Service) PostTransactionService() {}

func GetIService(rest resttemplate.IRestTemplate) IService {

	return &Service{rest}

}
