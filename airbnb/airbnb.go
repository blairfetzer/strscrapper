package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Missing location argument")
		os.Exit(1)
	}


	location := os.Args[1]

	c := colly.NewCollector(colly.AllowedDomains(
		"airbnb.com",
		"https://airbnb.com",
		"https://airbnb.com/",
		"airbnb.com/",
		"www.airbnb.com",
		"https://www.airbnb.com",
	))

	c.OnHTML(`div[aria-live="polite"]`, func(e *colly.HTMLElement) {
		fmt.Println("result: ", e)
	})

	//c.OnHTML(`div[data-pageslot] * > div[aria-live="polite"]`, func(e *colly.HTMLElement) {
	//	//e.ForEach("meta", func(_ int, elem *colly.HTMLElement) {
	//	//	elem.Request.Visit(elem.Attr("itemprop"))
	//	//})
	//	fmt.Println("div: ", e.ChildAttr("div","class"))
	//	fmt.Println("e: " , e)
	//	//url, _ := e.DOM.ParentsUntil("~").Attr("content")
	//	//metaTags := e.DOM.Find(`div[aria-live="polite"] * > div[itemprop="itemListElement"]`)
	//	//metaTags := e.ChildAttrs("meta", `itemProp="url"`)
	//	//fmt.Println("metaTags: ", metaTags)
	//	e.ForEach("meta[itemprop]", func(_ int, el *colly.HTMLElement) {
	//		imgSrc := el.Attr("content")
	//		fmt.Println("imgSrc", imgSrc)
	//		fmt.Println("text:", el.Text)
	//	})
	//	//urls.Each(func(_ int, s *goquery.Selection) {
	//	//	fmt.Println("hello", s.Contents())
	//	//	val := s.Find(`meta[itemprop="url"]`)
	//	//	fmt.Println("val ", val)
	//	//	foo := s.Find("#gsgwcjk")
	//	//	fmt.Println("foo size ", foo.Size())
	//		//link, bar := val.Attr(".content")
	//		//fmt.Println("link ", link, bar)
	//		//val, _ := s.Attr("data-pageslot")
	//		//fmt.Println("Attr ", val)
	//		//requestIDURL := e.Request.AbsoluteURL(e.ChildAttr(`meta[itemprop="url"]`, "content"))
	//		//e.Request.Visit(requestIDURL)
	//	//})
	//
	//	//goquerySelection := e.DOM
	//	//fmt.Println(goquerySelection.Find("a[href]").Attr("href"))
	//})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code:", r.StatusCode)
	})

	c.Visit("https://www.airbnb.com/s/" + location + "/homes")
}