package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// Sitemapindex struct
type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

// News struct
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// NewsMap struct
type NewsMap struct {
	Keyword  string
	Location string
}

// NewsAggPage struct
type NewsAggPage struct {
	Title       string
	Description map[string]NewsMap
}

var wg sync.WaitGroup

func newsRoutine(c chan News, Location string) {
	defer wg.Done()
	var n News
	resp, err := http.Get(Location)
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()

	c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s Sitemapindex

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	newsMap := make(map[string]NewsMap)
	resp.Body.Close()
	queue := make(chan News, 50)

	// range will iterate over data structure -- returns key and value
	for _, Location := range s.Locations {
		wg.Add(1)
		Location = strings.TrimSpace(Location)
		go newsRoutine(queue, Location)
	}
	wg.Wait()
	close(queue)

	for elem := range queue {
		for idx, err := range elem.Keywords {
			if err != "" {
				fmt.Println(err)
			}

			newsMap[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "Try a new chance well its need", Description: newsMap}
	t, err := template.ParseFiles("src\\basic-web\\templating\\1\\index.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Execute(w, p))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Who, yo got golang</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/newsagg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
