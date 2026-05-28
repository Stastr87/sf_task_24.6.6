package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	sr "24.6.6/pkg"
	"golang.org/x/net/html"
)

func extractLinks(url sr.SimpleNewReader) {

	resp, err := http.Get(url.Data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var visit func(*html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					fmt.Println(attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(doc)
}

func readURLFromConsole() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите URL: ")
	scanner.Scan()
	return scanner.Text()
}

func main() {

	var url sr.SimpleNewReader

	input := readURLFromConsole()
	url = *sr.NewSimpleReader(input)
	url.CountData([]byte(input))
	fmt.Println("Принято символов:", url.DataLen)
	extractLinks(url)
}
