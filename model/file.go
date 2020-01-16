package model

import (
	"log"
	"time"
)

type file_uploader struct {
	fileid     int
	uploaderid int
	uploadtime string
}

type file_collecter struct {
	fileid     int
	collecterid int
	collecttime string
}

type file_downloader struct {
	fileid     int
	downloaderid int
	downloadtime string
}

func CreateNewuploadRecord(fileid int, uploaderid int) bool {
	 var tmprecord file_uploader
	 var t time.Time = time.Now()
	 timestring := t.Format("2006-01-02 15:04")
	 tmprecord.fileid = fileid
	 tmprecord.uploaderid = uploaderid
     tmprecord.uploadtime = timestring
     if err := Db.Self.Model(&file_uploader{}).Create(&tmprecord).Error; err != nil {
     	log.Print("记录新建失败")
     	log.Println(err)
     	return false
	 }
	 return true
}

func CreateNewdownloadRecord(fileid int, downloaderid int) bool {
	var tmprecord file_downloader
	var t time.Time = time.Now()
	timestring := t.Format("2006-01-02 15:04")
	tmprecord.fileid = fileid
	tmprecord.downloaderid = downloaderid
	tmprecord.downloadtime = timestring
	if err := Db.Self.Model(&file_downloader{}).Create(&tmprecord).Error; err != nil {
		log.Print("记录新建失败")
		log.Println(err)
		return false
	}
	return true
}

func CreateNewCollectRecord(fileid int, collecterid int) bool {
	var tmprecord file_collecter
	var t time.Time = time.Now()
	timestring := t.Format("2006-01-02 15:04")
	tmprecord.fileid = fileid
	tmprecord.collecterid = collecterid
	tmprecord.collecttime = timestring
	if err := Db.Self.Model(&file_collecter{}).Create(&tmprecord).Error; err != nil {
		log.Print("记录新建失败")
		log.Println(err)
		return false
	}
	return true
}