# notify_slack_of_amplify_events

## 概要

EventBridgeから転送されたAmplifyイベントに情報を追加し，Slackに転送します．

Goは開発者に実装方法を強制させられるため，可読しやすく，後続者が開発しやすいです．

また，静的解析のルールが厳しいため，バグを事前に見つけることができます．

そのため，実装がスケーリングしやすく，特にバグが許されないような基盤部分に適しています．

今回，これらには当てはまりませんが，Lambdaでは実装からデプロイまでの一連の流れを俯瞰できるため，Goを採用しました．

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
