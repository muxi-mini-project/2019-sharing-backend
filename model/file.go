package model

import (
	"log"
	"time"
)

type File_uploader struct {
	Fileid     int  `gorm:"file_id"`
	Uploaderid int  `gorm:"uploader_id"`
	Uploadtime string  `gorm:"upload_time"`
}

type File_collecter struct {
	Fileid      int     `gorm:"file_id"`
	Collecterid int     `gorm:"collecter_id"`
	Collecttime string  `gorm:"collect_time"`
}

type File_downloader struct {
	Fileid       int    `gorm:"file_id"`
	Downloaderid int    `gorm:"downloader_id"`
	Downloadtime string `gorm:"download_time"`
}

func CreateNewuploadRecord(fileid int, uploaderid int) bool {
	 var tmprecord File_uploader
	 var t time.Time = time.Now()
	 timestring := t.Format("2006-01-02 15:04")
	 tmprecord.Fileid = fileid
	 tmprecord.Uploaderid = uploaderid
     tmprecord.Uploadtime = timestring
     if err := Db.Self.Model(&File_uploader{}).Create(&tmprecord).Error; err != nil {
     	log.Print("记录新建失败")
     	log.Println(err)
     	return false
	 }
	 return true
}

func CreateNewdownloadRecord(fileid int, downloaderid int) bool {
	var tmprecord File_downloader
	var t time.Time = time.Now()
	timestring := t.Format("2006-01-02 15:04")
	tmprecord.Fileid = fileid
	tmprecord.Downloaderid = downloaderid
	tmprecord.Downloadtime = timestring
	if err := Db.Self.Model(&File_downloader{}).Create(&tmprecord).Error; err != nil {
		log.Print("记录新建失败")
		log.Println(err)
		return false
	}
	return true
}

func CreateNewCollectRecord(fileid int, collecterid int) bool {
	var tmprecord File_collecter
	var t time.Time = time.Now()
	timestring := t.Format("2006-01-02 15:04")
	tmprecord.Fileid = fileid
	tmprecord.Collecterid = collecterid
	tmprecord.Collecttime = timestring
	if err := Db.Self.Model(&File_collecter{}).Create(&tmprecord).Error; err != nil {
		log.Print("记录新建失败")
		log.Println(err)
		return false
	}
	return true
}

func DeleteCollection(fileid int, collecterid int) bool {
	 var tmprecord File_collecter
	 if err := Db.Self.Model(&File_collecter{}).Where(&File_collecter{Fileid:fileid,Collecterid:collecterid}).First(&tmprecord).Error; err != nil {
	 	log.Print("查无此数据")
	 	log.Println(err)
	 	return false
	 }
	if err := Db.Self.Model(&File_collecter{}).Where(&File_collecter{Fileid:fileid,Collecterid:collecterid}).Delete(&File_collecter{}).Error; err != nil {
		log.Print("删除失败")
		log.Println(err)
		return false
	}
	return true
}

func DownloadFile(fileid int) string {
	var tmpfile File
	if err := Db.Self.Model(&File{}).Where(&File{FileId:fileid}).First(&tmpfile).Error; err != nil {
		log.Print("查无此数据")
		log.Println(err)
	}
	return tmpfile.FileUrl
}