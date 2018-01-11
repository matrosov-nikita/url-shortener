package main

import (
	"fmt"
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/url-shortener/handler"
	"github.com/url-shortener/storage/mysql"
	"github.com/url-shortener/storage/redis"
)

// Config describes configuration of redis client
type Config struct {
	redisAddr,
	redisPass,
	mysqlAddr string
}

func main() {
	conf := new(Config)
	service := micro.NewService(
		micro.Name("url-shortener"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:        "REDIS_URL",
				Usage:       "Describe Redis URL in 'host:port' format",
				Destination: &conf.redisAddr,
				EnvVar:      "REDIS_URL",
			},
			cli.StringFlag{
				Name:        "REDIS_PASS",
				Usage:       "Describe Redis password",
				Destination: &conf.redisPass,
				EnvVar:      "REDIS_PASS",
			},
			cli.StringFlag{
				Name:        "MYSQL_URL",
				Usage:       "Describe MySQL address in 'username:password@/db",
				Destination: &conf.mysqlAddr,
				EnvVar:      "MYSQL_URL",
			},
		),
	)

	service.Init()

	if len(conf.redisAddr) == 0 {
		log.Fatal("redis address is required, please specify REDIS_URL option")
	}

	if len(conf.mysqlAddr) == 0 {
		log.Fatal("mysql address is required, please specify MYSQL_URL option")
	}

	log.Printf("Redis address set to %v", conf.redisAddr)
	log.Printf("MySQL address set to %v", conf.mysqlAddr)

	s := service.Server()
	cacher, err := redis.New(conf.redisAddr, conf.redisPass)
	if err != nil {
		log.Fatalf("could not create redis cacher: %v", err)
	}

	storage, err := mysql.New(conf.mysqlAddr)
	if err != nil {
		log.Fatalf("could not create mysql storage: %v", err)
	}

	urlHandler := s.NewHandler(handler.New(cacher, storage))
	s.Handle(urlHandler)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
