package model

import (
	"log"
)

type File struct {
	FileId      int     `gorm:"column:file_id" json:"-"`
	FileUrl     string  `gorm:"column:file_url" json:"file_url"`
	FileName    string  `gorm:"column:file_name" json:"file_title"`
	Format      string  `gorm:"column:format" json:"format"`
	Content     string  `gorm:"column:content" json:"content"`
	Subject     string  `gorm:"column:subject" json:"subject"`
	College     string  `gorm:"column:college" json:"college"`
	Type        string  `gorm:"column:type" json:"type"`
	Grade       float64 `gorm:"column:grade" json:"-"`
	LikeNum     int     `gorm:"column:like_num" json:"-"`
	CollcetNum  int     `gorm:"column:collect_num" json:"-"`
	DownloadNum int     `gorm:"column:download_num" json:"-"`
	Scored      int     `gorm:"column:scored" json:"-"`
}

func CreateNewfile(tmpfile File) bool {
	tmpfile.LikeNum = 0
	tmpfile.CollcetNum = 0
	tmpfile.DownloadNum = 0
	tmpfile.Grade = 0
	if err := DB.Self.Model(&File{}).Create(&tmpfile).Error; err != nil {
		log.Print("数据库创建数据失败")
		log.Println(err)
		return false
	}
	return true
}

func Deletefile(fileid int) bool {
	var tmpfile File
	if err := DB.Self.Model(&File{}).Where(&File{FileId: fileid}).First(&tmpfile).Error; err != nil {
		log.Print("查无此数据 ")
		log.Println(err)
		return false
	}
	if err := DB.Self.Model(&File{}).Where(&File{FileId: fileid}).Delete(&File{}).Error; err != nil {
		log.Print("删除失败")
		log.Println(err)
		return false
	}
	return true
}
