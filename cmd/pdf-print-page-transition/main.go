package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	var (
		userDataDir = flag.String("user-data-dir", "./tmp/", "user data dir")
		id          = flag.Int64("id", 1, "id number")
	)
	flag.Parse()

	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.DisableGPU,
		chromedp.UserDataDir(*userDataDir),
	}

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, _ := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.flierinc.com/summary/`+strconv.FormatInt(*id, 10)),
		chromedp.Sleep(5*time.Second),
		chromedp.Click(`//*[@id="block1"]/div[2]/div[2]/div[2]/div[2]`, chromedp.NodeVisible),
		chromedp.Sleep(5*time.Second),
		chromedp.Click(`//*[@id="print"]`, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
}
