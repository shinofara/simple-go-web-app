GolangでWAFを使わずシンプルにWEBアプリケーションを構築できないか検討してみるリポジトリ
===========================

# Useing Pacages

router
github.com/julienschmidt/httprouter

middleware
github.com/urfave/negroni

context
github.com/nbio/httpcontext

render
github.com/unrolled/render

logger
github.com/uber-go/zap

dbはこれかな
https://github.com/go-gorp/gorp

# Setup

```
$ make setup
```

# Test Run

Execute this command in your terminal.

```
$ make run
```

Open this url in your browser.
http://localhost:8080/

# HTTP2

```
$ go run /usr/local/go/src/crypto/tls/generate_cert.go --host localhost
2017/01/17 18:08:33 written cert.pem
2017/01/17 18:08:33 written key.pem
```

リポジトリCRUDを利用してentiryを返却
model(service）、トランザクション処理なども含めて一つの処理の流れを担当
handler(controller)、routeとしての受けて、各modelからの処理の結果をHTMLや、JSONなどで返却を担当する
handlersはentiryを知らない
servideはrepositoryを知らない