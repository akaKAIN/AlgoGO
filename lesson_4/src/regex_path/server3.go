package main

import (
	"many_path/views"
	"net/http"
	"regexp"
)

type regexResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp //сохранение скомпилированных регулярок
}

func newPathResolver() *regexResolver {
	return &regexResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

// Метод добавления маршрута в структуру
func (r *regexResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile(regex)
	r.cache[regex] = cache
}

func (r *regexResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			handlerFunc(res, req)
			return
		}
	}
	http.NotFound(res, req)
}

func main() {
	router := newPathResolver()
	router.Add("GET /admin$", views.AdminView)
	router.Add("GET|POST /", views.HomeView)
	if err := http.ListenAndServe(":5555", router); err != nil {
		panic("Listen!! PANIC!!")
	}
}
