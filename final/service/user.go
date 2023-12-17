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

type General struct {
	Message string `json:"message"`
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
		return General{Message: "票已售完"}, nil
	}
	if err := model.DB.Model(&show).
		Update("curr_capacity", show.CurrCapacity-1).
		Error; err != nil {
		return 0, errors.New("系统有误")
	}
	if err := model.DB.Model(&show).Update("sold", show.Sold+1).Error; err != nil {
		return 0, errors.New("系统有误")
	}
	info := struct {
		ShowId int `column:"show_id"`
		UserId int `column:"user_id"`
	}{
		ShowId: int(Id),
		UserId: int(userId),
	}
	if err := model.DB.Table("show_user").Create(&info).Error; err != nil {
		return 0, errors.New("抢票失败")
	}
	return struct {
		TicketId uint
		ShowId   uint
	}{
		TicketId: uint(show.Sold),
		ShowId:   show.ID,
	}, nil

}

func (u *User) GetTicketInfo(Id int) (interface{}, error) {
	var shows model.User
	if err := model.DB.Model(&model.User{}).
		Where("id = ?", Id).Preload("Show").
		Find(&shows).Error; err != nil {
		return 0, errors.New("用户不存在")
	}
	return struct {
		Total int          `json:"total"`
		Shows []model.Show `json:"show"`
	}{
		Total: len(shows.Show),
		Shows: shows.Show,
	}, nil
}

func (u *User) AbandonTicket(tId int, uId int) (interface{}, error) {
	var user model.User
	if err := model.DB.Model(&model.User{}).
		Where("id = ?", uId).
		Preload("Show").
		Find(&user).Error; err != nil {
		return 0, errors.New("用户不存在")
	}
	if len(user.Show) > 0 {
		var info struct {
			StarId uint
			UserId uint
		}
		if err := model.DB.Table("show_user").
			Where("show_id = ? AND user_id = ?", tId, uId).
			Find(&info).Error; err != nil {
			return 0, errors.New("演出不存在")
		}
		if err := model.DB.Table("show_user").
			Where("show_id = ? AND user_id = ?", tId, uId).
			Delete(&info).Error; err != nil {
			return 0, errors.New("删除失败")
		}
		var show model.Show
		if err := model.DB.Model(&model.Show{}).
			Where("id = ?", tId).Find(&show).Error; err != nil {
			return 0, errors.New("删除失败2")
		}
		if err := model.DB.Model(&show).
			Update("sold", show.Sold-1).
			Update("curr_capacity", show.CurrCapacity+1).Error; err != nil {
			return 0, errors.New("删除失败3")
		}
	}
	return General{
		Message: "票已成功放弃",
	}, nil
}
