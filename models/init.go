// +build !sqlite

package model

import (
	"fmt"
	"github.com/HFO4/cloudreve/pkg/conf"
	"github.com/HFO4/cloudreve/pkg/util"
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// Init 初始化 MySQL 链接
func Init() {
	util.Log().Info("初始化数据库连接")

	var (
		db  *gorm.DB
		err error
	)

	db, err = gorm.Open(conf.DatabaseConfig.Type, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DatabaseConfig.User,
		conf.DatabaseConfig.Password,
		conf.DatabaseConfig.Host,
		conf.DatabaseConfig.Name))

	// 处理表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.DatabaseConfig.TablePrefix + defaultTableName
	}

	// Debug模式下，输出所有 SQL 日志
	if conf.SystemConfig.Debug {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}

	//db.SetLogger(util.Log())
	if err != nil {
		util.Log().Panic("连接数据库不成功, %s", err)
	}

	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	//执行迁移
	migration()
}
