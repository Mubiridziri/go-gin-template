package main

import (
	"app/internal/app"
	"flag"
	"log"
)

var (
	bindAddr string
)

func init() {
	flag.StringVar(&bindAddr, "port", "8080", "listen server port")
}

func main() {
	flag.Parse()

	a := app.New(":" + bindAddr)

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
