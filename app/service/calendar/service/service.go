package service

import (
	"for-share/app/service/calendar/conf"
	"for-share/app/service/calendar/dao"
)

type Service struct {
	dao *dao.Dao
}

func New(c *conf.Service) *Service {
	return &Service{
		dao: dao.New(),
	}
}
