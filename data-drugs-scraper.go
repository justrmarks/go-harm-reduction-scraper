package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
)

func main() {
	url := "https://www.ecstasydata.org/"

	// Instantiate default collector
	c := colly.NewCollector(
		// Turn on asynchronous requests
		colly.Async(true),
		// Attach a debugger to the collector
		colly.Debugger(&debug.LogDebugger{}),
		// colly.AllowedDomains("ecstasydata.org"),
		// Allow visiting the same page multiple times
		colly.AllowURLRevisit(),
		// Allow crawling to be done in parallel / async
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*.ecstasydata.org",
		Parallelism: 2,
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL, r.StatusCode)
	})

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	// Wait until threads are finished

	err := c.Visit(url)
	if err != nil {
		fmt.Println(err)
	}

	c.Wait()
}
