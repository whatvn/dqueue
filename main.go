package main

import (
	"log"

	"github.com/whatvn/dqueue/handler"
	"github.com/whatvn/dqueue/helper"
	"github.com/whatvn/dqueue/protobuf"
	"github.com/whatvn/dqueue/worker"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
	"github.com/whatvn/dqueue/database"
	"flag"
)

const (
	Domain      = "DelayQueue"
	Consul      = "127.0.0.1:8500"
	OpenTracing = "127.0.0.1:8200"
)

func init() {
	var db = database.NewDatabase()
	db.Init()
	flag.Set("logtostderr", "true")
	flag.Set("v", "2")
	flag.Parse()
}

func WebServer() {
	// Create service
	apiService := web.NewService(
		web.Name("go.micro.api.message"),
	)

	apiService.Init()

	// Create RESTful handler (using Gin)
	messageHandler := new(handler.MessageHandler)
	router := gin.Default()
	router.POST("/message/getall", messageHandler.GetAllMessages)
	router.POST("/message/getlistbydata/:data", messageHandler.GetListMessageByData)
	router.POST("/message/force/:data", messageHandler.ForceMessage)
	router.POST("/message/getlist/:offset/:limit", messageHandler.GetListMessage)

	// Register Handler
	apiService.Handle("/", router)

	// Run server
	if err := apiService.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	service := helper.NewServer(Domain, Consul, OpenTracing).GetService()
	delayQueue.RegisterDelayQueueHandler(service.Server(), handler.NewMicroServiceHandler())
	queueWorker := worker.NewWorker(helper.GetQueueType())

	go WebServer()

	go queueWorker.Run()

	err := service.Run()
	if err != nil {
		log.Println("server cannot start, error: ", err)
	}
}
