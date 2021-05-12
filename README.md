# 概要
freeeの取引履歴のうち、売上高のみ取り出しCSV(sales.csvへ出力する)

# 事前インストール
## Go(1.16)のインストール
以下のリンクを参照
https://golang.org/doc/install

## GNU Makeのインストール
インストールされていれば、以下のコマンドを実行して確認可能
```
make --version
```

## ビルド用のバイナリインストール
```
make devel-deps
```

# ビルド
## バイナリのビルド
`./bin/gen-sales-csv`が生成される
```
make build
```

# 環境変数の設定
## .envファイルの作成
プロジェクト直下の`.env.template`をコピーし、`.env`を作成する。


## Client ID、Client Secretの取得
以下のリンクに従って行う。
https://developer.freee.co.jp/tutorials/starting-api
ただし、コールバックURLには`urn:ietf:wg:oauth:2.0:oob`を入力する。

取得した値を、`.env`の`CLIENT_ID`と`CLIENT_SECRET`へ入力する

## Tokenの取得
以下のリンクに従って行う。
https://developer.freee.co.jp/tutorials/getting-access-token

まずは認可コードを取得する。その後、以下のPostを実行する。
```
curl -i -X POST \
 -H "Content-Type:application/x-www-form-urlencoded" \
 -d "grant_type=authorization_code" \
 -d "client_id={クライアントId}" \
 -d "client_secret={クライアントSecret}" \
 -d "code={認可コード}" \
 -d "redirect_uri=urn%3Aietf%3Awg%3Aoauth%3A2.0%3Aoob" \
 'https://accounts.secure.freee.co.jp/public_api/token'
```
以下のようなレスポンスが返ってくる
```
{"access_token":"{アクセストークン}","token_type":"bearer","expires_in":86399,"refresh_token":"{リフレッシュトークン}","scope":"read write default_read","created_at":1620802518}
```

## Company Idの取得
https://developer.freee.co.jp/docs/accounting/reference#/
上記のリンクからアクセスして取得する。

`/api/1/companies`の鍵マークにアクセストークンを入力して実行する。

取得したら、`.env`の`COMPANY_ID`へ入力する

```
{
  "companies": [
    {
      "id": 1234567,
      "name": null,
      "name_kana": null,
      "display_name": "開発用テスト事業所",
      "role": "admin"
    }
  ]
}
```
# 実行
## 初回の実行(Tokenの入力が必要)
```
./bin/gen-sales-csv --access-token={アクセストークン} --refresh-token={リフレッシュトークン}
```
`./out/sales.csv`が生成される。また、トークンが`./token.b`へ保存される

## 2回目の実行
```
./bin/gen-sales-csv
```
`./out/sales.csv`が生成される。また、トークンの有効期限が切れた場合のリフレッシュは自動で行われ、`./token.b`へ保存される