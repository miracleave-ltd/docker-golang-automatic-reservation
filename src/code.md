# ex. プログラム全量

```go
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
    page.FindByID("buttonOneWay").Click()
    page.Screenshot("screenshot/ana-2.png")
    //到着地「札幌」選択
    page.FindByID("arrivalAirport").Select("札幌(千歳)")
    page.Screenshot("screenshot/ana-3.png")
    //カレンダーテキスト押下
    page.FindByID("outwardEmbarkationDate").Click()
    page.Screenshot("screenshot/ana-4.png")
    //搭乗日カレンダーを指定
    page.FirstByXPath("/html/body/div[8]/div/div/div/div/div[1]/table/tbody/tr[5]/td[3]/a").Click()
    page.Screenshot("screenshot/ana-5.png")
    //最安値指定
    page.FirstByLabel("最安運賃を検索").Click()
    page.Screenshot("screenshot/ana-6.png")
    //検索ボタン押下
    page.FindByButton("検索する").Click()
    page.Screenshot("screenshot/ana-7.png")
    //値段を押下
    page.FirstByLabel("23,960円").Click()
    page.Screenshot("screenshot/ana-8.png")
    //確認ボタン押下
    page.FindByButton("確認画面へ").Click()
    page.Screenshot("screenshot/ana-9.png")
    //一般の方押下
    page.FindByButton("一般の方").Click()
    page.Screenshot("screenshot/ana-10.png")

    //お客様情報入力
    page.FirstByXPath("/html/body/div[3]/div/div[1]/form/div[1]/table/tbody/tr/td[2]/input").Fill("ソラノ")
    page.FirstByXPath("/html/body/div[3]/div/div[1]/form/div[1]/table/tbody/tr/td[3]/input").Fill("タロウ")
    page.FirstByXPath("/html/body/div[3]/div/div[1]/form/div[1]/table/tbody/tr/td[4]/input").Fill("25")
    page.FirstByLabel("男性").Click()
    page.FirstByName("telNo").Fill("123-456-7890")
    page.FirstByName("assistMailAddress").Fill("test@test.test.test")
    page.FirstByName("assistConfirmMailAddress").Fill("test@test.test.test")
    page.Screenshot("screenshot/ana-11.png")
    //  page.FindByID("m_ticket02").Click()
    // log.Printf(page.Title())
}
```