package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"sr-skilltest/internal/model"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getProperties() (properties model.Properties) {
	isLocalDev := flag.Bool("local", false, "=(true/false)")
	flag.Parse()
	if *isLocalDev {
		return loadPropertiesLocal()
	} else {
		return loadEnv()
	}
}

func loadPropertiesLocal() model.Properties {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	properties := loadEnv()
	return properties
}

func loadEnv() model.Properties {
	timestart := time.Now()
	fmt.Println("Starting Load Config " + timestart.Format("2006-01-02 15:04:05"))

	properties := model.Properties{
		App: model.App{
			Mode:            os.Getenv("APP_MODE"),
			Debug:           os.Getenv("APP_DEBUG_MODE"),
			BaseURL:         os.Getenv("APP_BASE_URL"),
			CorsAllowOrigin: os.Getenv("CORS_ALLOW_ORIGIN"),
		},
		Database: model.Database{
			DBHost:     os.Getenv("SQL_HOST"),
			DBPort:     os.Getenv("SQL_PORT"),
			DBName:     os.Getenv("SQL_DATABASE"),
			DBUser:     os.Getenv("SQL_USER"),
			DBPassword: os.Getenv("SQL_PASSWORD"),
		},
		Cache: model.Cache{
			CacheHost:     os.Getenv("REDIS_HOST"),
			CachePassword: os.Getenv("REDIS_PASSWORD"),
			CachePort:     os.Getenv("REDIS_PORT"),
		},
	}

	timefinish := time.Now()
	fmt.Println("Finish Load Config " + timefinish.Format("2006-01-02 15:04:05"))
	return properties
}

func databaseConnect(properties model.Properties) *gorm.DB {
	params := properties.Database
	dsn := "host=" + params.DBHost + " port=" + params.DBPort + " user=" + params.DBUser + " dbname=" + params.DBName + " password=" + params.DBPassword + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}

	return db
}

func ConnectCache(properties model.Properties) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     properties.Cache.CacheHost + ":" + properties.Cache.CachePort,
		Password: properties.Cache.CachePassword,
		DB:       0, // use default DB
	})

	return client
}
