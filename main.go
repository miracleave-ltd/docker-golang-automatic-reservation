package main

import (
    "github.com/sclevine/agouti"
    "log"
)

func main() {
    // Chromeを利用することを宣言
    driver := agouti.ChromeDriver(
        agouti.ChromeOptions("args", []string{
            "--headless",
            "--disable-gpu",
            "--window-size=1280,1024",
            "--disable-dev-shm-usage",
            "--no-sandbox",
        }),
        agouti.Debug,
    )

    if err := driver.Start(); err != nil {
        log.Printf("Failed to start driver: %v", err)
    }
    defer driver.Stop()

    page, err := driver.NewPage(agouti.Browser("chrome"))
    if err != nil {
        log.Printf("Failed to open page: %v", err)
    }

    // Access to a target page
    url := "https://www.google.co.jp/"
    err = page.Navigate(url)
    if err != nil {
        log.Printf("Failed to navigate: %v", err)
    }
    // Get screen shot
    page.Screenshot("screenshot/Google.png")
		// 自動操作
	//ANA日本語ページ遷移
	page.Navigate("https://www.ana.co.jp/ja/jp")
	log.Printf(page.Title())
        page.Screenshot("screenshot/ana-top.png")
	//検索ボタン押下
	page.FirstByName("arrivalAirport").Submit()
	page.Screenshot("screenshot/ana-1.png")
	//区間検索「片道」押下
	page.FindByID("hoge1").Click()
	page.Screenshot("screenshot/ana-2.png")

}
