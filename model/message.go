package model

import (
	"log"
	"time"
)
type Message struct {
	WriterId string  `gorm:"writer_id"`
	HostId   string  `gorm:"host_id"`
	WriteTime string `gorm:"write_time"`
	Content string   `gorm:"content"`
}

func CreateNewMessage(writerid string, hostid string, content string) bool {
	var tmpnote Message
	tmpnote.WriterId = writerid
	tmpnote.HostId = hostid
	tmpnote.Content = content
	tNow := time.Now()
	timeNow := tNow.Format("2006-01-02 15:04:05")
	tmpnote.WriteTime = timeNow
	if err := DB.Self.Model(&Message{}).Create(&tmpnote).Error; err != nil {
		log.Println(err)
		log.Print("留言失败")
		return false
	}
	return true
}