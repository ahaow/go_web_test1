package initialize

import (
	"go_web_test1/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := global.Config.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to initialize database, got error: %v", err)
	}

	sqlDb, err := db.DB()

	sqlDb.SetMaxIdleConns(global.Config.Database.MaxIdleConns) //最大连接池数量
	sqlDb.SetMaxOpenConns(global.Config.Database.MaxOpenCons)  // 打开数据库最大数量
	sqlDb.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("Failed to configure database, got error: %v", err)
	}

	return db
}
