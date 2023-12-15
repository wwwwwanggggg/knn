package controller

import (
	"errors"
	"final/common"
	"final/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u *User) Register(c *gin.Context) {
	var info service.UserInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.User.Register(info)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (u *User) Login(c *gin.Context) {
	var info service.UserInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err, success := srv.User.Login(info)
	if err != nil {
		c.Error(common.ErrNew(err, common.AuthErr))
		return
	}
	if success {
		SessionSet(c, "tenzor", UserSession{
			ID:       resp,
			Username: "tenzor"})
		c.JSON(http.StatusOK, ResponseNew(c, resp))
		return
	}
	c.Error(common.ErrNew(err, common.AuthErr))
}

func (u *User) Logout(c *gin.Context) {
	session := SessionGet(c, "tenzor")
	if session == nil {
		c.Error(common.ErrNew(errors.New("未登录"), common.AuthErr))
		return
	}
	SessionClear(c)
	c.JSON(http.StatusOK, ResponseNew(c, nil))
}

func (u *User) Getinfo(c *gin.Context) {
	session, ok := SessionGet(c, "tenzor").(UserSession)
	if session.ID == 0 || !ok {
		c.Error(common.ErrNew(errors.New("未登录"), common.AuthErr))
		return
	}
	resp := struct {
		UserName string
	}{
		UserName: session.Username,
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))

}

func (u *User) GetShow(c *gin.Context) {
	session, ok := SessionGet(c, "tenzor").(UserSession)
	if session.ID == 0 || !ok {
		c.Error(common.ErrNew(errors.New("未登录"), common.AuthErr))
		return
	}
	var form common.PagerForm
	if err := c.ShouldBindQuery(&form); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.User.GetShow(form)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (u *User) GrabTicket(c *gin.Context) {
	session, ok := SessionGet(c, "tenzor").(UserSession)
	if session.ID == 0 || !ok {
		c.Error(common.ErrNew(errors.New("未登录"), common.AuthErr))
		return
	}
	var info struct {
		ShowId uint `json:"showId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&info); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.User.GrabTicket(info.ShowId, uint(session.ID))
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}
