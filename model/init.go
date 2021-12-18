package model

import (
	"fmt"
	"go-restful-api-server-simple-practice/config"
	"strings"

	"github.com/lexkong/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Datebase struct {
	Self *gorm.DB
}

var DB *Datebase

func (db *Datebase) Init() {
	DB = &Datebase{
		Self: getSelfDb(),
	}
}

// func (db *Datebase) Close() {
// 	db.Self.Close()
// }

func getSelfDb() *gorm.DB {
	return InitSelfDb()
}

func InitSelfDb() *gorm.DB {
	return openDb(&config.C.ViperConfig.Mysql)
}

func openDb(c *config.MySqlConfig) *gorm.DB {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=%s", c.Username, c.Password, strings.Join([]string{c.Host, c.Port}, ":"), c.Database, "Local")
	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", c.Username, c.Password, strings.Join([]string{c.Host, c.Port}, ":"), c.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(err, "Database connection failed, Dabases info: [%s] \n", dsn)
	}
	log.Infof("Database connection success, Dabases info: [%s] \n", dsn)
	return db
}
