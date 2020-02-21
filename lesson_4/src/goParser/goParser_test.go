package main

import (
	"fmt"
	"goParser/LoggerParser"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
)

func TestLoginClient(t *testing.T) {
	LoginClient()
}

func TestParseVideoUrl(t *testing.T) {
	cookieJar, _ := cookiejar.New(nil)
	client := http.Client{Jar: cookieJar}

	resp, err := client.PostForm("https://www.linux.org.ru/login.jsp", url.Values{
		"username":    {"gokain@yandex.ru"},
		"password": {"598759875987"},
	})
	LoggerParser.ForError(err)
	fmt.Println("post:", resp.StatusCode)

}
