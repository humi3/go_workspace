# go_workspace

## go test repository

試験的なものを試すリポジトリー

## リポジトリーのクローンに関して

gopath の直下に以下のように作成する。  
`$GOPATH/src/github.com/clone対象のRepositoryのgithubユーザ/`

このリポジトリーは以下  
`$GOPATH/src/github.com/humi3/`

\$GOPATH については、mac だった以下になっているはず
`/Users/ユーザ/go`

## package について

他の自作パッケージを import する際には、以下のように読み込む。
`github.com/ユーザ/プロジェクト/対象のpackage`

## 参考資料

- [他言語から来た人が Go を使い始めてすぐハマったこととその答え](https://qiita.com/mumoshu/items/0d2f2a13c6e9fc8da2a4)
- [Go でオブジェクト指向を学ぶ](https://qiita.com/__init__/items/5f6c71eafd2d5e8ccb39)

## DB セットアップ

- `docker-compose up -d`
- `./init-mysql.sh`

### sh に権限がない場合

- `chmod +x 権限をつけたいファイル名`

## orm

gorm を使用する

### install

`go get github.com/jinzhu/gorm`

## config ファイル読み込みで使用

`go get gopkg.in/ini.v1`
