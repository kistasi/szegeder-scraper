package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

const (
	baseUrl        = "http://szegeder.hu"
	numberOfPages  = 31
	articleLink    = "article.post h1.entry-title a"
	articleTitle   = "article.post h1.entry-title"
	articleDate    = "article.post time.entry-date"
	articleAuthor  = "article.post span.author a.url"
	articleContent = "article.post div.entry-content p"
)

func main() {
	for page := 0; page < numberOfPages; page++ {
		scrape(fmt.Sprintf("%s/page/%d", baseUrl, page))
	}
}

func scrape(url string) {
	listingCollector := colly.NewCollector()

	listingCollector.OnHTML(articleLink, func(article *colly.HTMLElement) {
		articleCollector := colly.NewCollector()

		articleCollector.OnHTML("body", func(articlePage *colly.HTMLElement) {
			title := articlePage.DOM.Find(articleTitle).Text()
			date := articlePage.DOM.Find(articleDate).Text()
			author := articlePage.DOM.Find(articleAuthor).Text()
			content := articlePage.DOM.Find(articleContent).Text()

			fmt.Println(title)
			fmt.Println(date)
			fmt.Println(author)
			fmt.Println(content)
		})

		articleCollector.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})

		err := articleCollector.Visit(article.Attr("href"))

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(2 * time.Second)
	})

	listingCollector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := listingCollector.Visit(url)

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)
}
