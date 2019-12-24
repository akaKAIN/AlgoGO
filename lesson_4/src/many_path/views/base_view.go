package views

import (
	"fmt"
	"net/http"
)

func AdminView(w http.ResponseWriter, r *http.Request){
	getData := r.URL.Query()
	_,_ = fmt.Fprintf(w, "URL: %v\n", getData["hello"])
}

func ListView(w http.ResponseWriter, r *http.Request){
	_,_ = fmt.Fprintf(w, "URL.Path: %s", r.URL.Path)
}

func HomeView(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}
	_,_ = fmt.Fprintf(w, "URL: <h1>'%s'</h1>", "HomePage")
}