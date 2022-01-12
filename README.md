# README

## Install

- docker-compose up`でコンテナが起動します
- コンテナのashを起動し、`go run server.go`でgraphqlのAPIサーバが起動します


## Tips

- ashを起動する

`bash`の代わりに`ash`が起動できる。alpineのデフォルトのターミナルは`ash`

```
docker exec -it go_container ash
```