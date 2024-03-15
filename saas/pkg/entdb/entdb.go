package entdb

import (
	"context"
	"database/sql"
	"fmt"
	"runtime"
	"saas/gen/ent"
	"strconv"
	"strings"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/fatih/color"
)

type EntDB struct {
	client *ent.Client
}

func (e EntDB) Client() *ent.Client {
	return e.client
}

func New(db *sql.DB) EntDB {
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	// client := ent.NewClient(ent.Driver(drv), ent.Debug())
	client := ent.NewClient(ent.Driver(&CustomDriver{drv}), ent.Debug())

	client.Intercept(
		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
			return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
				count, err := next.Query(ctx, query)
				return count, err
			})
		}),
	)

	// ent.InterceptFunc(func(next ent.Querier) ent.Querier {
	// 	return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
	// 		// Do something before the query execution.
	// 		value, err := next.Query(ctx, query)
	// 		// Do something after the query execution.
	// 		return value, err
	// 	})
	// })

	return EntDB{
		client: client,
	}
}

type CustomDriver struct {
	*entsql.Driver
}

func (d *CustomDriver) Query(ctx context.Context, query string, args, v any) error {
	err := d.Driver.Query(ctx, query, args, v)
	// fmt.Println("Custom ERROR", err)
	FileWithLineNum()
	return err
}

// FileWithLineNum return the file name and line number of the current file
func FileWithLineNum() string {
	// the second caller usually from gorm internal, so set i start from 2
	for i := 1; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.Contains(file, "/vendor") && !strings.Contains(file, "/ent") && !strings.Contains(file, "generated.go")) {
			// fmt.Println(file, line, ok)
			fmt.Println(time.Now().Format(time.RFC3339), color.GreenString(file))
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}

	return ""
}
