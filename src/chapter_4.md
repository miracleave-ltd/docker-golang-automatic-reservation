# 4. ANAの予約ページを自動化してみる

## 4.1. `main.go`ファイルを編集し、ANAのページで予約出来るようコードを修正

前頁で編集したプログラムに以下コードを追加します。

```go
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
```

## 4.2. FindByIDのhoge1を書き換える

実際の要素IDを取得し、前項の`page.FindByID("hoge1").Click()`を書き換えます。

* https://www.ana.co.jp/ja/jp

上記URLにアクセスし、検索ボタンを押下すると、区間ページへ移動できます。
該当ページにてANAの区間検索ページ要素をChromeの検証ツールで特定します。

![キャプチャ](https://user-images.githubusercontent.com/66953939/84682510-b240a380-af70-11ea-9aaa-381d7f67df82.png)
```
page.FindByID("hoge1").Click()
↓
page.FindByID("buttonOneWay").Click()
```

続けて以下コードを追加します。

```go
//到着地「札幌」選択
page.FindByID("arrivalAirport").Select("札幌(千歳)")
page.Screenshot("screenshot/ana-3.png")
//カレンダーテキスト押下
page.FindByID("outwardEmbarkationDate").Click()
page.Screenshot("screenshot/ana-4.png")
//カレンダーでXpathを指定して8月10日を指定
page.FirstByXPath("hoge2").Click()
page.Screenshot("screenshot/ana-5.png")
//最安値指定
page.FirstByLabel("最安運賃を検索").Click()
page.Screenshot("screenshot/ana-6.png")
//検索ボタン押下
page.FirstByXPath("/html/body/div[4]/div/div[1]/form/div[2]/div[4]/p/input").Click()
page.Screenshot("screenshot/ana-7.png")
//値段を押下
page.FirstByLabel("hoge3").Click()
page.Screenshot("screenshot/ana-8.png")
//確認ボタン押下
page.FindByButton("確認画面へ").Click()
page.Screenshot("screenshot/ana-9.png")
//一般の方押下
page.FindByButton("一般の方").Click()
page.Screenshot("screenshot/ana-10.png")
```

## 4.3. FirstByXPathのhoge2を書き換える

前項と同ページにて**往路搭乗日**内のカレンダーをクリック、カレンダーが表示された状態で検証ツールを使い、搭乗日の`Xpath`を特定します。<br>※下図の①→②→③の順番で取得します。
![キャプチャ](https://user-images.githubusercontent.com/66953939/84685535-6f34ff00-af75-11ea-8d6e-4ff2b8d3893e.png)

下記の通り、`hoge2`部分を書き換えます。
```go
page.FirstByXPath("hoge2").Click()
// ↓ 取得したXpathを設定
page.FirstByXPath("/html/body/div[9]/div/div/div/div/div[3]/table/tbody/tr[3]/td[2]/a").Click()
```

## 4.4. FirstByLabelのhoge3を書き換える

値段のラベルを取得し"hoge3"を書き換えます。<br>**※上記ページ内で出発地、到着地、搭乗日を選択し検索すると表示されます。** 

![a](https://user-images.githubusercontent.com/66953939/84686408-d8694200-af76-11ea-8e3b-551d611ba86d.png)

下記の通り、`hoge3`部分を書き換えます。

```go
page.FirstByLabel("hoge3").Click()
// ↓ 取得したラベルを設定
page.FirstByLabel("23,960円").Click()
```

`hoge3`を書き換えたら以下を追加します。

```go
//お客様情報入力
page.FirstByXPath("/html/body/div[3]/div/div[1]/form/div[1]/table/tbody/tr/td[2]/input").Fill("ソラノ")
page.FirstByXPath("/html/body/div[3]/div/div[1]/form/div[1]/table/tbody/tr/td[3]/input").Fill("タロウ")
page.FirstByXPath("/html/body/div[3]/div/div[1]/form/div[1]/table/tbody/tr/td[4]/input").Fill("25")
page.FirstByLabel("男性").Click()
page.FirstByName("hoge4").Fill()
page.FirstByName("assistMailAddress").Fill("test@test.test.test")
page.FirstByName("assistConfirmMailAddress").Fill("test@test.test.test")
page.Screenshot("screenshot/ana-11.png")
```

## 4.5. FirstByNameのhoge4を書き換える

電話番号のName属性を取得し、`hoge4`を書き換えます。

```go
page.FirstByName("hoge4").Fill()
// ↓ 
page.FirstByName("telNo").Fill("123-456-7890")
```

ここまで完了したら、再びコンテナに入り、`Go`を実行します。

## 4.6. 完成したプログラムの動作確認

コンテナで以下を実行します。

```go
go run main.go
```

`screenshot`フォルダ内に`ana-11.png`までのキャプチャが撮れていることを確認します。
ここまで出来たら完成です。
厳密に組む場合はもっと`if`と`err`は駆使することになると思いますが、どちらかというと`Selenuim`の側面が強い感じとなりました。お疲れ様でした。

## 4.7. 付録

完成したコードはの全量は[こちら](./code.md)になります。


