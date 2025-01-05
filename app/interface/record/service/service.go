package service

import (
	"for-share/app/interface/record/conf"
	recordServiceConf "for-share/app/service/record/conf"
	recordServiceGrpc "for-share/app/service/record/server/grpc"
	recordServiceService "for-share/app/service/record/service"

	categoryServiceConf "for-share/app/service/category/conf"
	categoryServiceGrpc "for-share/app/service/category/server/grpc"
	categoryServiceService "for-share/app/service/category/service"
)

type Service struct {
	recordService   *recordServiceGrpc.Server
	categoryService *categoryServiceGrpc.Server
}

func New(c *conf.Service) *Service {
	service := &Service{
		recordService: recordServiceGrpc.New(recordServiceService.New(&recordServiceConf.Service{
			Dao: &recordServiceConf.Dao{
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
