package model

import (
     "github.com/stretchr/testify/assert"
     "log"
      "time"
)

type File struct {
     FileId int
     FileUrl string
     FileName string
     Format string
     Content string
     Subject string
     College string
     Type string
     Grade float32
     Likes int
     CollcetNumber int
     DownloadNumber int
     Status bool
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
          return assert.False()
     }
     return true
}

