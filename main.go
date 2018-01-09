package main

import (
	"fmt"
	"log"
	"url-shortener/handler"
	"url-shortener/storage"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	fmt.Println("main")
	service := micro.NewService(
		micro.Name("url-shortener"),
		micro.Version("latest"),
		micro.Flags(cli.StringFlag{
			Name:  "mysql",
			Usage: "Describe MySQL URL with 'user:password@/dbname' format",
		}),
	)

	fmt.Println("before init")
	service.Init(
		micro.Action(func(c *cli.Context) {
			fmt.Println("I'm there")
			source := c.String("mysql")
			fmt.Println("MySQL source set to", source)
			s := service.Server()
			fmt.Printf("source = %s", source)
			st, err := storage.New(source)
			if err != nil {
				log.Fatalf("could not initialize MySQL db with source %s", source)
			}

			s.Handle(s.NewHandler(handler.New(st)))
		}),
	)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(service)
}
