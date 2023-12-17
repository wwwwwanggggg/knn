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
		// end
	}
	r.GET("/oauth2/authorize", ctr.User.OauthLogin)
	r.GET("cb", ctr.User.Callback)
	r.POST("/oauth2/callback", ctr.User.Callback)
	r.POST("/oauth2/authorize", ctr.User.Token)
}

// gto_rpfl6ww3lmmqauqqyqrwhsalnpk3jnyzwd2rjkjejd2wahne452a

// 157256ba-9133-459b-b970-d88044b2f56e
