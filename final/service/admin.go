package service

import (
	"errors"
	"final/common"
	"final/model"
	"fmt"
	"time"
)

type StarInfo struct {
	Name  string `json:"name" binding:"required"`
	Intro string `json:"intro" binding:"required"`
}

type ShowInfo struct {
	Title       string    `json:"title" binding:"required"`
	Starttime   time.Time `json:"starttime" time_fotmat:"2006-01-02 15:04:05" binding:"required"`
	Endtime     time.Time `json:"endtime" time_format:"2006-01-02 15:04:05" binding:"required"`
	Location    string    `json:"location"`
	MaxCapacity int       `json:"maxcapacity" binding:"required"`
	ShowContent string    `json:"showcontent"`
	Promo       string    `json:"promo"`
	Performers  []int     `json:"performers" binding:"required"`
}
type GetStarInfo struct {
	Id    uint   `json:"StarId"`
	Name  string `json:"name"`
	Intro string `json:"intro"`
}

type UpdateShowInfo struct {
	ShowId      uint      `json:"id" binding:"required"`
	Title       string    `json:"title"`
	Starttime   time.Time `json:"starttime" time_format:"2006-01-02 15:04:05"`
	Endtime     time.Time `json:"endtime" time_format:"2006-01-02 15:04:05"`
	Showcontent string    `json:"showcontent"`
	Promo       string    `json:"promo"`
	Performers  []uint    `json:"performers"`
	MaxCapacity int       `json:"maxcapacity"`
	Location    string    `json:"location"`
}

type GetShowInfo struct {
	Id        uint      `json:"ShowId"`
	Title     string    `json:"title"`
	Starttime time.Time `json:"startTime" time_format:"2006-01-02 15:04:05"`
	Endtime   time.Time `json:"endTime" time_format:"2006-01-02 15:04:05"`
	Location  string    `json:"location"`
}

type Uinfo struct {
	Id   int    `json:"userId"`
	Name string `json:"userName"`
}

type GetUserInfo struct {
	Total int     `json:"total"`
	Users []Uinfo `json:"users"`
}

type Admin struct {
}

func (a *Admin) AddStar(info StarInfo) (interface{}, error) {
	s := model.Star{
		Name:  info.Name,
		Intro: info.Intro,
	}
	if err := model.DB.Model(&model.Star{}).
		Where("name = ?", info.Name).
		FirstOrCreate(&s).
		Error; err != nil {
		return nil, err
	}
	return struct {
		StarId uint
	}{
		StarId: s.ID,
	}, nil
}

func (a *Admin) AddShow(show ShowInfo) (interface{}, error) {
	s := model.Show{
		Title: show.Title, MaxCapacity: show.MaxCapacity,
		CurrCapacity: show.MaxCapacity, Sold: 0,
		ShowContent: show.ShowContent, Location: show.Location,
		Promo:     show.Promo,
		Starttime: show.Starttime, Endtime: show.Endtime,
	}
	var stars []model.Star
	for _, va := range show.Performers {
		var tem model.Star
		if err := model.DB.Model(&model.Star{}).
			Where("id = ?", va).
			First(&tem).
			Error; err != nil {
			return 0, errors.New("明星不存在")
		}
		stars = append(stars, tem)
	}
	s.Performers = stars
	if err := model.DB.Model(&model.Show{}).Create(&s).Error; err != nil {
		return 0, err
	}
	return struct {
		ShowId int `json:"showId"`
	}{
		ShowId: int(s.ID),
	}, nil
}

func (a *Admin) UpdateStar(name string, intro string, id int) (interface{}, error) {
	var tem model.Star
	if err := model.DB.Model(&model.Star{}).
		Where("id = ?", id).
		First(&tem).Error; err != nil {
		return 0, errors.New("用户不存在")
	}
	if name != "" {
		err := model.DB.Model(&model.Star{}).Where("id = ?", id).
			Update("name", name).Error
		if err != nil {
			return 0, errors.New("更改失败")
		}
	}
	if intro != "" {
		err := model.DB.Model(&model.Star{}).Where("id = ?", id).
			Update("intro", intro).Error
		if err != nil {
			return 0, errors.New("更改失败")
		}
	}
	return struct {
		StarId int `json:"starId"`
	}{
		StarId: id,
	}, nil
}

