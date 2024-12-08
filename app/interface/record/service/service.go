package service

import "for-share/app/interface/record/conf"

type Service struct {
}

func New(c *conf.Service) *Service {
	service := &Service{}
	return service
}
