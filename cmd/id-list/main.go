package main

import (
	"context"
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
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.flierinc.com/summary/list`+pageURL),
		chromedp.Sleep(5*time.Second),
		chromedp.Nodes(`//div[contains(@class, "summary-md")]//div[@class="summary-md"]/a/@href`, &links),
	)
	for _, v := range links {
		str := v.Attributes[1]
		rep := regexp.MustCompile(`/summary/([0-9]*)`)
		i, _ := strconv.Atoi(rep.FindStringSubmatch(str)[1])
		fmt.Fprintf(os.Stdout, "%#v\n", i)
	}
	if err != nil {
		log.Fatal(err)
	}
}
