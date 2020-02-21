package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"goParser/LoggerParser"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func main() {
	//url := "https://geekbrains.ru"
	//r := getResp(url)
	//doc, err := goquery.NewDocumentFromReader(r.Body)
	//MyLogger.ForError(err)
}

func LoginClient() {
	cookieJar, _ := cookiejar.New(nil)
	client := http.Client{Jar: cookieJar}

	resp, err := client.PostForm("https://geekbrains.ru/chapters/5414", url.Values{
		"user_email":    {"gokain@yandex.ru"},
		"user_password": {"598759875987"},
	})
	LoggerParser.ForError(err)

	r, err := client.Get("https://geekbrains.ru/chapters/5414")
	LoggerParser.ForError(err)
	defer resp.Body.Close()

	text, err := ioutil.ReadAll(r.Body)
	LoggerParser.ForError(err)
	if strings.Contains(string(text), "Запомнить меня"){
		fmt.Println("No access.")
	}


}

func getResp(url string) *http.Response {
	resp, err := http.Get(url)
	LoggerParser.ForError(err)
	defer func() {
		err := resp.Body.Close()
		LoggerParser.ForError(err)
	}()
	if resp.StatusCode != 200 {
		errText := fmt.Sprintf("Error of connection: %d  url: %s", resp.StatusCode, url)
		LoggerParser.ForError(errors.New(errText))
	}
	return resp
}

func ParseVideoUrl(url string) (string, error) {
	r := getResp(url)
	doc, err := goquery.NewDocumentFromReader(r.Body)
	LoggerParser.ForError(err)

	videoLink, ok := doc.Find("a.lesson-contents__download_row").Attr("href")
	if !ok {
		err := errors.New(fmt.Sprintf("На странице урока (url: %s) ссылка на скачку не найдена.", url))
		return "", err
	}
	return videoLink, nil
}

func GetUrlsList(courseUrl string) []string {
	var UrlsList []string
	r := getResp(courseUrl)
	doc, err := goquery.NewDocumentFromReader(r.Body)
	LoggerParser.ForError(err)

	allLessons := doc.Find("div.lessons-list__item lesson-list__item-forthcoming")
	for i := range allLessons.Nodes {
		lessonBlock := allLessons.Eq(i)
		lessonUrl, ok := lessonBlock.Find("a").Attr("href")
		if ok {
			UrlsList = append(UrlsList, lessonUrl)
		}
	}

	return UrlsList
}

func MakeAbsUrl(url string) string {
	return ""
}
