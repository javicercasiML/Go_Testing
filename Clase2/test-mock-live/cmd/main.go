package main

import (
	"database/sql"
	"fmt"
	"testmock/internal/items"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// instances
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/item_db?parseTime=true")
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// app
	st := items.NewStorageSQL(db)
	sv := items.NewService(st)

	// -> get
	total, err := sv.GetTotalByName("Pepsi", 5)
	if err != nil {
		panic(err)
	}
	fmt.Println("item pepsi price:", total)


	var weight float64 = 34
	var price float64 = 200
	i, err := sv.UpdateByName("Pepsi", &weight, &price, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("item pepsi updated: %+v", i)
}