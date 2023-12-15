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
		apiRouter.GET("/getinfo", ctr.User.Getinfo)
		apiRouter.POST("/admin/show", ctr.Admin.AddShow)
		apiRouter.PUT("/admin/star/:id", ctr.Admin.UpdateStar)
		// apiRouter.PUT("/admin/show/:id", ctr.Admin.UpdateShow)
		apiRouter.GET("/admin/stars", ctr.Admin.GetStar)
		apiRouter.GET("/admin/shows", ctr.Admin.GetShow)
		apiRouter.GET("/user/shows", ctr.User.GetShow)
		apiRouter.POST("/user/grab-ticket", ctr.User.GrabTicket)
		apiRouter.DELETE("/admin/show/:id", ctr.Admin.DeleteShow)
		apiRouter.DELETE("/admin/star/:id", ctr.Admin.DeleteStar)
		// end
	}

}
