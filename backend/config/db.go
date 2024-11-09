package config

import (
	"demo_hubu_backend/middleware"
	"demo_hubu_backend/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var err error

func DbInit() (*gorm.DB, error) {
	// 初始化 Viper 并读取配置文件
	viper.SetConfigName("config") // 配置文件名（不带扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 配置文件路径
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 解析配置文件
	var config struct {
		Database struct {
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			Name     string `mapstructure:"name"`
		}
		JWT struct {
			Secret string `mapstructure:"secret"`
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// 构建 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	// 初始化数据库连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移模式
	ModelsAutoMigrate()
	middleware.SetJWTSecret(config.JWT.Secret)
	return db, err
}

func ModelsAutoMigrate() {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.ApprovalResource{})
}
