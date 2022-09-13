package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/snsvistunov/films-app/pkg/handler"
	"github.com/snsvistunov/films-app/pkg/repository"
	server "github.com/snsvistunov/films-app/pkg/server"
	"github.com/snsvistunov/films-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error while loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.PGConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	jwtStorage, err := repository.NewRedisStorage(repository.RedisConfig{
		Host: viper.GetString("redis.host"),
		Port: viper.GetString("redis.port"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize jwt storage: %s", err.Error())
	}

	repos := repository.NewRepository(db, jwtStorage)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
