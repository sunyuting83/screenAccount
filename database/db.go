package database

import (
	"AccountSell/utils"
	"database/sql"
	"os"
	"runtime"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Eloquent *sql.DB
	sqlDB    *gorm.DB
)

// InitDB init db
func InitDB() {
	CurrentPath, _ := utils.GetCurrentPath()
	OS := runtime.GOOS
	linkString := "/"
	if OS == "windows" {
		linkString = "\\"
	}
	dbPath := strings.Join([]string{CurrentPath, "db"}, linkString)
	if !utils.IsExist(dbPath) {
		os.MkdirAll(dbPath, 0755)
	}
	dbFile := strings.Join([]string{dbPath, "data.db"}, linkString)
	sqlDB, _ = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})

	Eloquent, _ = sqlDB.DB()
	Eloquent.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	Eloquent.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	Eloquent.SetConnMaxLifetime(time.Hour)
	sqlDB.AutoMigrate(&User{}, &Accounts{})
}
