package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"

	"github.com/spf13/viper"
)

// 全局MySQL数据库变量
var DB *gorm.DB

func InitDb() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("db.UserName"),
		viper.GetString("db.Password"),
		viper.GetString("db.Host"),
		viper.GetString("db.Port"),
		viper.GetString("db.Database"),
		viper.GetString("db.Charset"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名单数
		},
	})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Failed to get raw database connection:", err)
		return
	}

	// 数据库连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
}
