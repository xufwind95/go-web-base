package database

import (
	"fmt"

	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 下划线为仅仅调用该包的Init函数

	"github.com/xufwind95/go-web-base/config"
)

var (
	DB *gorm.DB
)

// 连接数据库
func openDB(username, password, addr, name string, logMode bool) {
	config_ := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local",
	)

	log.Info(config_)

	db, err := gorm.Open("mysql", config_)
	if err != nil {
		panic(err)
	}
	DB = db
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.LogMode(logMode)
}

func InitDB(conf *config.Config) {
	// 打开数据库连接
	openDB(
		conf.DataBase.UserName,
		conf.DataBase.PassWord,
		conf.DataBase.Addr,
		conf.DataBase.DbName,
		conf.Gormlog,
	)
}

func CloseDB() {
	DB.Close()
}
