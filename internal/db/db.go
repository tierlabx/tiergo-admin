package db

import (
	"database/sql"
	"fmt"
	"tier-up/internal/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(c config.Config) (*sql.DB, *gorm.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", c.DB.Host, c.DB.User, c.DB.Password, c.DB.DriverName, c.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print("\n", "db", db)
		fmt.Print("\n", "-----------------")
		fmt.Print("\n", "err", err)
		fmt.Print("\n", "链接数据库失败")
		panic(err)
	}
	// 迁移表
	if c.DB.AutoCreateTable {
		AutoMigrate(db)
	}

	// 链接池
	sqldb, err := db.DB()
	if err != nil {
		panic("连接池启动出错")
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqldb.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqldb.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqldb.SetConnMaxLifetime(time.Hour)

	if err = sqldb.Ping(); err != nil {
		panic(err)
	}

	return sqldb, db
}
