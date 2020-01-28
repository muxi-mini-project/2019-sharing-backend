package model

import (
	"time"
	"log"
)

type Message struct {
	 Writerid int     `gorm:"writer_id"`
	 Hostid   int     `gorm:"host_id"`
	 Writetime string `gorm:"write_time"`
}

func LeaveMessage(writerid int, hostid int) bool {
	var tmpmessage Message
	var t time.Time = time.Now()
	timestring := t.Format("2006-01-02 15:04")
	tmpmessage.Writerid = writerid
	tmpmessage.Hostid = hostid
	tmpmessage.Writetime = timestring
	if err := Db.Self.Model(&Message{}).Create(&tmpmessage).Error; err != nil {
		log.Print("记录新建失败")
		log.Println(err)
		return false
	}
	return true
}
