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
		colly.AllowedDomains(url),
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
		RandomDelay: 1 * time.Second
	})

	// Start scraping in five threads on https://httpbin.org/delay/2
	c.OnHTML(".tablet", func(e *colly.HTMLElement) {

	})

	c.Visit(url)
	// Wait until threads are finished
	
}
