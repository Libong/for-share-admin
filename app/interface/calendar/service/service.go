package service

import (
	"for-share/app/interface/calendar/conf"
	calendarServiceConf "for-share/app/service/calendar/conf"
	calendarServiceGrpc "for-share/app/service/calendar/server/grpc"
	calendarServiceService "for-share/app/service/calendar/service"
	googleGrpc "google.golang.org/grpc"
	"libong/common/server/grpc"

	categoryServiceConf "for-share/app/service/category/conf"
	categoryServiceGrpc "for-share/app/service/category/server/grpc"
	categoryServiceService "for-share/app/service/category/service"

	accountServiceApi "libong/login/rpc/account/api"
)

type Service struct {
	calendarService *calendarServiceGrpc.Server
	categoryService *categoryServiceGrpc.Server
	accountService  accountServiceApi.AccountServiceClient
}

func New(c *conf.Service) *Service {
	var (
		accountConn *googleGrpc.ClientConn
		err         error
	)

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

	if accountConn, err = grpc.NewConnection(c.RBACService); err != nil {
		panic(-1)
	}
	service.accountService = accountServiceApi.NewAccountServiceClient(accountConn)
	return service
}
