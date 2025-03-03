package service

type Service struct {
	serviceName string
}

func NewService(serviceName string) *Service {
	return &Service{
		serviceName,
	}
}

func (s *Service) GetServiceName() string {
	return s.serviceName
}
