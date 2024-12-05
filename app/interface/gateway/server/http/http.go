package http

import (
	"for-share/app/interface/gateway/conf"
	"libong/common/server/http"
	loginConf "libong/login/app/interface/login/conf"
	loginHttp "libong/login/app/interface/login/server/http"
	loginService "libong/login/app/interface/login/service"
)

func New(conf *conf.Config) *http.Server {
	server := http.New(conf.Server.HTTP)

	loginHttp.ConfigHttp(loginService.New(&loginConf.Service{
		//LoginRSAKey: &loginConf.RSAKey{
		//	PublicKeyPath:  conf.Service.LoginRSAKey.LoginPublic,
		//	PrivateKeyPath: conf.Service.LoginRSAKey.LoginPrivate,
		//},
		Dao: &loginConf.Dao{
			Redis: conf.Service.Dao.Redis,
			Mysql: conf.Service.Dao.Mysql,
		},
		RBACService: conf.Service.RBACService,
	}), server)
	return server
}
