package model

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type File_uploader struct {
	Id         int    `gorm : "id"`
	UploaderId string `gorm : "uploader_id"`
	FileId     int    `gorm : "file_id"`
	Uploadtime string `gorm : "upload_time"`
}

type File_downloader struct {
	Id           int    `gorm : "id"`
	DownloaderId string `gorm : "downloader_id"`
	FileId       int    `gorm : "file_id"`
	Downloadtime string `gorm : "download_time"`
}

type File_collecter struct {
	Id          int    `gorm : "id"`
	CollecterId string `gorm : "collecter_id"`
	FileId      int    `gorm : "file_id"`
	Collecttime string `gorm : "collect_time"`
}

type Likes struct {
	id     int    `gorm : "id"`
	UserId string `gorm : "user_id"`
	FileId int    `gorm : "file_id"`
}

func CreateNewDownloadRecord(fileid int, downloaderid string) bool {
	var tmprecord File_downloader
	tmprecord.FileId = fileid
	tmprecord.DownloaderId = downloaderid
	tNow := time.Now()
	timeNow := tNow.Format("2006-01-02 15:04:05")
	tmprecord.Downloadtime = timeNow
	if err := DB.Self.Model(&File_downloader{}).Create(&tmprecord).Error; err != nil {
		log.Println(err)
		log.Print("记录创建失败")
		return false
	}
	return true
}

func CreateNewCollectRecord(fileid int, collecterid string) bool {
	var tmprecord File_collecter
	tmprecord.FileId = fileid
	tmprecord.CollecterId = collecterid
	tNow := time.Now()
	timeNow := tNow.Format("2006-01-02 15:04:05")
	tmprecord.Collecttime = timeNow
	if err := DB.Self.Model(&File_downloader{}).Create(&tmprecord).Error; err != nil {
		log.Println(err)
		log.Print("记录创建失败")
		return false
	}
	return true
}

func CreateNewUploadRecord(fileid int, uploaderid string) bool {
	var tmprecord File_uploader
	tmprecord.FileId = fileid
	tmprecord.UploaderId = uploaderid
	tNow := time.Now()
	timeNow := tNow.Format("2006-01-02 15:04:05")
	tmprecord.Uploadtime = timeNow
	if err := DB.Self.Model(&File_downloader{}).Create(&tmprecord).Error; err != nil {
		log.Println(err)
		log.Print("记录创建失败")
		return false
	}
	return true
}

func Like(fileid int, userid string) bool {
	var tmprecord Likes
	tmprecord.FileId = fileid
	tmprecord.UserId = userid
	if err := DB.Self.Model(&Likes{}).Create(&tmprecord).Error; err != nil {
		log.Println(err)
		log.Print("记录创建失败")
		return false
	}
	return true
}