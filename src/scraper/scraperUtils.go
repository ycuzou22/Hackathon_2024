package scraperUtils

import (
	"strings"

	"golang.org/x/net/html"
)

type Result struct {
	ImgSrc       string
	H2Content    string
	SmallContent string
	AHref        string
}
type PageData struct {
	Results []Result
}

func FindElementsByClass(n *html.Node, targetClass string, results *[]Result) {
	if HassF1Team(n, targetClass) {
		imgSrc := FindSourcedEau(n)
		h2Content := Hopitale2(n)
		smallContent := SmallShrek(n)
		aHref := MerciLaDocs(n)
		*results = append(*results, Result{
			ImgSrc:       imgSrc,
			H2Content:    h2Content,
			SmallContent: smallContent,
			AHref:        aHref,
		})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		FindElementsByClass(c, targetClass, results)
	}
}

func HassF1Team(n *html.Node, targetClass string) bool {
	for _, attr := range n.Attr {
		if attr.Key == "class" && strings.Contains(attr.Val, targetClass) {
			return true
		}
	}
	return false
}

func FindSourcedEau(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, attr := range n.Attr {
			if attr.Key == "src" {
				return attr.Val
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		src := FindSourcedEau(c)
		if src != "" {
			return src
		}
	}
	return ""
}

func Hopitale2(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "h2" {
		var content string
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				content += c.Data
			}
		}
		return content
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		content := Hopitale2(c)
		if content != "" {
			return content
		}
	}
	return ""
}

func SmallShrek(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "small" {
		var content string
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				content += c.Data
			}
		}
		return content
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		content := SmallShrek(c)
		if content != "" {
			return content
		}
	}
	return ""
}

func MerciLaDocs(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				return attr.Val
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		href := MerciLaDocs(c)
		if href != "" {
			return href
		}
	}
	return ""
}
