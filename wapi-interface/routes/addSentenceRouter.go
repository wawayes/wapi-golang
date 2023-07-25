package routes

import "wapi-interface/api"

func addSentenceRouter() {
	apiGroup.GET("/sentence", api.RootApi.GetSentence)
}
