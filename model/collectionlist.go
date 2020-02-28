package model

import log "github.com/sirupsen/logrus"

type Collect_list struct {
	CollectlistId   int    `gorm:"column:collectlist_id"`
	CollectlistName string `gorm:"column:collectlist_name"`
	UserID          string `gorm:"column:user_id"`
}

func CreateNewcollectlist(collect_name string, userid string) bool {
	var tmpcollectlist Collect_list
	tmpcollectlist.CollectlistName = collect_name
	tmpcollectlist.UserID = userid
	if err := DB.Self.Model(&Collect_list{}).Create(&tmpcollectlist).Error; err != nil {
		log.Print("新建收藏夹失败")
		log.Println(err)
		return false
	}
	return true
}
