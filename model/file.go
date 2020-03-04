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
	var tmpfile File
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
	if err := DB.Self.Model(&File{}).Where(&File{FileId: fileid}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		return false
	}
	tmpfile.DownloadNum++
	if err := DB.Self.Model(&File{}).Where(&File{FileId: fileid}).Update("download_num", tmpfile.DownloadNum).Error; err != nil {
		log.Println(err)
		log.Print("点赞统计失败")
		return false
	}
	return true
}

func CreateNewCollectRecord(fileid int, collecterid string, collectlistid int) bool {
	var tmprecord1 File_collecter
	var tmpfile File
	var tmprecord2 File_collecter
	//对tmprecord1赋值，进行新建操作
	tmprecord1.FileId = fileid
	tmprecord1.CollecterId = collecterid
	tNow := time.Now()
	timeNow := tNow.Format("2006-01-02 15:04:05")
	tmprecord1.Collecttime = timeNow
	tmprecord1.CollectlistId = collectlistid
	//利用tmprecord2进行一个记录是否存在的检测
	if err := DB.Self.Model(&File_collecter{}).Where(&File_collecter{FileId:fileid,CollectlistId:collectlistid,CollecterId:collecterid}).Find(&tmprecord2).Error; tmprecord2.Collecttime != ""{
		log.Println(err)
		log.Print("已收藏")
		return false
	}
	if err := DB.Self.Model(&File_collecter{}).Create(&tmprecord1).Error; err != nil {
		log.Println(err)
		log.Print("记录创建失败")
		return false
	}
	if err := DB.Self.Model(&File{}).Where(&File{FileId: fileid}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		return false
	}
	tmpfile.CollcetNum++
	if err := DB.Self.Model(&File{}).Where(&File{FileId: fileid}).Update("collect_num", tmpfile.CollcetNum).Error; err != nil {
		log.Println(err)
		log.Print("点赞统计失败")
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
