package model

import (
     "log"
)

type File struct {
     FileId   int       `gorm:"file_id" json:"-"`
     FileUrl  string    `gorm:"file_url" json:"file_url"`
     FileName string    `gorm:"file_name" json:"file_title"`
     Format   string    `gorm:"format" json:"format"`
     Content  string    `gorm:"content" json:"content"`
     Subject  string    `gorm:"subject" json:"subject"`
     College  string    `gorm:"college" json:"college"`
     Type     string    `gorm:"type" json:"type"`
     Grade    float32   `gorm:"grade" json:"-"`
     Likes    int       `gorm:"like_num" json:"-"`
     CollcetNumber  int `gorm:"collect_num" json:"-"`
     DownloadNumber int `gorm:"download_num" json:"-"`
}

func CreateNewfile(tmpfile File) bool {
     tmpfile.Likes = 0
     tmpfile.CollcetNumber = 0
     tmpfile.DownloadNumber = 0
     tmpfile.Grade = 0
     if err := Db.Self.Model(&File{}).Create(&tmpfile).Error; err != nil {
          log.Print("数据库创建数据失败")
          log.Println(err)
          return false
     }
     return true
}

func Deletefile(fileid int) bool {
     var tmpfile File
     if err := Db.Self.Model(&File{}).Where(&File{FileId:fileid}).First(&tmpfile).Error; err != nil {
          log.Print("查无此数据 ")
          log.Println(err)
          return false
     }
     if err := Db.Self.Model(&File{}).Where(&File{FileId:fileid}).Delete(&File{}).Error; err != nil {
          log.Print("删除失败")
          log.Println(err)
          return false
     }
     return true
}

