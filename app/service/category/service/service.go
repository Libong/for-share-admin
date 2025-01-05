package service

import (
	"for-share/app/service/category/conf"
	"for-share/app/service/category/dao"
)

type Service struct {
	dao *dao.Dao
}

func New(c *conf.Service) *Service {
	return &Service{
		dao: dao.New(),
	}
}
