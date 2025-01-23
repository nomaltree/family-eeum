package main

import (
	"fmt"
	"github.com/nomaltree/family-eeum/feed/handler"
	"github.com/nomaltree/family-eeum/feed/service"
	"github.com/nomaltree/family-eeum/feed/storage"
	"log"
	"net/http"
)

func main() {
	fTest, err := storage.SetStorage()
	if err != nil {
		log.Fatal("Error while set up storage: ", err)
	}
	fService := service.NewFeedService(fTest)
	fRouter := handler.FeedHandler(fService)

	fmt.Println("Starting server while at port 8083: ")
	log.Fatal(http.ListenAndServe(":8083", fRouter))
}
