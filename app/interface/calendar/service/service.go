package service

import (
	"for-share/app/interface/calendar/conf"
	calendarServiceConf "for-share/app/service/calendar/conf"
	calendarServiceGrpc "for-share/app/service/calendar/server/grpc"
	calendarServiceService "for-share/app/service/calendar/service"

	categoryServiceConf "for-share/app/service/category/conf"
	categoryServiceGrpc "for-share/app/service/category/server/grpc"
	categoryServiceService "for-share/app/service/category/service"
)

type Service struct {
	calendarService *calendarServiceGrpc.Server
	categoryService *categoryServiceGrpc.Server
}

func New(c *conf.Service) *Service {
	service := &Service{
		calendarService: calendarServiceGrpc.New(calendarServiceService.New(&calendarServiceConf.Service{
			Dao: &calendarServiceConf.Dao{
				Mysql: c.Dao.Mysql,
			},
		})),
		categoryService: categoryServiceGrpc.New(categoryServiceService.New(&categoryServiceConf.Service{
			Dao: &categoryServiceConf.Dao{
				Mysql: c.Dao.Mysql,
			},
		})),
	}
	return service
}
