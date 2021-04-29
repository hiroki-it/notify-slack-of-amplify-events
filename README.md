# notify-slack-of-amplify-events

## 概要

EventBridgeから転送されたAmplifyイベントに情報を追加し，Slackに転送します．

Goは開発者に実装方法を強制させられるため，可読性が高く，後続者が開発しやすいです．

また，静的解析のルールが厳しいため，バグを事前に見つけることができます．

そのため，実装がスケーリングしやすく，特にバグが許されないような基盤部分に適しています．

今回，これらには当てはまりませんが，Goの環境構築からデプロイまでの一連の流れを俯瞰するために，Goを採用しました．

## 環境構築

### 1. RIEのインストール

Lambdaをローカルで擬似的に再現するために，RIEをインストールする必要があります．

いくつか方法が用意されており，Goソースコードの再利用性の観点から，ホストPCにRIEをインストールする方法を採用いたしました．

ホストPCで以下のコマンドを実行します．

```shell
$ mkdir -p ~/.aws-lambda-rie
$ curl -Lo ~/.aws-lambda-rie/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie
$ chmod +x ~/.aws-lambda-rie/aws-lambda-rie
```

その他のインストール方法につきましては，以下を参考に．

https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/go-image.html#go-image-other

## ディレクトリ構造

本リポジトリのディレクトリ構造は以下の通りに構成しております．

cmdディレクトリの構成は，クリーンアーキテクチャを参考にいたしました．

参考：https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

```
notify-slack-of-amplify-events
├── build # ビルド処理
├── cmd # エントリポイントを含む処理
|   ├── controllers # コントローラ
|   ├── entities # エンティティ
|   └── usecases # ユースケース
|
├── configs # セットアップ処理
└── test
    ├── mock # モック処理
    └── unit # ユニットテスト処理
```

### 2. ビルドイメージ & コンテナ構築

Dockerfileからイメージをビルドし，コンテナを構築します．

デタッチモードを使用して起動します．

```shell
$ docker-compose up -d
````

### 3. 起動

コンテナを起動します．また，起動時に```shell```を実行し，コンテナに接続します．

また，```rm```オプションを使用して，処理後にコンテナを削除するようにします．

```shell
$ docker-compose run --rm  notify-slack-of-amplify-events sh

# 接続中です．
/go/src $
# 接続を切断します．
/go/src $ exit
```

### 4. モジュールのインストール

起動中のコンテナで，アプリケーションで使用するモジュールをインストールします．

```shell
$ docker-compose exec notify-slack-of-amplify-events go mod download -x
```

停止中のコンテナでこれを実行する場合は，```run```コマンドを使用します．

また，```rm```オプションを使用して，処理後にコンテナを削除するようにします．

```shell
$ docker-compose run --rm  notify-slack-of-amplify-events go mod download -x
```

## ホットリロード

ツールとして，[air](https://github.com/cosmtrek/air) を使用いたしました．

ホスト側でソースコードを修正すると，コンテナ側のアーティファクトが随時更新されます．

また，ホットリロードの実行時に，合わせてソースコード整形と静的解析を実行します．

```shell
$ docker-compose exec notify-slack-of-amplify-events air -c .air.toml
```

停止中のコンテナでこれを実行する場合は，```run```コマンドを使用します．

また，```rm```オプションを使用して，処理後にコンテナを削除するようにします．

```shell
$ docker-compose run --rm notify-slack-of-amplify-events air -c .air.toml
```
