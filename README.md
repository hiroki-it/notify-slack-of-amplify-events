# notify_slack_of_amplify_events

## 概要

EventBridgeから転送されたAmplifyイベントに情報を追加し，Slackに転送します．

## 環境構築

### 1. ビルド

```sh
$ docker build --file ./build/Dockerfile --tag <イメージ名>:<タグ> .
```

### 2. コンテナ構築

```sh
$ docker run -d -it --name <コンテナ名> <イメージ名>:<タグ> sh
```

### 3. 接続

```sh
$ docker exec -it <コンテナ名> sh 
```
