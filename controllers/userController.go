package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	user "github.com/haseebarifseecs/rest-api/model"
)

var users []user.User

type Response struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	OK   string `json:"status"`
}

func CreateHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")

	response := Response{
		Id:   1,
		Name: "Haseeb",
		OK:   "Success",
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	if req.Method == "GET" {
		resp.Write(jsonData)
	}

	if req.Method == "POST" {
		err := req.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(req.Form.Get("Test"))
		resp.Write([]byte(req.Form.Get("Test")))
	}
	// resp.Write([]byte("message:hello"))
}

func UpdateHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(200)
}

func DeleteHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(200)
}

func ReadHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(200)
}
