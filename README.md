# notify-slack-of-amplify-events

## 概要

### アプリケーションについて

EventBridgeから転送されたAmplifyイベントに情報を追加し，Slackに転送します．

### 採用言語について

Goを採用しました．

Goは開発者に実装方法を強制させられるため，可読性が高く，後続者が開発しやすいです．

また，静的解析のルールが厳しいため，バグを事前に見つけることができます．

そのため，実装がスケーリングしやすく，特にバグが許されないような基盤部分に適しています．

今回，これらには当てはまりませんが，Goを採用し，環境構築からデプロイまでの一連の流れを学習しようと考えました．

<br>

## Goのディレクトリ構造

本リポジトリのディレクトリ構造は以下の通りに構成しております．

クリーンアーキテクチャに則ったレイヤー名，パッケージ名，依存関係を取り入れております．

また，インターフェースは各レイヤーパッケージのルートに配置しました．

ただし，インターフェース層のプレゼンター，ユースケース層のアウトプットポートを廃止しております．

これに伴い，ユースケース層のインターラクターは，プレゼンターではなくレスポンスモデルを返却するようにしております．

```
project
├── build # ビルド処理
├── cmd   # エントリポイントを含む処理，ユニットテスト
|   ├── domain           # ** ドメイン層 **
|   |   ├── entities     # エンティティ
|   |   ├── ids          # ID
|   |   ├── repositories # インターフェースリポジトリ
|   |   └── values       # 値オブジェクト
|   |
|   ├── infrastructure      # ** インフラストラクチャ層 ** 
|   |   ├── logger          # ロガー 
|   |   └── foo             # 任意のルートエンティティ
|   |       └── reposiories # リポジトリ
|   |
|   ├── interaces           # ** インターフェース層 **
|   |   └── foo             # 任意のルートエンティティ
|   |       ├── controllers # コントローラ
|   |       ├── presenters  # プレゼンター
|   |       └── validators  # バリデーション
|   |     
|   └── usecase             # ** ユースケース層 **
|       └── foo             # 任意のルートエンティティ
|           ├── interactors # インターラクタ
|           ├── boundaries  # インプットバウンダリ
|           ├── requests    # リクエストモデル
|           ├── responses   # レスポンスモデル
|           └── services    # サービス
| 
├── integration_test # 統合テスト
|   ├── request      # テストデータ 
|   └── response     # ゴールデンファイル
| 
├── mock # モック
├── ops  # CICD用シェルスクリプト
└── serverless_configs # Serverless Frameworkの共通部品
```

<br>

## 環境構築

### RIEのインストール

ローカルで統合テストを行うために，RIEをインストールする必要があります．

RIEにより，ローカルで擬似的にLambdaを再現できます．

いくつか方法が用意されており，Goソースコードの再利用性の観点から，ホストPCにRIEをインストールする方法を採用いたしました．

ホストPCで以下のコマンドを実行します．

```shell
$ mkdir -p ~/.aws-lambda-rie
$ curl -Lo ~/.aws-lambda-rie/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie
$ chmod +x ~/.aws-lambda-rie/aws-lambda-rie
```

その他のインストール方法につきましては，以下のリンクを参考に．

https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/go-image.html#go-image-other

### ビルドイメージ & コンテナ構築

Dockerfileからイメージをビルドし，コンテナを構築します．

イメージのビルドにおいては，マルチステージビルドを採用しております．

```shell
$ docker-compose run -d --rm app
````

### モジュールのインストール

コンテナで，アプリケーションで使用するモジュールをインストールします．

```shell
$ docker-compose run --rm app go mod download -x
```

<br>

## ホットリロード

ホスト側のソースコードを修正に合わせて，コンテナ側のアーティファクトを随時更新します．

また，ホットリロードの実行時に，合わせてソースコード整形と静的解析を実行します．

ツールとして，[air](https://github.com/cosmtrek/air) を使用いたしました．

```shell
$ docker-compose run --rm app air -c .air.toml
```

<br>

## ユニットテスト

### テスト関数の命名

Roy Osherove氏の命名戦略を参考にいたしました．

参考リンク：https://osherove.com/blog/2005/4/3/naming-standards-for-unit-tests.html

### テスト設計

テーブル駆動テストを採用しました．

参考リンク：https://github.com/golang/go/wiki/TableDrivenTests

### ユニットテスト

Goサービスのユニットテストを実行します．

```shell
$ docker-compose run --rm app go test -v -cover ./cmd/...
```

<br>

## 統合テスト（ローカル環境のみ対応）

GoサービスからLambdaサービスにPOSTリクエストを送信し，一連の処理をテストします．

ローカル環境でLambdaを擬似的に再現するために，RIEを使用いたしました．

CircleCIにおけるジョブにて，Lambdaのホストを指定してリクエストを送信できないため，ローカル環境でのみ実行可能です．

なお，修正したソースコードをLambdaのRIEに再反映するためには，イメージを再ビルドする必要があります．

アタッチモードで起動中のlambdaサービスに対して，appサービスからリクエストを送信します．

```shell
$ docker-compose up --build lambda

# 別のターミナルを開いた上で実行する．
$ docker-compose run --rm app go test -v -cover ./integration_test/...
```

<br>

## CI/CD

### CI

以下を行います．

1. docker-compose.ymlの静的解析
2. コンテナ構築
3. パッケージのダウンロード
4. Goの静的解析
5. ユニットテスト
6. カバレッジレポートの作成
7. CircleCIのアーティファクトとして，カバレッジレポートをアップロード

カバレッジレポートは，CirlcleCIのアーティファクトタブで確認できます．

### CD

原則として，ローカル環境からソースコードをLambdaにデプロイしないようにします．

代わりに，CircleCIのCDにて，事前にイメージをECRにプッシュし，またソースをLambdaにデプロイします．

デプロイツールとして，[Serverless Framework](https://github.com/serverless/serverless) を使用いたしました．

<br>

## Amplifyについて

Amplifyを用いて，SSGアプリのCI/CDを構築します．

AmplifyがS3にSSGアプリをデプロイし，このイベントをEventBridgeがキャッチします．

AmplifyのCI/CDにつきましては，[こちらのリポジトリ](https://github.com/hiroki-it/deploy-ssg-to-amplify) を参考に．
