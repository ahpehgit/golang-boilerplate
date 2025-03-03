package sample

import service "github.com/ahpehgit/golang-boilerplate/service"

type SampleService struct {
	baseService *service.Service
}

func NewSampleService(serviceName string) *SampleService {
	return &SampleService{
		service.NewService(serviceName),
	}
}

func (service *SampleService) GetServiceName() string {
	return service.baseService.GetServiceName()
}
