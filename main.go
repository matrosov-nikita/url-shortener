package main

import (
	"fmt"
	"log"
	"url-shortener/handler"
	"url-shortener/storage"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

// Config describes configuration of redis client
type Config struct {
	redisAddr, redisPass string
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
		),
	)

	service.Init()

	if len(conf.redisAddr) == 0 {
		log.Fatalln("redis address is required, please specufy redis-addr flag")
	}

	log.Printf("redis address set to %v", conf.redisAddr)

	s := service.Server()
	st, err := storage.New(conf.redisAddr, conf.redisPass)
	if err != nil {
		log.Fatalf("could not create storage: %v", err)
	}

	s.Handle(s.NewHandler(handler.New(st)))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
