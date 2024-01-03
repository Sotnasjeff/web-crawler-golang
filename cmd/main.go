package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

var (
	links   []string
	visited map[string]bool = map[string]bool{}
)

func main() {
	visitLink("https://google.com.br")

	fmt.Println(len(links))
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
			links = append(links, link.String())

			visitLink(link.String())
		}

	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		extractLinks(n)
	}
}
