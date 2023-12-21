package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"final/common"
	"final/service"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u *User) Register(c *gin.Context) {
	var info service.UserInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	_, err := srv.User.Register(info)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, nil))
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

func (u *User) GetTicketInfo(c *gin.Context) {
	session, ok := SessionGet(c, "tenzor").(UserSession)
	if session.ID == 0 || !ok {
		c.Error(common.ErrNew(errors.New("未登录"), common.AuthErr))
		return
	}
	resp, err := srv.User.GetTicketInfo(session.ID)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (u *User) AbandonTicket(c *gin.Context) {
	session, ok := SessionGet(c, "tenzor").(UserSession)
	if session.ID == 0 || !ok {
		c.Error(common.ErrNew(errors.New("未登录"), common.AuthErr))
		return
	}
	var id struct {
		Id int `form:"id" binding:"required"`
	}
	if err := c.ShouldBindQuery(&id); err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	resp, err := srv.User.AbandonTicket(id.Id, session.ID)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (u *User) OauthLogin(c *gin.Context) {
	url := "https://git.tiaozhan.tech/oauth/authorize?client_id=78d198cc-9cc0-49a3-bfab-63f195123895&redirect_uri=http://127.0.0.1:8080/oauth2/callback&response_type=code&state=STATE"
	fmt.Println(url)
	c.Redirect(http.StatusFound, url)

}

func (u *User) Callback(c *gin.Context) {
	var info struct {
		Code  string `form:"code" binding:"required"`
		State string `form:"state" binding:"required"`
	}
	if err := c.ShouldBindQuery(&info); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	uri := "https://git.tiaozhan.tech/login/oauth/access_token"
	body := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {info.Code},
		"client_id":     {"78d198cc-9cc0-49a3-bfab-63f195123895"},
		"client_serect": {"gto_wuui7wurtx6trr2nqzxn4ubps6dvl55a27w74syzh3tf63q3s6aq"},
		"redirect_url":  {"http://127.0.0.1:8080/oauth2/callback"},
	}
	req, err := http.NewRequest("POST", uri, bytes.NewBufferString(body.Encode()))
	if err != nil {
		c.Error(common.ErrNew(err, common.AuthErr))
		return
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("78d198cc-9cc0-49a3-bfab-63f195123895", "gto_wuui7wurtx6trr2nqzxn4ubps6dvl55a27w74syzh3tf63q3s6aq")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Error(common.ErrNew(err, common.AuthErr))
		return
	}
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Error(common.ErrNew(err, common.AuthErr))
		return
	}
	var Access_token struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(respbody, &Access_token); err != nil {
		c.Error(common.ErrNew(err, common.AuthErr))
		return
	}
	accessToken := Access_token.AccessToken
	userApi := "https://git.tiaozhan.tech/api/v1/user"
	getUser, err := http.NewRequest("GET", userApi, nil)
	if err != nil {
		c.Error(common.ErrNew(err, common.AuthErr))
	}
	getUser.Header.Set("Authorization", "Bearer "+accessToken)
	getInfo := &http.Client{}
	r, e := getInfo.Do(getUser)
	if e != nil {
		c.Error(common.ErrNew(e, common.AuthErr))
		return
	}
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		c.Error(common.ErrNew(err, common.AuthErr))
		return
	}
	var userinfo struct {
		UserName string `json:"username"`
	}

	if err := json.Unmarshal(rBody, &userinfo); err != nil {
		c.Error(common.ErrNew(err, common.AuthErr))
		return
	}
	fmt.Println(userinfo)
	id, err := srv.User.OauthLogin(userinfo.UserName)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	SessionSet(c, "tenzor", UserSession{
		Username: userinfo.UserName,
		ID:       id,
	})
	c.JSON(http.StatusOK, ResponseNew(c, nil))
}
