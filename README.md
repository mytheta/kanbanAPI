# kanbanAPI
kanbanのAPI
## Swaggerの設定(自分用)
`gin-swagger`ってのを使った．

`swag`のコマンドインストール
```
go get github.com/swaggo/swag/cmd/swag
```
これも
```
go get -u github.com/swaggo/http-swagger
```
`sagger`の生成
```
swag init
```

[ここ](https://github.com/swaggo/gin-swagger)みてやった．

### swaggerを見て見る
`http://localhost:9000/swagger/index.html`