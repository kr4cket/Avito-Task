package app

import (
	"avitoTask/pkg/handlers"
	"avitoTask/pkg/repository"
	"avitoTask/pkg/server"
	"avitoTask/pkg/services"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func Run() {

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error loading env vars: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configuration %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("error initializing db %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := services.NewService(repos)
	handlers := handlers.New(service)
	server := new(server.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running Http-Server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
