package model

import (
	"log"
	"strconv"
	"time"
)

type File_uploader struct {
	UploaderId string `gorm:"column:uploader_id"`
	FileId     int    `gorm:"column:file_id"`
	Uploadtime string `gorm:"column:upload_time"`
}

type File_downloader struct {
	DownloaderId string `gorm:"column:downloader_id"`
	FileId       int    `gorm:"column:file_id"`
	Downloadtime string `gorm:"column:download_time"`
}

type File_collecter struct {
	CollecterId   string `gorm:"column:collecter_id"`
	FileId        int    `gorm:"column:file_id"`
	Collecttime   string `gorm:"column:collect_time"`
	CollectlistId int    `gorm:"column:collectlist_id"`
}

type Likes struct {
	UserId string `gorm:"column:user_id"`
	FileId int    `gorm:"column:file_id"`
}

type Score struct {
	Score  int    `gorm:"column:score"`
	Userid string `gorm:"column:user_id"`
	Fileid int    `gorm:"column:file_id"`
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

func CreateNewCollectRecord(fileid int, collecterid string, collectlistid int) bool {
	var tmprecord File_collecter
	tmprecord.FileId = fileid
	tmprecord.CollecterId = collecterid
	tNow := time.Now()
	timeNow := tNow.Format("2006-01-02 15:04:05")
	tmprecord.Collecttime = timeNow
	tmprecord.CollectlistId = collectlistid
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
	var tmpfile File
	tmprecord.FileId = fileid
	tmprecord.UserId = userid
	if err := DB.Self.Model(&Likes{}).Create(&tmprecord).Error; err != nil {
		log.Println(err)
		log.Print("记录创建失败")
		return false
	}
	if err := DB.Self.Model(&File{}).Where(&File{FileId: fileid}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		return false
	}
	tmpfile.LikeNum++
	if err := DB.Self.Model(&File{}).Where(&File{FileId: fileid}).Update("like_num", tmpfile.LikeNum).Error; err != nil {
		log.Println(err)
		log.Print("点赞统计失败")
		return false
	}
	return true
}

func Unlike(fileid int, userid string) bool {
	var tmprecord Likes
	var tmpfile File
	tmprecord.FileId = fileid
	tmprecord.UserId = userid
	if err := DB.Self.Where(&Likes{FileId: fileid, UserId: userid}).Delete(&Likes{}).Error; err != nil {
		log.Println(err)
		log.Print("记录创建失败")
		return false
	}
	if err := DB.Self.Model(&File{}).Where(&File{FileId: fileid}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		return false
	}
	tmpfile.LikeNum--
	if err := DB.Self.Model(&File{}).Where(&File{FileId: fileid}).Update("like_num", tmpfile.LikeNum).Error; err != nil {
		log.Println(err)
		log.Print("点赞统计失败")
		return false
	}
	return true
}

func CreateScoreRecord(userid string, fileid int, score int) bool {
	var tmpscore Score
	tmpscore.Score = score
	tmpscore.Fileid = fileid
	tmpscore.Userid = userid
	if err := DB.Self.Model(&Score{}).Create(&tmpscore).Error; err != nil {
		log.Println(err)
		log.Print("评分失败")
		return false
	}
	return true
}

func InttoFloat(a int) float64 {
	i := strconv.Itoa(a)
	float, err := strconv.ParseFloat(i, 64)
	if err != nil {
		log.Println(err)
	}
	return float
}
