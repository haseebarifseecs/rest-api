package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	user "github.com/haseebarifseecs/rest-api/model"
)

var users []*user.User
var ID int = 0

type Response struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	OK   string `json:"status"`
}

func CreateHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		err := req.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		name := req.Form.Get("name")
		user := user.User{
			Id:   ID,
			Name: name,
		}
		users = append(users, &user)
		response := Response{
			Id:   ID,
			Name: name,
			OK:   "Success",
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ID += 1
		resp.WriteHeader(http.StatusOK)
		resp.Header().Set("Content-Type", "application/json")
		resp.Write([]byte(jsonData))
	}
	// resp.Write([]byte("message:hello"))
}

func UpdateHandler(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := req.Form.Get("name")
	newName := req.Form.Get("newName")
	fmt.Println(name)
	fmt.Println(newName)
	if len(users) > 0 {
		for _, j := range users {
			if name == j.Name {
				// j.Name = newName
				j.UpdateName(newName)
				resp.Write([]byte("Updated"))
				return
			}
		}
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("User Doesn't exist"))
	} else {
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("User Doesn't Exist"))
	}

}

func DeleteHandler(resp http.ResponseWriter, req *http.Request) {
	// err := req.ParseForm()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	name := req.FormValue("name")
	fmt.Println(name)
	if len(users) > 0 {
		for i, j := range users {
			if name == j.Name {
				fmt.Println("Matched")
				users = append(users[:i], users[i+1:]...)
				resp.WriteHeader(http.StatusOK)
				resp.Write([]byte(name + " Deleted"))
				return
			}
		}
	} else {
		resp.Write([]byte(name + " Doesn't exist"))
	}
}

func ReadHandler(resp http.ResponseWriter, req *http.Request) {
	var resps []Response
	if len(users) > 0 {
		for _, j := range users {
			resps = append(resps, Response{
				Id:   j.Id,
				Name: j.Name,
				OK:   "success",
			})
		}
		jsonData, err := json.Marshal(resps)
		if err != nil {
			log.Fatal(err)
		}
		resp.WriteHeader(http.StatusOK)
		resp.Header().Set("Content-Type", "application/json")
		resp.Write(jsonData)
	} else {
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("{0}"))
	}
}
