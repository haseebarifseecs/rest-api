package controllers

import "net/http"

func CreateHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(200)
	resp.Write([]byte("You hit an endpoint"))
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
