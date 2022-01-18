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

## tutorial

### graphqlでデータを取得する

apiのエンドポイント(http://localhost:8000/query)に以下のクエリを実行して、データを取得する

```
query findTodos {
  todos {
    text
    done

  }
}
```


### graphqlでデータを更新する

apiのエンドポイント(http://localhost:8000/query)に以下のクエリを実行して、データを保存する

```
mutation createTodo {
	createTodo(input:{ text: "test", userId:"1" }){
    id
    text
    done
    user {
      id
      name
    }
  }
}
```

### ディレクトリgraph以下のテストを実行する


```
go test -v ./graph
```

