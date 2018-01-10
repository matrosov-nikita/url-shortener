package main

import (
	"fmt"
	"log"
	"url-shortener/handler"
	"url-shortener/storage/mysql"
	"url-shortener/storage/redis"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
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
				Name:        "redis-addr",
				Usage:       "Describe Redis URL with 'host:port' format",
				Destination: &conf.redisAddr,
			},
			cli.StringFlag{
				Name:        "redis-pass",
				Usage:       "Describe Redis password",
				Destination: &conf.redisPass,
			},
			cli.StringFlag{
				Name:        "mysql-addr",
				Usage:       "Describe MySQL address",
				Destination: &conf.mysqlAddr,
			},
		),
	)

	service.Init()

	if len(conf.redisAddr) == 0 {
		log.Fatalln("redis address is required, please specufy redis-addr flag")
	}

	if len(conf.mysqlAddr) == 0 {
		log.Fatalln("MySQL address is required, please specufy mysql-addr flag")
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
		log.Fatalf("could not create MySQL storage: %v", err)
	}

	s.Handle(s.NewHandler(handler.New(cacher, storage)))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
