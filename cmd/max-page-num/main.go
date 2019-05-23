package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var page []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.flierinc.com/summary/list`),
		chromedp.Sleep(5*time.Second),
		chromedp.Nodes(`//div[@id="pagination"]/div/ul/li[7]/a[@class="page-link"]/text()`, &page),
	)
	for _, v := range page {
		i, _ := strconv.Atoi(v.NodeValue)
		fmt.Printf("%#v\n", i)
	}
	if err != nil {
		log.Fatal(err)
	}
}
