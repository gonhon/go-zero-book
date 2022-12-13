// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	class "github.com/gonhon/go-zero-book/service/user/api/internal/handler/class"
	"github.com/gonhon/go-zero-book/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: loginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/class/add",
				Handler: class.AddClassHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/class/update",
				Handler: class.UpdateClassHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/class/delete/:id",
				Handler: class.DeleteClassHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/class/findClass",
				Handler: class.FindClassHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/class/:id",
				Handler: class.GetByIdHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}
