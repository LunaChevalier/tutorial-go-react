package graph_test

import (
	"bytes"
	"context"
	"strings"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/LunaChevalier/tutorial-go-react/graph"
	"github.com/LunaChevalier/tutorial-go-react/graph/generated"
)

func TestTodos(t *testing.T)  {
	const dataSource = "root:root@tcp(mysql:3306)/react-go-app?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dataSource)

	ts := httptest.NewServer(
		handler.NewDefaultServer(
			generated.NewExecutableSchema(
				generated.Config{
					Resolvers: &graph.Resolver{DB: db},
				},
			),
		),
	)

	defer ts.Close()

	q := struct {
		Query string
	}{
		Query: "{todos{text}}",
	}
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, ts.URL, &body)
	if err != nil {
		t.Fatal("error new request", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err!= nil {
		t.Fatal("error request", err)
	}

	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("error read body", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal("error request code:", res.StatusCode, string(result))
	}

	// TODO: graphqlでのjsonをパースするように修正する
	assert.Equal(t, strings.Contains(string(result), "data"), true)
	assert.Equal(t, strings.Contains(string(result), "todos"), true)
	assert.Equal(t, strings.Contains(string(result), "text"), true)
	assert.Equal(t, strings.Contains(string(result), "id"), false)
	// t.Log(string(result))
}