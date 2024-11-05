package bldr

import (
	"fmt"

	"github.com/uptrace/bun"
)

type GetParams struct {
	Model string
}

func Get(params GetParams, db *bun.DB) error {
	// sch := GetSchema()
	// model, err := sch.FindModel(params.Model)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(model)

	// db := &bun.DB{}
	// db1 := bun.NewDB(&sql.DB{}, pgdialect.New())

	qb := db.NewSelect().Table("users").Column("id").Where("id = ?", 123)

	fmt.Println(qb.String())

	return nil
}

func GetSchema() *Schema {
	return &Schema{
		Version: "v1",
		Models: []*Model{
			{
				Table:        "users",
				PrimaryKey:   "id",
				KeyType:      "string",
				Incrementing: true,
				Fields: []Field{
					{Name: "id", Type: FieldTypeString, Unique: true},
					{Name: "name", Type: FieldTypeString},
					{Name: "created_at", Type: FieldTypeDateTime},
					{Name: "updated_at", Type: FieldTypeDateTime},
				},
			},
		},
	}
}
