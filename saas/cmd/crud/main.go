package main

import (
	"context"
	"fmt"
	"saas/cmd/bldr"
	"saas/pkg/app"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func main() {
	app := app.New()
	// app.Cli.Execute()

	db := bun.NewDB(app.DB, pgdialect.New())

	ctx := context.Background()
	printNone(db, ctx)

	err := bldr.Get(bldr.GetParams{
		Model: "users",
	}, db)
	fmt.Println(err)
	// users := make([]User, 0)
	// err := db.NewRaw(
	// 	"SELECT id, email FROM ? LIMIT ?",
	// 	bun.Ident("users"), 100,
	// ).Scan(ctx, &users)

	// fmt.Println(err)
	// goutil.PrintToJSON(users)
}

type User struct {
	ID    string
	Email string
}

func printNone(v ...interface{}) {

}
