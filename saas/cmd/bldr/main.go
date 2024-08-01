package bldr

import "fmt"

func main() {

}

type Mutation struct {
	model *Model
}

type CreateInput struct {
	ResourceName string   `json:"resourceName"`
	KeyValues    KeyValue `json:"keyValues"`
}

type KeyValue map[string]interface{}

func getSchema() *Schema {
	return &Schema{
		Version: "v1",
		Models: []*Model{
			{
				Table:        "todos",
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

type Schema struct {
	Version string   `json:"version"`
	Models  []*Model `json:"models"`
}

type Model struct {
	Table        string  `json:"table"`
	PrimaryKey   string  `json:"primarKey"`
	KeyType      string  `json:"keyType"`
	Incrementing bool    `json:"incrementing"`
	Fields       []Field `json:"fields"`
}

func (cls *Model) FindField(name string) (*Field, error) {
	for _, field := range cls.Fields {
		if field.Name == name {
			return &field, nil
		}
	}
	return nil, fmt.Errorf("cant find model %s", name)
}

type Field struct {
	Name         string      `json:"name"`
	Type         FieldType   `json:"type"`
	Unique       bool        `json:"unique"`
	DefaultValue interface{} `json:"default"`
}

type FieldType string

const (
	FieldTypeInt      FieldType = "int"
	FieldTypeBigInt   FieldType = "bigint"
	FieldTypeString   FieldType = "string"
	FieldTypeDateTime FieldType = "datetime"
)

func (cls *Schema) FindModel(name string) (*Model, error) {
	for _, model := range cls.Models {
		if model.Table == name {
			return model, nil
		}
	}
	return nil, fmt.Errorf("cant find model %s", name)
}
