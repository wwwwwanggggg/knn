package controller

import (
	"final/common"
	"final/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Admin struct {
}

func (a *Admin) AddStar(c *gin.Context) {
	var info service.StarInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.Admin.AddStar(info)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (a *Admin) AddShow(c *gin.Context) {
	var show service.ShowInfo
	if err := c.ShouldBindJSON(&show); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.Admin.AddShow(show)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (a *Admin) UpdateStar(c *gin.Context) {
	var info struct {
		StarId int    `json:"id" binding:"required"`
		Name   string `json:"name"`
		Intro  string `json:"intro"`
	}
	if err := c.ShouldBindJSON(&info); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.Admin.UpdateStar(info.Name, info.Intro, info.StarId)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))

}

func (a *Admin) UpdateShow(c *gin.Context) {
	var show service.UpdateShowInfo
	if err := c.ShouldBindJSON(&show); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.Admin.UpdateShow(show, int(show.ShowId))
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (a *Admin) GetStar(c *gin.Context) {
	var form common.PagerForm
	if err := c.ShouldBindQuery(&form); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.Admin.GetStar(form)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (u *Admin) GetShow(c *gin.Context) {
	var form common.PagerForm
	if err := c.ShouldBindQuery(&form); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.Admin.GetShow(form)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (a *Admin) DeleteShow(c *gin.Context) {
	Id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.Admin.DeleteShow(Id)
	if err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (a *Admin) DeleteStar(c *gin.Context) {
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.Admin.DeleteStar(Id)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (a *Admin) GetUser(c *gin.Context) {
	var form common.PagerForm
	if err := c.ShouldBindQuery(&form); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp, err := srv.Admin.GetUser(form)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}
