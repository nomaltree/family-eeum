package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	router := mainHandler()

	fmt.Println("Starting server while at port 8082: ")
	http.ListenAndServe(":8082", router)
}

func mainHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
		userId, err := authHandler(r.Header.Get("token"))
		path := mux.Vars(r)["path"]
		if err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
		} else {
			Id := strconv.Itoa(userId)
			path = path + "?userId=" + Id
			// path에 따라 Redirect 로직 처리
			// 각 기능에 대한 포트번호는 임의로 할당
			if strings.HasPrefix(path, "feed") {
				path = "http://localhost:8083/" + path
				http.Redirect(w, r, path, http.StatusMovedPermanently)
			} else if strings.HasPrefix(path, "mission") {
				path = "http://localhost:8084/" + path
				http.Redirect(w, r, path, http.StatusMovedPermanently)
			} else if strings.HasPrefix(path, "notification") {
				path = "http://localhost:8085/" + path
				http.Redirect(w, r, path, http.StatusMovedPermanently)
			} else if strings.HasPrefix(path, "survey") {
				path = "http://localhost:8086/" + path
				http.Redirect(w, r, path, http.StatusMovedPermanently)
			}
		}
	}).Methods("GET")

	return router
}

func authHandler(token string) (int, error) {
	Body := []byte(token)
	bodyReader := bytes.NewReader(Body)
	requestURL := fmt.Sprintf("http://localhost:port/회원인증 엔드포인트")
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		log.Println("Error while making http request: ", err)
		os.Exit(1)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error while sending http request: ", err)
		os.Exit(1)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error while reading response body: ", err)
	}
	userId, err := strconv.Atoi(string(resBody))
	if err != nil {
		log.Println("Error while converting response body to int: ", err)
	}
	return userId, err
}
