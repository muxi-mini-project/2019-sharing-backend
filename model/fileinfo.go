package model

import (
     "log"
)

type File struct {
     FileId   int       `gorm:"file_id"`
     FileUrl  string    `gorm:"file_url"`
     FileName string    `gorm:"file_name"`
     Format   string    `gorm:"format"`
     Content  string    `gorm:"content"`
     Subject  string    `gorm:"subject"`
     College  string    `gorm:"college"`
     Type     string    `gorm:"type"`
     Grade    float32   `gorm:"grade"`
     Likes    int       `gorm:"like_num"`
     CollcetNumber  int `gorm:"collect_num"`
     DownloadNumber int `gorm:"download_num"`
}

func CreateNewfile(uploader string,filename string, fileurl string,format string, content string, subject string ,college string, typename string) bool {
     var tmpfile File
     tmpfile.FileName = filename
     tmpfile.FileUrl = fileurl
     tmpfile.Likes = 0
     tmpfile.CollcetNumber = 0
     tmpfile.DownloadNumber = 0
     tmpfile.Content = content
     tmpfile.College = college
     tmpfile.Format = format
     tmpfile.Subject = subject
     tmpfile.Grade = 0
     tmpfile.Type = typename
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

