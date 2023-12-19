package router

import (
	"final/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Error)

	apiRouter := r.Group("/api")
	{
		// example
		// begin
		apiRouter.POST("/user", ctr.User.Register)
		apiRouter.POST("/admin/star", ctr.Admin.AddStar)
		apiRouter.POST("/login", ctr.User.Login)
		apiRouter.DELETE("/logout", ctr.User.Logout)
		apiRouter.GET("/user/info", ctr.User.Getinfo)
		apiRouter.POST("/admin/show", ctr.Admin.AddShow)
		apiRouter.PUT("/admin/star/", ctr.Admin.UpdateStar)
		apiRouter.PUT("/admin/show/", ctr.Admin.UpdateShow)
		apiRouter.GET("/admin/stars", ctr.Admin.GetStar)
		apiRouter.GET("/admin/shows", ctr.Admin.GetShow)
		apiRouter.GET("/user/shows", ctr.User.GetShow)
		apiRouter.POST("/user/tickets", ctr.User.GrabTicket)
		apiRouter.DELETE("/admin/show/:id", ctr.Admin.DeleteShow)
		apiRouter.DELETE("/admin/star/:id", ctr.Admin.DeleteStar)
		apiRouter.GET("/user/tickets", ctr.User.GetTicketInfo)
		apiRouter.DELETE("/user/adandon-ticket", ctr.User.AbandonTicket)
		apiRouter.GET("/admin/users", ctr.Admin.GetUser)
		apiRouter.GET("/test", ctr.Hello.Hello)
		// end
	}
	r.GET("/oauth2/authorize", ctr.User.OauthLogin)
	r.GET("/oauth2/callback", ctr.User.Callback)

}

// 78d198cc-9cc0-49a3-bfab-63f195123895

// gto_wuui7wurtx6trr2nqzxn4ubps6dvl55a27w74syzh3tf63q3s6aq
