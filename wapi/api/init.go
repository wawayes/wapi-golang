package api

import (
	"wapi/api/interfaceApi"
	"wapi/api/userApi"
)

type Api struct {
	UserApi      userApi.UserApi
	InterfaceApi interfaceApi.InterfaceApi
}

var RootApi = new(Api)
