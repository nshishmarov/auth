package main

import (
	"auth/main/api"
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)



func main()  {
	fmt.Println("App starts!")
	r := router.New()	
	r.POST("/token", controller.Validate)
	log.Fatal(fasthttp.ListenAndServe("localhost:8080", r.Handler))
}