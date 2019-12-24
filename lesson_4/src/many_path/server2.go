package main

import (
	"many_path/views"
	"net/http"
)

func main() {
	http.HandleFunc("/admin", views.AdminView)
	http.HandleFunc("/list", views.ListView)
	http.HandleFunc("/", views.HomeView)

	if err := http.ListenAndServe("localhost:5555", nil); err != nil {
		panic("PANIC!!!")
	}
}

