package main

import (
	"net/http"
)

type webserver struct {
	serverRun string
}

func (ws *webserver) Test(w http.ResponseWriter, r *http.Request) {
	logger.Println("now in testing")
	ret := "aaa"
	w.Write([]byte(ret))
}

func (ws *webserver) webserver() {
	logger.Println("now create webserver")
	http.HandleFunc("/test", ws.Test)        // set router
	err := http.ListenAndServe(":3030", nil) // set listen port
	if err != nil {
		logger.Println("ListenAndServe: ", err)
	}
}
