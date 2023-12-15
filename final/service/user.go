package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"final/common"
	"final/model"
)

type UserInfo struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Convert(s string) string {
	after := md5.Sum([]byte(s))
	res := hex.EncodeToString(after[:])
	return res
}

type User struct {
}

func (u *User) Register(info UserInfo) (int, error) {
	user := &model.User{

		Name:     info.Name,
		Password: Convert(info.Password),
	}

	if err := model.DB.Model(&model.User{}).
		Where("name = ?", info.Name).
		FirstOrCreate(&user).
		Error; err != nil {
		return 0, err
	}
	return int(user.ID), nil
}

func (u *User) Login(info UserInfo) (int, error, bool) {
	user := model.User{
		Name:     info.Name,
		Password: Convert(info.Password),
	}
	var tem model.User
	if err := model.DB.Model(&model.User{}).
		Where("name = ?", user.Name).
		Find(&tem).Error; err != nil {
		return 0, errors.New("用户名或密码错误"), false
	}
	if tem.Password == Convert(info.Password) {
		return int(tem.ID), nil, true
	}
	return 0, nil, false

}

func (u *User) GetShow(form common.PagerForm) (interface{}, error) {
	var shows []model.Show
	if err := model.DB.Model(&model.Show{}).
		Select("id", "title", "starttime", "endtime", "location", "max_capacity", "curr_capacity").
		Preload("Performers").
		Find(&shows).Error; err != nil {
		return 0, errors.New("查询错误")
	}
	return shows, nil
}

func (u *User) GrabTicket(Id uint, userId uint) (interface{}, error) {
	var show model.Show
	if err := model.DB.Model(&model.Show{}).
		Where("id = ?", Id).
		First(&show).Error; err != nil {
		return 0, errors.New("演出不存在")
	}
	if show.CurrCapacity == 0 {
		return 0, errors.New("票已售完")
	}
	if err := model.DB.Model(&show).
		Update("curr_capacity", show.CurrCapacity-1).
		Error; err != nil {
		return 0, errors.New("系统有误")
	}
	if err := model.DB.Model(&show).Update("sold", show.Sold+1).Error; err != nil {
		return 0, errors.New("系统有误")
	}
	var user model.User
	if err := model.DB.Model(&model.User{}).
		Where("id = ?", userId).
		Find(&user).Error; err != nil {
		return 0, errors.New("这是用户查找")
	}
	if err := model.DB.Model(&model.Show{}).Where("id = ?", show.ID).Association("Users").
		Append(&user).Error; err != nil {
		return 0, errors.New("系统有误")
	}
	return struct {
		TicketId uint
		ShowId   uint
	}{
		TicketId: uint(show.Sold),
		ShowId:   show.ID,
	}, nil

}
