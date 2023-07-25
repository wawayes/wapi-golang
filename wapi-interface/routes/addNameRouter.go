package routes

import (
	"wapi-interface/api"
	"wapi-interface/middleware"
)

func addNameRouter() {
	nameRouter := apiGroup.Group("user")
	{
		nameRouter.GET("name", middleware.ValidHeader, api.RootApi.GetUserName)
		nameRouter.GET("get", api.RootApi.GetName)
		nameRouter.POST("post", api.RootApi.PostName)
	}
}
