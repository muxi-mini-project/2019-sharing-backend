package model
 
import (
        "fmt"
       _"github.com/go-sql-driver/mysql"
        "github.com/jinzhu/gorm"
	"log"
)

const dns = "root:password@/muxi_sharing?charset=utf8&parseTime=True&loc=Localo"

type Database struct {
     Self *gorm.DB
}

var Db *Database

func getDatabase() (*gorm.DB, error) {
     db, err := gorm.Open("mysql", dns)
     if err != nil {
	fmt.Print("getDatabase")
	log.Println(err)
     }
     return db, err
}

func (db *Database) Init() {
     newDb, err := getDatabase()
     if err != nil {
	log.Print("数据库初始化错误")
	log.Println(err)
     }
     Db = &Database{Self: newDb}
     return
}

func (db *Database) Close() {
      if err := Db.Self.Close(); err != nil {
	 log.Print("数据库关闭错误")
	 log.Println(err)
      }
      return
}
