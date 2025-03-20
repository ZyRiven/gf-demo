package middleware

import "github.com/gogf/gf/v2/net/ghttp"

func AuthMiddleware(r *ghttp.Request) {
	handler := r.GetServeHandler()

	// 检查是否有noAuth元数据，如果有且值为true，则跳过验证
	noAuth := handler.GetMetaTag("noAuth")
	if noAuth == "true" {
		r.Middleware.Next()
		return
	}

	// 获取当前用户信息（从上下文或JWT、Session等中获取）
	userId := r.GetCtxVar("userId").Int()
	if userId == 0 {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    errcode.Unauthorized.Code(),
			Message: errcode.Unauthorized.Message(),
		})
		return
	}

	// 检查是否需要特定角色
	requiredRole := handler.GetMetaTag("role")
	if requiredRole != "" {
		// 检查用户是否有该角色
		hasRole, err := checkUserRole(userId, requiredRole)
		if err != nil || !hasRole {
			r.Response.WriteJson(ghttp.DefaultHandlerResponse{
				Code:    errcode.Forbidden.Code(),
				Message: errcode.Forbidden.Message(),
			})
			return
		}
	}

	r.Middleware.Next()
}
