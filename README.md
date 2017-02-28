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

# Useing Pacages

| 役割       | パッケージ                           |
| ---------- | ------------------------------------ |
| router     | github.com/pressly/chi               |
| middleware | github.com/pressly/chi               |
| render     | github.com/unrolled/render           |
| logger     | github.com/uber-go/zap               |
| gorp       | https://github.com/go-gorp/gorp      |

# レイヤー設計

| レイヤ                    | 別名                 | 説明                                                     | 備考                |
| ------------------------- | -------------------- | -------------------------------------------------------- | ------------------- |
| ユーザインターフェース層  | プレゼンテーション層 | ユーザが触る場所                                         | HTML/JS/CSS         |
| アプリケーション層        | -                    | Requestを受けて下位層の結果をviewなどを使ってユーザに返す| Application Service |
| ドメイン層                | モデル層             | ビジネスロジック                                         | -                   |
| インフラストラクチャ層    | -                    | 上位レイヤを支えるインフラとの橋渡し                     | ORマッパなど        |


# sub-package設計

思考的には完全ではないけどDDDよりの思考で設計
http://qiita.com/haazime/items/6119097071149a362f7f
https://www.ogis-ri.co.jp/otc/hiroba/technical/DDDEssence/chap2.html#Repositories
違ったらPRください

| name        | 担当する世界                                                                      | その他                                                       |
| ----------- | ----------------------------------------------------------------------------------| ------------------------------------------------------------ |
| handler     | 発生したHTTPリクエストに対して、パス毎の処理を定義する場所                        |                                                              |
| service     | 目的毎の処理を各場所、主に`service`がrepositotyなど連絡を行い一つの処理を達成する | ユビキタス言語とすること（例: User.Register = 登録する)      |
| context     | リクエスト発生から返却までの間、維持したい情報を保持                              |                                                              |
| repository  | 各データストアとのやり取りを行い、適したentityを返却                              |                                                              |
| entity      | データストアのデータ構成                                                          | Identityを持っており値は変わっても同じものと扱う物。例えば人 |
| valueObject | `entity` とは逆に「色」とか「量」とかの用に不変のオブジェクトを管理               | 現在は未使用                                                 |
| middleware  | middleware                                                                        |                                                              |
| render      | HTTPレスポンスを返却する際に、htmlやjsonなどの描画を担当                          |                                                              |
| transfer    | データの転送（メール、ファイルUPLOADなどなど）                                    |                                                              |
| public      | HTML/JS/CSSなど                                                                   |                                                              |

# Contribution

1. Fork it
2. Clone to your local gopath (git clone git@github.com:<YOUR NAME>/simple-go-web-app.git $GOPATH/src/github.com/shinofara/simple-go-web-app)
3. Create your feature branch (git checkout -b my-new-feature)
4. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request