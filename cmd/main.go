package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
)

var X_URL = []string{
	"https://x.com/shanren2011/status/1880591123405770857",
}

func main() {
	fmt.Println("ROD SCRAPING....")
	Run()
	fmt.Println("ROD SCRAPING....DONE")

}
func ReadCSV() {

}
func Run() {
	page := rod.New().MustConnect().MustPage(X_URL[0])
	mctx, txtCancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(10 * time.Second)
		txtCancel()
	}()
	txtEl, err := page.Context(mctx).Element("div[data-testid=\"tweetText\"]")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(txtEl.Text())
	}
	imgEl, err := page.Context(mctx).Element("div[data-testid=\"tweetPhoto\"] > img ")
	if err != nil {
		fmt.Println(err)
	} else {
		_ = utils.OutputFile("b.png", imgEl.MustResource())
	}
	// page.MustWaitDOMStable().MustScreenshot("a.png")
}
