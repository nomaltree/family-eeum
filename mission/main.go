package main

import (
	"fmt"
	"github.com/nomaltree/family-eeum/mission/handler"
	"github.com/nomaltree/family-eeum/mission/service"
	"github.com/nomaltree/family-eeum/mission/storage"
	"log"
	"net/http"
)

func main() {
	mTest, err := storage.SetStorage()
	if err != nil {
		log.Fatal("Error while set up storage: ", err)
	}
	mService := service.NewMissionService(mTest)
	mRouter := handler.MissionHandler(mService)

	fmt.Println("Starting server while at port 8084: ")
	log.Fatal(http.ListenAndServe(":8084", mRouter))
}
