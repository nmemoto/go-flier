package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {

	var (
		page     = flag.Int("page", 1, "page number")
		headless = flag.Bool("headless", true, "Headless or not")
	)
	flag.Parse()
	var pageURL string
	if *page == 1 {
		pageURL = ""
	} else if *page > 1 {
		pageStr := strconv.Itoa(*page)
		pageURL = "?page=" + pageStr
	} else {
		fmt.Fprintf(os.Stderr, "%v is wrong page number. you should enter more than 1.\n", *page)
		os.Exit(1)
	}

	var opts []chromedp.ExecAllocatorOption
	if *headless == false {
		opts = []chromedp.ExecAllocatorOption{
			chromedp.NoFirstRun,
			chromedp.NoDefaultBrowserCheck,
			chromedp.DisableGPU,
		}
	} else {
		opts = []chromedp.ExecAllocatorOption{
			chromedp.NoFirstRun,
			chromedp.NoDefaultBrowserCheck,
			chromedp.DisableGPU,
			chromedp.Headless,
		}
	}

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	if *headless == true {
		defer cancel()
	}

	var links []*cdp.Node
	var titles []*cdp.Node
	var authors []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.flierinc.com/summary/list`+pageURL),
		chromedp.Sleep(10*time.Second),
		chromedp.Nodes(`//div[@class="summary-md"]/*/a/@href`, &links),
		chromedp.Nodes(`//div[@class="summary-md"]//div[@class="summary-title"]//text()`, &titles),
		chromedp.Nodes(`//div[@class="summary-md"]//div[@class="summary-author"]/span[@class="pr5"][1]//text()`, &authors),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "page %v is failed.\n", *page)
		log.Fatal(err)
	}

	records := [][]string{}

	for i, v := range links {
		str := v.Attributes[1]
		rep := regexp.MustCompile(`/summary/([0-9]*)`)
		idx := rep.FindStringSubmatch(str)[1]
		title := titles[i].NodeValue
		author := authors[i].NodeValue
		url := "https://www.flierinc.com" + str
		records = append(records, []string{idx, title, author, url})
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
