package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"sleep-service/internal/model"
	"sleep-service/internal/service"
)

// AuthMiddleware 路由鉴权
func AuthMiddleware(r *ghttp.Request) {
	handler := r.GetServeHandler()
	// 检查是否有noAuth元数据，如果有且值为true，则跳过验证
	noAuth := handler.GetMetaTag("noAuth")
	if noAuth == "true" {
		r.Middleware.Next()
		return
	}

	// 获取当前用户信息
	ctx := r.GetCtx()
	r.SetCtx(r.GetNeverDoneCtx())
	data, err := service.GfToken().ParseToken(r)
	if err != nil || data == nil {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    401,
			Message: "无效token",
		})
		return
	} else {
		contextModel := new(model.Context)
		err = gconv.Struct(data.Data, &contextModel.User)
		if err != nil {
			g.Log().Error(ctx, err)
			r.Response.WriteJson(ghttp.DefaultHandlerResponse{
				Code:    400,
				Message: "数据转化格式错误",
			})
			return
		}
		service.Context().Init(r, contextModel)
	}

	// 检查是否需要特定角色
	//requiredRole := handler.GetMetaTag("role")
	//if requiredRole != "" {
	//	// 检查用户是否有该角色
	//	hasRole, err := checkUserRole(userId, requiredRole)
	//	if err != nil || !hasRole {
	//		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
	//			Code:    403,
	//			Message: "无权限访问",
	//		})
	//		return
	//	}
	//}

	r.Middleware.Next()
}
