//数据库的连接

package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

const dns = "root:@tcp(localhost:3306)/muxi_sharing"

type Database struct {
	Self *gorm.DB
}

var DB *Database

func getDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Print("getDatabase")
		log.Println(err)
	}
	db.SingularTable(true)
	return db, err
}

func (db *Database) Init() {
	newDb, err := getDatabase()
	if err != nil {

	}
	DB = &Database{Self: newDb}
}

func (db *Database) Close() {
	if err := DB.Self.Close(); err != nil {

	}
	return
}
