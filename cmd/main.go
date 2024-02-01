package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Sotnasjeff/web-crawler-golang/db"
	"golang.org/x/net/html"
)

var (
	visited map[string]bool = map[string]bool{}
)

func main() {
	visitLink("https://google.com.br")
}

type VisitedLink struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"linke"`
	VisitedDate time.Time `bson:"visited_date"`
}

func visitLink(url string) {
	if ok := visited[url]; ok {
		return
	}
	visited[url] = true
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic(err)
	}

	doc, err := html.Parse(res.Body)

	if err != nil {
		panic(err)
	}

	extractLinks(doc)
}

func extractLinks(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, att := range node.Attr {
			if att.Key != "href" {
				continue
			}

			link, err := url.Parse(att.Val)

			if err != nil || link.Scheme == "" {
				continue
			}

			visitedLink := VisitedLink{
				Website:     link.Host,
				Link:        link.String(),
				VisitedDate: time.Now(),
			}

			db.Insert("links", visitedLink)

			visitLink(link.String())
		}

	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		extractLinks(n)
	}
}
