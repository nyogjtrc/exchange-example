package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nyogjtrc/exchange-example/internal/rest"
)

func main() {
	err := rest.LoadRateFile("rate.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	r := gin.Default()

	rest.QueryExchange(r)

	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
