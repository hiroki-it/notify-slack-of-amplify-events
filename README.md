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
