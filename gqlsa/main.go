package gqlsa

import (
	"fmt"
	"gqlsa/graph/generated"
	"gqlsa/graph/gqlsaresolver"
	"gqlsa/internal/gqlgin"
	"lace/gqlgenfn"
	"reflect"
	"saas/pkg/middleware"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/ubgo/goutil"
	"github.com/vektah/gqlparser/v2/ast"
)

func Boot(ginEngine *gin.Engine, resolver *gqlsaresolver.Resolver) {
	ginEngine.Use(middleware.MiddlewareSilent(resolver.Plugin.EntDB.Client()))
	ginEngine.Use(gqlgenfn.GinContextToContextMiddleware())

	c := generated.Config{Resolvers: resolver}

	schm := generated.NewExecutableSchema(c)
	// extendSchema(&schm)
	gserver := handler.NewDefaultServer(schm)

	gqlgin.New(gqlgin.Config{
		GinEngine:        ginEngine,
		GqlServer:        gserver,
		PlaygroundKey:    resolver.AppConfig.Graphql.Key,
		RouteGroupPrefix: "/sa/",
	})
	fmt.Println("boot gqlsa")
}

// Custom schema extension
func extendSchema(schema *graphql.ExecutableSchema) {
	// fields, err := fetchDynamicFields()
	// if err != nil {
	// 	log.Fatalf("failed to fetch dynamic fields: %v", err)
	// }

	fmt.Println("sfsdhjfkdshjfkhsdjkhds")

	aa := *schema
	// fmt.Println(aa.Schema().Query.Fields)
	for _, field := range aa.Schema().Types {
		if field.Name == "Student" {

			fmt.Println(field.Name, reflect.TypeOf(field))

			field.Fields = append(field.Fields, &ast.FieldDefinition{
				Name: "class",
				Type: ast.NamedType("String", nil),
			})

			goutil.PrintToJSON(field.Fields[0])
		}

	}
	// // Add dynamic fields to the ExampleType
	// for _, field := range fields {
	// 	fieldName := field.FieldName
	// 	fieldType := field.FieldType

	// 	// Add field to introspection
	// 	schema.Query().Fields()["example"].Type.(*graphql.Object).AddFieldConfig(fieldName, &graphql.Field{
	// 		Type: graphql.String, // Adjust the type based on fieldType
	// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// 			example := p.Source.(*ExampleType)
	// 			if value, ok := example.DynamicFields[fieldName]; ok {
	// 				return value, nil
	// 			}
	// 			return nil, fmt.Errorf("field %s not found", fieldName)
	// 		},
	// 	})
	// }
}