func (u *Admin) UpdateShow(show UpdateShowInfo, Id int) (interface{}, error) {
	var s model.Show
	if err := model.DB.Model(&model.Show{}).
		Where("id = ?", Id).
		First(&s).Error; err != nil {
		return 0, errors.New("演出不存在")
	}
	if show.Title != "" {
		if err := model.DB.Model(&s).Update("title", show.Title).Error; err != nil {
			return 0, errors.New("修改失败")
		}
	}
	if show.Promo != "" {
		if err := model.DB.Model(&s).Update("promo", show.Promo).Error; err != nil {
			return 0, errors.New("修改失败")
		}
	}
	if show.Showcontent != "" {
		if err := model.DB.Model(&s).Update("show_content", show.Showcontent).Error; err != nil {
			return 0, errors.New("修改失败")
		}
	}
	if show.Location != "" {
		if err := model.DB.Model(&s).Update("location", show.Location).Error; err != nil {
			return 0, errors.New("修改失败")
		}
	}
	if show.MaxCapacity != 0 && show.MaxCapacity > s.Sold {
		if err := model.DB.Model(&s).
			Update("max_capacity", show.MaxCapacity).
			Update("curr_capacity", show.MaxCapacity-s.Sold).
			Error; err != nil {
			return 0, errors.New("修改失败")
		}
	}
	if show.Performers != nil {
		var tem []struct {
			ShowId uint `column:"show_id"`
			StarId uint `column:"star_id"`
		}
		if err := model.DB.Table("show_star").
			Where("show_id = ?", Id).
			Delete(&tem).Error; err != nil {
			return 0, errors.New("修改失败")
		}
		for _, i := range show.Performers {
			var t struct {
				ShowId uint `column:"show_id"`
				StarId uint `column:"star_id"`
			}
			t.ShowId = s.ID
			t.StarId = i
			model.DB.Table("show_star").Create(&t)
		}
	}

	return struct {
		Message string
	}{
		Message: "演出详情已修改",
	}, nil
}

func (a *Admin) GetStar(form common.PagerForm) (interface{}, error) {
	var info []GetStarInfo
	if err := model.DB.Model(&model.Star{}).Select("id", "name", "intro").
		Offset((form.Page - 1) * form.Limit).
		Limit(form.Limit).Find(&info).Error; err != nil {
		return 0, errors.New("查询失败")
	}
	return info, nil
}

func (a *Admin) GetShow(form common.PagerForm) (interface{}, error) {
	var info []GetShowInfo
	if err := model.DB.Model(&model.Show{}).
		Select("id", "title", "starttime", "location", "endtime").
		Offset((form.Page - 1) * form.Limit).
		Limit(form.Limit).Find(&info).Error; err != nil {
		return 0, errors.New("查询失败")
	}
	return info, nil
}

func (a *Admin) DeleteShow(Id int) (interface{}, error) {
	var show model.Show
	if err := model.DB.Model(&model.Show{}).Preload("Performers").Preload("Users").
		Where("id = ?", Id).
		Find(&show).Error; err != nil {
		return 0, errors.New("演出不存在")
	}
	fmt.Println(show)
	if len(show.Performers) > 0 {
		if err := model.DB.Model(&show).
			Association("Performers").
			Clear().Error; err != nil {
			return 0, errors.New("删除失败")
		}
	}
	if len(show.Users) > 0 {
		if err := model.DB.Model(&model.Show{}).Where("id = ?", Id).
			Association("Users").
			Clear().Error; err != nil {
			return 0, errors.New("删除失败")
		}
	}
	if err := model.DB.Model(&model.Show{}).Delete(&show).Error; err != nil {
		return 0, errors.New("删除失败")
	}
	return struct {
		Message string
	}{
		Message: "演出已成功删除",
	}, nil
}

func (a *Admin) DeleteStar(Id int) (interface{}, error) {
	var star model.Star
	if err := model.DB.Model(&model.Star{}).
		Where("id = ?", Id).Preload("Show").
		Find(&star).Error; err != nil {
		return 0, errors.New("明星不存在")
	}
	if len(star.Show) > 0 {
		return struct {
			Shows   []model.Show
			Message string
		}{
			Shows:   star.Show,
			Message: "明星参加了这些演出,请先删除这些演出",
		}, nil
	}
	if err := model.DB.Model(&model.Star{}).Delete(&star).Error; err != nil {
		return 0, errors.New("删除失败")
	}
	return struct {
		Message string
	}{
		Message: "删除成功",
	}, nil

}

func (a *Admin) GetUser(form common.PagerForm) (interface{}, error) {
	var infos []Uinfo
	if err := model.DB.Model(&model.User{}).
		Select("id", "name").
		Offset((form.Page - 1) * form.Limit).
		Limit(form.Limit).Find(&infos).Error; err != nil {
		return 0, errors.New("查询失败")
	}
	return GetUserInfo{
		Total: len(infos),
		Users: infos,
	}, nil

}
