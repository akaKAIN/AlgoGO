package http_server

import (
	"fmt"
	"log"
	"net/http"
)

func serverStatus(res http.ResponseWriter, req *http.Request) {
	if _, err := fmt.Fprint(res, "<h1 style='text-align: center;'>Server was running ... </h1>"); err != nil {
		log.Printf("Error: %s\n", err)
	}
}

func RunServer() {
	http.HandleFunc("/", serverStatus)
	if err := http.ListenAndServe("localhost:5556", nil); err != nil {
		log.Printf("Listen error: %s", err)
	}
}
