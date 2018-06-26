package router

import (
	"net/http"
)

var CreateBuildHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte(payload))
})
