package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

const (
	baseUrl        = "http://szegeder.hu/"
	articleLink    = "article.post h1.entry-title a"
	articleTitle   = "article.post h1.entry-title"
	articleDate    = "article.post time.entry-date"
	articleAuthor  = "article.post span.author a.url"
	articleContent = "article.post div.entry-content p"
)

func main() {
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
	})

	listingCollector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := listingCollector.Visit(baseUrl)

	if err != nil {
		fmt.Println(err)
	}
}
