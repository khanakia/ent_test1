package main

import (
	"fmt"
	"log"
	"os"
	"saas/pkg/constants"
	"strings"

	_ "github.com/lib/pq"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	var err error

	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaPath("../gqlsa/graph/ent.graphql"),
		entgql.WithConfigPath("../gqlsa/gqlgen.yml"),

		// add @canAdmin directive to the node and nodes query in ent.graphql
		// https://github.com/ent/ent/issues/3173
		entgql.WithSchemaHook(func(_ *gen.Graph, s *ast.Schema) error {
			for _, name := range []string{"node", "nodes"} {
				f := s.Types["Query"].Fields.ForName(name)
				if f == nil {
					return fmt.Errorf("missing query field %q", name)
				}
				f.Directives = append(f.Directives, &ast.Directive{Name: constants.DirectiveCanApp})
			}
			return nil
		}),

		entgql.WithOutputWriter(outputWriter),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		// entc.TemplateDir("./tmpl"),
		entc.FeatureNames("sql/execquery", "intercept", "schema/snapshot", "sql/modifier", "sql/upsert"),
		entc.Extensions(ex),
	}

	err = entc.Generate("./schema",
		&gen.Config{
			Target:  "./gen/ent",
			Package: "saas/gen/ent",
		},
		opts...,
	)
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}

func outputWriter(schema *ast.Schema) error {
	// generatePublicGql(schema)
	os.WriteFile("../gqlsa/graph/ent.graphql", []byte(printSchema(schema)), 0644)

	// should be after gqlsa generated as schema is pointer so publicgql conditions will affect
	// gqlsa ent.graphql file

	return nil
}

// we want also the copy the ent.graphql to gql not only gqlsa but we need only speicific query
// which can be publicly queried and ignore all other queries so we created an empty directory
// appGql with which we can know this query need to be in gql/ent.graphql file and skip other queries
func generatePublicGql(schema *ast.Schema) string {

	ps := NewPublicSchema(*schema)
	ps.Build()

	// s := &ast.Schema{
	// 	Directives: make(map[string]*ast.DirectiveDefinition),
	// }
	// var queryFields ast.FieldList

	// for name, definition := range schema.Types {
	// 	// fmt.Println(name, definition.Kind)

	// 	if name == entgql.QueryType {
	// 		// var fieldsNew ast.FieldList
	// 		for _, fieldDef := range definition.Fields {
	// 			// fmt.Println("--", fieldDef.Name)
	// 			// continue
	// 			skip := true

	// 			for _, directive := range fieldDef.Directives {
	// 				// fmt.Println(directive.Name)
	// 				if directive.Name == constants.DirectiveAppGql {
	// 					skip = false
	// 					break
	// 				}
	// 			}

	// 			if skip {
	// 				continue
	// 			}

	// 			// remove the @appGql directive as that we need only to identify the query so it's not actually a directive
	// 			fieldDef.Directives = fieldDef.Directives[:0]

	// 			// fieldsNew = append(fieldsNew, field)

	// 			queryFields = append(queryFields, fieldDef)
	// 		}

	// 		// definition.Fields = fieldsNew

	// 		// fmt.Println(definition.Fields)
	// 	}

	// }

	// s.AddTypes(&ast.Definition{
	// 	Name:   entgql.QueryType,
	// 	Kind:   ast.Object,
	// 	Fields: queryFields,
	// })

	// s = addTypeToSchema(s, &schema)

	sb := &strings.Builder{}
	formatter.
		NewFormatter(sb, formatter.WithIndent("  ")).
		FormatSchema(ps.Schema())

	os.WriteFile("../gql/graph/ent.graphql", []byte(sb.String()), 0644)
	return sb.String()
}

func printSchema(schema *ast.Schema) string {
	// updateSchema(schema)

	sb := &strings.Builder{}
	formatter.
		NewFormatter(sb, formatter.WithIndent("  ")).
		FormatSchema(schema)
	return sb.String()
}

func NewPublicSchema(schema ast.Schema) *PublicSchema {
	return &PublicSchema{
		schema: schema,
		s: &ast.Schema{
			Directives: schema.Directives,
		},
		defTypeNames: DefTypeNames{entgql.QueryType: true},
		// finalDefNamesBucket: DefTypeNames{},
		typesMap: map[string]*ast.Definition{},
		// defsMap:  map[string]*ast.Definition{},
		filterFieldMiddlewares: []FilterFieldMiddleware{FilterFieldMiddlewarePublicSkip(), FilterFieldMiddlewareRemoveDirective()},
		queryFieldMiddlewares:  []FilterFieldMiddleware{FilterFieldMiddlewareCanApp()},
	}
}

type DefTypeNames map[string]bool

// filter some fields e.g. PostType { postTypeForm: PostTypeForm} we do not need this in public api or
// skip some sensitive fields
type FilterFieldMiddleware func(fieldDef *ast.FieldDefinition) (*ast.FieldDefinition, error)

type PublicSchema struct {
	schema      ast.Schema
	s           *ast.Schema
	queryFields ast.FieldList
	// defsMap     map[string]*ast.Definition

	// definations to keep in new schema
	defTypeNames DefTypeNames
	// finalDefNamesBucket DefTypeNames

	typesMap map[string]*ast.Definition

	filterFieldMiddlewares []FilterFieldMiddleware
	queryFieldMiddlewares  []FilterFieldMiddleware
}

