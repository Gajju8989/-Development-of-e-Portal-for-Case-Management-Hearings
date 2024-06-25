package router

import "Pc_Build_Service/Internal/userinterface/handler"

type Router interface {
	MapRoutes()
}
type impl struct {
	userInterfaceHandler handler.UserInterfaceHandler
}

func NewRouter(userInterfaceHandler handler.UserInterfaceHandler) Router {
	return &impl{
		userInterfaceHandler: userInterfaceHandler,
	}
}
