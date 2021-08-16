# 2. Dockerの準備

<!-- toc -->

## 2.1. プログラムをダウンロード

下図のとおり、GitHubから`docker-golang-automatic-reservation-master.zip`をダウンロードします。
ダウンロードしたらユーザー直下に解答し、配置します。

![キャプチャ](https://user-images.githubusercontent.com/66953939/84660963-94fddc00-af54-11ea-90ff-e8fcc0af3e5f.png)

![キャプチャ](https://user-images.githubusercontent.com/66953939/84664878-23289100-af5a-11ea-9818-8b3dd7beb8ca.png)

## 2.2. Dockerイメージのビルド

コマンドプロンプトで解凍した`docker-golang-automatic-reservation-master`まで移動します。

```sh
cd C:\Users\docker-golang-automatic-reservation-master
```

移動したら以下を実行します。

```sh
docker-compose build
```

## 2.3. コンテナの起動

ディレクトリはそのままで以下コマンドを実行します。

```sh
docker-compose up -d
```

更に以下のコマンドを実行しコンテナが正常に起動しているか確認します。

```sh
docker ps
# Creating docker-golang-automatic-reservation-master_app_1 ... done
# ↑のように「docker-golang-automatic-reservation-master_app_1」のコンテナがupしていればOK！
```