package main

import (
	categoryRoutes "Project1/internal/category/routes"
	postRoutes "Project1/internal/post/routes"
	userRoutes "Project1/internal/users/routes"
	"Project1/pkg/database/dbcon"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	DB, err := dbcon.ConnectToDB(dbcon.Config{
		Host:     viper.GetString("DB.host"),
		Port:     viper.GetString("DB.port"),
		Username: viper.GetString("DB.username"),
		DBName:   viper.GetString("DB.dbname"),
		SSLMode:  viper.GetString("DB.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initialize DB: %s", err.Error())
	}

	app := gin.Default()
	app.Static("/static", "./uploads")
	api := app.Group("/api")
	userRoutes.InitUserRoutes(api, DB)
	categoryRoutes.InitCategoryRoutes(api, DB)
	postRoutes.InitPostRoutes(api, DB)

	if err := app.Run("localhost:8000"); err != nil {
		log.Fatalf("Failed running app: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("Configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
