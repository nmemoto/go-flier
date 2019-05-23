package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	var (
		userDataDir = flag.String("user-data-dir", "./tmp/", "user data dir")
		headless    = flag.Bool("headless", true, "Headless or not")
	)
	flag.Parse()

	var opts []chromedp.ExecAllocatorOption
	if *headless == false {
		opts = []chromedp.ExecAllocatorOption{
			chromedp.NoFirstRun,
			chromedp.NoDefaultBrowserCheck,
			chromedp.DisableGPU,
			chromedp.UserDataDir(*userDataDir),
		}
	} else {
		opts = []chromedp.ExecAllocatorOption{
			chromedp.NoFirstRun,
			chromedp.NoDefaultBrowserCheck,
			chromedp.DisableGPU,
			chromedp.UserDataDir(*userDataDir),
			chromedp.Headless,
		}
	}

	user := os.Getenv("FLIER_USER")
	pass := os.Getenv("FLIER_PASS")

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	if *headless == true {
		defer cancel()
	}

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.flierinc.com/`),
		chromedp.Sleep(5*time.Second),
		chromedp.Click(`#header-utils > button.btn.btn-fl-login > span`, chromedp.NodeVisible),
		chromedp.Sleep(5*time.Second),
		chromedp.Focus("#email", chromedp.ByID),
		chromedp.KeyAction(user),
		chromedp.Focus("#password", chromedp.ByID),
		chromedp.KeyAction(pass),
		chromedp.Click(`#login-form > div > div > div.modal-body > form > div.form-group.center.mt30 > button.btn.btn-fl-sm.btn-fl-submit.arrow-right-white`, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
}
