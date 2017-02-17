GolangでWAFを使わずシンプルにWEBアプリケーションを構築できないか検討してみるリポジトリ
===========================

# Usage

## Setup

```
$ make setup
```

## Test Run

Execute this command in your terminal.

```
$ make run
```

Open this url in your browser.
http://localhost:8080/

## Check HTTP2

```
$ go run /usr/local/go/src/crypto/tls/generate_cert.go --host localhost
2017/01/17 18:08:33 written cert.pem
2017/01/17 18:08:33 written key.pem
```


# Useing Pacages

| 役割       | パッケージ                            |
| ---------- | ------------------------------------ |
| router     | github.com/julienschmidt/httprouter  |
| middleware | github.com/urfave/negroni            |
| render     |  github.com/unrolled/render          |
| logger     | github.com/uber-go/zap               |
| gorp       | https://github.com/go-gorp/gorp      |




# sub-package設計

| name    | 担当する世界                                                                         | その他 |
| -------- | ------ | -- |
| handler | 発生したHTTPリクエストに対して、パス毎の処理を定義する場所                        | |
| service | 目的毎の処理を各場所、主に`service`がrepositotyなど連絡を行い一つの処理を達成する | ユビキタス言語とすること（例: User.Register = 登録する) |
| context | リクエスト発生から返却までの間、維持したい情報を保持                              |
| repository | 各データストアとのやり取りを行い、適したentityを返却 |
| entity | データストアのデータ構成 |
| middleware | middleware |
| render | HTTPレスポンスを返却する際に、htmlやjsonなどの描画を担当 |

