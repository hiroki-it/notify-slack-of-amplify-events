# notify_slack_of_amplify_events

## 概要

EventBridgeから転送されたAmplifyイベントに情報を追加し，Slackに転送します．

## 環境構築

### 1. ビルド & コンテナ構築

```sh
$ docker-compose up -d
````

### 2. 接続

```sh
$ docker exec -it notify-slack-of-amplify-events sh 
```

## ホットリロード

ツールとして，[air](https://github.com/cosmtrek/air) を使用いたしました．

ホスト側でソースコードを修正すると，コンテナ側のアーティファクトが随時更新されます．

また，ホットリロードの実行時に，合わせてソースコード整形と静的解析を実行します．

```sh
# 接続
$ docker exec -it notify-slack-of-amplify-events sh 

# コンテナ内でairを起動
$ air -c .air.toml
```
