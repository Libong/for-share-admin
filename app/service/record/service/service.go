package service

import (
	"for-share/app/service/record/conf"
	"for-share/app/service/record/dao"
)

type Service struct {
	dao *dao.Dao
}

func New(c *conf.Service) *Service {
	return &Service{
		dao: dao.New(),
	}
}
