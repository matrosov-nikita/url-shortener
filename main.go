package main

import (
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/url-shortener/handler"
	"github.com/url-shortener/storage/mysql"
	"github.com/url-shortener/storage/redis"
)

// Config describes configuration for databases
type Config struct {
	redisURI,
	mysqlURI string
}

func main() {
	conf := new(Config)
	service := micro.NewService(
		micro.Name("my.service.urlshortener"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:        "redis_uri",
				Usage:       "Describe Redis URI",
				Destination: &conf.redisURI,
				EnvVar:      "REDIS_URI",
				Value:       "localhost: 6379",
			},
			cli.StringFlag{
				Name:        "mysql_uri",
				Usage:       "Describe MySQL URI",
				Destination: &conf.mysqlURI,
				EnvVar:      "MYSQL_URI",
				Value:       "test:test@/urls",
			},
		),
	)

	service.Init()

	if len(conf.redisURI) == 0 {
		log.Fatal("redis URI is required")
	}

	if len(conf.mysqlURI) == 0 {
		log.Fatal("mysql URI is required")
	}

	log.Printf("Redis address set to %v", conf.redisURI)
	log.Printf("MySQL address set to %v", conf.mysqlURI)

	cacher, err := redis.New(conf.redisURI)
	if err != nil {
		log.Fatalf("could not create redis cacher: %v", err)
	}

	storage, err := mysql.New(conf.mysqlURI)
	if err != nil {
		log.Fatalf("could not create mysql storage: %v", err)
	}

	s := service.Server()
	urlHandler := s.NewHandler(handler.New(cacher, storage))
	s.Handle(urlHandler)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
