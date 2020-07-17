package main

import (
	"encoding/json"
	"fmt"

	"github.com/boombaw/go-wilayah/route"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
	e := route.Init()
	data, err := json.MarshalIndent(e.Routes(), "", "	")
	if err != nil {
		panic(fmt.Sprint(err))
	}
	fmt.Println(string(data))
	e.Logger.Fatal(e.Start(":5000"))
}
