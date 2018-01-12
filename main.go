package main

import (
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/url-shortener/handler"
	"github.com/url-shortener/storage/mysql"
	"github.com/url-shortener/storage/redis"
)

// Config describes configuration of redis client
type Config struct {
	redisURI,
	redisPass,
	mysqlURI string
}

func main() {
	conf := new(Config)
	service := micro.NewService(
		micro.Name("go.micro.api.urlshortener"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:        "redis_uri",
				Usage:       "Describe Redis URI in 'host:port' format",
				Destination: &conf.redisURI,
				EnvVar:      "redis_uri",
			},
			cli.StringFlag{
				Name:        "redis_pass",
				Usage:       "Describe Redis password",
				Destination: &conf.redisPass,
				EnvVar:      "redis_pass",
			},
			cli.StringFlag{
				Name:        "mysql_uri",
				Usage:       "Describe MySQL URI in 'username:password@/db",
				Destination: &conf.mysqlURI,
				EnvVar:      "mysql_uri",
			},
		),
	)

	service.Init()

	if len(conf.redisURI) == 0 {
		log.Fatal("redis URI is required, please specify redis_uri option")
	}

	if len(conf.mysqlURI) == 0 {
		log.Fatal("mysql URI is required, please specify mysql_uri option")
	}

	log.Printf("Redis address set to %v", conf.redisURI)
	log.Printf("MySQL address set to %v", conf.mysqlURI)

	s := service.Server()
	cacher, err := redis.New(conf.redisURI, conf.redisPass)
	if err != nil {
		log.Fatalf("could not create redis cacher: %v", err)
	}

	storage, err := mysql.New(conf.mysqlURI)
	if err != nil {
		log.Fatalf("could not create mysql storage: %v", err)
	}

	urlHandler := s.NewHandler(handler.New(cacher, storage))
	s.Handle(urlHandler)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
