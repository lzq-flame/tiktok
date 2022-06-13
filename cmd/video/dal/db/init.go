package db

import (
	"example/pkg/constants"
	"example/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logdb "gorm.io/gorm/logger"
	gormopentracing "gorm.io/plugin/opentracing"
	"log"
	"time"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 16:15
 **/

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	newLogger := logdb.New(
		log.New(logger.L().Writer(), "", log.LstdFlags),
		logdb.Config{
			SlowThreshold: time.Second,
			LogLevel:      logdb.Info,
			Colorful:      true,
		},
	)
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 newLogger,
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&Video{}) {
		return
	}
	if err = m.CreateTable(&Video{}); err != nil {
		panic(err)
	}
}
