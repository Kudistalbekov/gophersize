package main

import (
	"log"
	"os"

	"golang.org/x/net/html"
)

type link struct {
	href string
	text string
}

func parsehtml(file string) *[]link {
	data := &[]link{}
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	node, err := html.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	var function func(*html.Node, string)
	function = func(n *html.Node, space string) {
		if n.Type == html.ElementNode && n.Data == "a" {
			if n.FirstChild.Type == html.TextNode {
				d := link{
					href: n.Attr[0].Val,
					text: n.FirstChild.Data,
				}
				*data = append(*data, d)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			function(c, space+"  ")
		}
	}
	function(node, "")
	return data
}