func (cls *PublicSchema) Build() {
	cls.buildQueryFields()
	cls.buildSchemaTypesMap()

	// extract and build schema for the QueryType Definition
	def := cls.typesMap[entgql.QueryType]
	cls.buildDef(def)
	// cls.extractDefTypeNames(cls.queryFields)
	// cls.extractDefTypeNamesForDefNames()
	// cls.extractDefTypeNamesForDefNames()

	cls.s.AddTypes(&ast.Definition{
		Name:   entgql.QueryType,
		Kind:   ast.Object,
		Fields: cls.queryFields,
	})

	cls.filterFields()
	// goutil.PrintToJSON(cls.defTypeNames)

	// goutil.PrintToJSON(cls.finalDefNamesBucket)
}

func (cls *PublicSchema) Schema() *ast.Schema {
	return cls.s
}

func (cls *PublicSchema) buildSchemaTypesMap() {
	for name, definition := range cls.schema.Types {
		// if name == "PostType" {
		// 	fmt.Printf("%+v", definition)
		// }
		cls.typesMap[name] = definition
	}

	// override the Query with filtered queryFields
	cls.typesMap[entgql.QueryType] = &ast.Definition{
		Name:   entgql.QueryType,
		Kind:   ast.Object,
		Fields: cls.queryFields,
	}
}

func (cls *PublicSchema) buildDef(def *ast.Definition) {
	// extract interfaces e.g. Node and add to new schema type
	for _, interfaceName := range def.Interfaces {
		defType, ok := cls.typesMap[interfaceName]
		_, ok1 := cls.defTypeNames[interfaceName]
		if ok && !ok1 {
			cls.s.AddTypes(defType)
			cls.defTypeNames[interfaceName] = true
		}
	}

	for _, fieldDef := range def.Fields {
		name := fieldDef.Type.Name()
		// fmt.Println(name)
		defType, ok := cls.typesMap[name]
		_, ok1 := cls.defTypeNames[name]
		if ok && !ok1 {
			// fmt.Println(name)
			cls.defTypeNames[name] = true
			cls.s.AddTypes(defType)
			cls.buildDef(defType)
			// cls.defTypeNames[name] = true
		}

		for _, argDef := range fieldDef.Arguments {
			name := argDef.Type.Name()

			defType, ok := cls.typesMap[name]
			_, ok1 := cls.defTypeNames[name]

			if ok && !ok1 {
				cls.defTypeNames[name] = true
				cls.s.AddTypes(defType)
				cls.buildDef(defType)
				// cls.defTypeNames[name] = true
			}
		}
	}

}

func FilterFieldMiddlewarePublicSkip() FilterFieldMiddleware {
	return func(fieldDef *ast.FieldDefinition) (*ast.FieldDefinition, error) {
		// for _, directive := range fieldDef.Directives {
		// 	if directive.Name == constants.DirectivePublicSkip {
		// 		return fieldDef, fmt.Errorf("skip")
		// 	}
		// }

		return fieldDef, nil
	}
}

func FilterFieldMiddlewareCanApp() FilterFieldMiddleware {
	return func(fieldDef *ast.FieldDefinition) (*ast.FieldDefinition, error) {
		for _, directive := range fieldDef.Directives {
			if directive.Name == constants.DirectiveAppGql {
				return fieldDef, nil
			}
		}

		return fieldDef, fmt.Errorf("skip")
	}
}

func FilterFieldMiddlewareRemoveDirective() FilterFieldMiddleware {
	return func(fieldDef *ast.FieldDefinition) (*ast.FieldDefinition, error) {
		var directivesNew ast.DirectiveList
		// for _, directive := range fieldDef.Directives {
		// 	switch directive.Name {
		// 	case constants.DirectiveAppGql, constants.DirectiveCanApp, constants.DirectivePublicSkip:
		// 		continue

		// 	default:
		// 		directivesNew = append(directivesNew, directive)
		// 	}
		// }

		fieldDef.Directives = directivesNew

		return fieldDef, nil
	}
}

func (cls *PublicSchema) filterFields() {
	for _, definition := range cls.s.Types {

		var fieldsNew ast.FieldList
		for _, fieldDef := range definition.Fields {

			var err error
			for _, fn := range cls.filterFieldMiddlewares {
				fieldDef, err = fn(fieldDef)
				if err != nil {
					break
				}
			}

			if err != nil {
				continue
			}
			// for _, directive := range fieldDef.Directives {
			// 	if directive.Name == constants.DirectivePublicSkip {
			// 		break
			// 	}
			// }

			// remove the @appGql directive as that we need only to identify the query so it's not actually a directive
			// fieldDef.Directives = fieldDef.Directives[:0]

			// cls.queryFields = append(cls.queryFields, fieldDef)
			fieldsNew = append(fieldsNew, fieldDef)
		}

		definition.Fields = fieldsNew
		// fmt.Println(definition.Fields)

	}
}

func (cls *PublicSchema) buildQueryFields() {
	for name, definition := range cls.schema.Types {
		if name == entgql.QueryType {
			for _, fieldDef := range definition.Fields {
				// fmt.Println("--", fieldDef.Name)
				// skip := true

				// for _, directive := range fieldDef.Directives {
				// 	if directive.Name == constants.DirectiveAppGql {
				// 		skip = false
				// 		break
				// 	}
				// }

				// if skip {
				// 	continue
				// }

				var err error
				for _, fn := range cls.queryFieldMiddlewares {
					fieldDef, err = fn(fieldDef)
					if err != nil {
						break
					}
				}

				if err != nil {
					continue
				}

				// remove the @appGql directive as that we need only to identify the query so it's not actually a directive
				// fieldDef.Directives = fieldDef.Directives[:0]

				cls.queryFields = append(cls.queryFields, fieldDef)
			}

		}
	}
}
