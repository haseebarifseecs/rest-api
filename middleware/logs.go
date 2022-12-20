package middleware

import (
	"log"
	"net/http"
)

func Logs(resp http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s", req.Method, req.URL.Path, req.Host)
}
