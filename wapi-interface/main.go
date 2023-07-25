package main

import (
	"wapi-interface/config"
	"wapi-interface/routes"
)

func main() {
	err := routes.Router.Run(config.Conf.Gin.GetAddr())
	if err != nil {
		return
	}
}
