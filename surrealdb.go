package schema

import (
	"fmt"

	"github.com/surrealdb/surrealdb.go"
)

var SurrealDBConn *surrealdb.DB

func init() {
	d, err := ConnectSurrealDB("ws://localhost:9000/rpc", "root", "root", "test", "test")
	if err != nil {
		panic(err)
	}
	SurrealDBConn = d
}

func ConnectSurrealDB(url string, user string, pass string, name string, namespace string) (*surrealdb.DB, error) {
	db, err := surrealdb.New(url)
	if err != nil {
		return nil, err
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": user,
		"pass": pass,
	}); err != nil {
		return nil, err
	}

	if _, err = db.Use(name, namespace); err != nil {
		return nil, err
	}

	return db, nil
}

func Select[T any](from string) (*[]T, error) {
	res, err := SurrealDBConn.Query(fmt.Sprintf("SELECT * FROM %s", from), nil)
	if err != nil {
		return nil, err
	}

	var data []T
	if ok, err := surrealdb.UnmarshalRaw(res, &data); !ok {
		return nil, err
	}

	return &data, nil
}

func SelectOnly[T any](from string) (*T, error) {
	res, err := SurrealDBConn.Query(fmt.Sprintf("SELECT * FROM %s", from), nil)
	if err != nil {
		return nil, err
	}

	var data []T
	if ok, err := surrealdb.UnmarshalRaw(res, &data); !ok {
		return nil, err
	}

	if len(data) == 1 {
		return &data[0], nil
	}

	return nil, fmt.Errorf("more than record returned")
}

func Graph[T any](from string) (*T, error) {
	res, err := SurrealDBConn.Query(fmt.Sprintf("SELECT *,fn::recurse($this) AS children FROM %s;", from), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	data := make([]T, 0)
	if ok, err := surrealdb.UnmarshalRaw(res, &data); !ok {
		return nil, err
	}

	return &data[0], nil
}
