package routers

func (api *Api) useUser() {
	router := api.router.Group("user")
	{
		router.GET("/demo", api.User.Demo)
	}
}
