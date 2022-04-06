package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/micro/go-micro/web"
)

func main() {
	// New web service
	service := web.NewService(
		web.Name("go.micro.api.webservice"),
		web.Address(":8081"),
	)

	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	// static files
	service.Handle("/webservice/", http.StripPrefix("/webservice/", http.FileServer(http.Dir("html"))))
	// webservice interface
	service.HandleFunc("/webservice/hi", hi)

	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}

func hi(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	_ = r.ParseForm()
	// 返回结果
	response := map[string]interface{}{
		"ref":  time.Now().UnixNano(),
		"data": "Hello! " + r.Form.Get("name"),
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
