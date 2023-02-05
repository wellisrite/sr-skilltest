package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"sr-skilltest/internal/domain"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Properties struct {
		App      App      `json:"app"`
		Database Database `json:"database"`
		Service  Service  `json:"service"`
		Cache    Cache    `json:"cache"`
	}

	App struct {
		Mode            string `json:"mode"`
		Debug           string `json:"debug"`
		CoachPortal     string `json:"coach_portal"`
		BaseURL         string `json:"base_url"`
		DjangoFCM       string `json:"django_fcm"`
		CorsAllowOrigin string `json:"cors_allow_origin"`
	}

	Database struct {
		DBHost     string `json:"db_host"`
		DBName     string `json:"db_name"`
		DBUser     string `json:"db_user"`
		DBPassword string `json:"db_password"`
		DBPort     string `json:"db_port"`
	}

	Cache struct {
		CacheHost     string `json:"cache_host"`
		CachePassword string `json:"cache_password"`
		CacheDB       int    `json:"cache_db"`
		CachePort     string `json:"cache_port"`
	}

	Service struct {
		ServicePort string `json:"service_port"`
		TimeZone    string `json:"time_zone"`
		PoolSize    int    `json:"pool_size"`
		LogPath     string `json:"log_path"`
	}
)

func getProperties() (properties Properties) {
	isLocalDev := flag.Bool("local", false, "=(true/false)")
	flag.Parse()
	if *isLocalDev {
		return loadPropertiesLocal()
	} else {
		return loadEnv()
	}
}

func loadPropertiesLocal() Properties {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	properties := loadEnv()
	return properties
}

func loadEnv() Properties {
	timestart := time.Now()
	fmt.Println("Starting Load Config " + timestart.Format("2006-01-02 15:04:05"))

	properties := Properties{
		App: App{
			Mode:            os.Getenv("APP_MODE"),
			Debug:           os.Getenv("APP_DEBUG_MODE"),
			BaseURL:         os.Getenv("APP_BASE_URL"),
			CorsAllowOrigin: os.Getenv("CORS_ALLOW_ORIGIN"),
		},
		Database: Database{
			DBHost:     os.Getenv("SQL_HOST"),
			DBPort:     os.Getenv("SQL_PORT"),
			DBName:     os.Getenv("SQL_DATABASE"),
			DBUser:     os.Getenv("SQL_USER"),
			DBPassword: os.Getenv("SQL_PASSWORD"),
		},
		Cache: Cache{
			CacheHost:     os.Getenv("REDIS_HOST"),
			CachePassword: os.Getenv("REDIS_PASSWORD"),
			CachePort:     os.Getenv("REDIS_PORT"),
		},
	}

	timefinish := time.Now()
	fmt.Println("Finish Load Config " + timefinish.Format("2006-01-02 15:04:05"))
	return properties
}

func databaseConnect(properties Properties) *gorm.DB {
	params := properties.Database
	dsn := "host=" + params.DBHost + " port=" + params.DBPort + " user=" + params.DBUser + " dbname=" + params.DBName + " password=" + params.DBPassword + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.OrderItems{})
	db.AutoMigrate(&domain.OrderHistories{})

	return db
}

func ConnectCache(properties Properties) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     properties.Cache.CacheHost + ":" + properties.Cache.CachePort,
		Password: properties.Cache.CachePassword,
		DB:       0, // use default DB
	})

	return client
}
