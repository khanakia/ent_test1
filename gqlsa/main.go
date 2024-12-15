package gqlsa

import (
	"fmt"
	"gqlsa/graph/generated"
	"gqlsa/graph/gqlsaresolver"
	"lace/gqlgin"
	"reflect"
	"saas/pkg/middleware/adminauthmiddleware"
	"saas/pkg/middleware/appmiddleware"
	"saas/pkg/middleware/gqldirectives"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-gonic/gin"
	"github.com/ubgo/goutil"
	"github.com/vektah/gqlparser/v2/ast"
)

func Boot(ginEngine *gin.Engine, resolver *gqlsaresolver.Resolver) {
	prefix := "/sa"
	rg := ginEngine.Group(prefix)

	c := generated.Config{Resolvers: resolver}
	setDirectives(&c)

	schm := generated.NewExecutableSchema(c)
	// extendSchema(&schm)
	gserver := handler.New(schm)
	gserver.AddTransport(transport.Options{})
	gserver.AddTransport(transport.GET{})
	gserver.AddTransport(transport.POST{})

	gqlgin.New(gqlgin.Config{
		RouterGroup:      rg,
		GqlServer:        gserver,
		PlaygroundKey:    resolver.AppConfig.Graphql.Key,
		RouteGroupPrefix: prefix,
		Middleware: []gin.HandlerFunc{
			adminauthmiddleware.MiddlewareApiKeySilent(resolver.Plugin.EntDB.Client()),
			adminauthmiddleware.MiddlewareSilent(resolver.Plugin.EntDB.Client()),
			appmiddleware.MiddlewareSilent(resolver.Plugin.EntDB.Client()),
		},
	})
	fmt.Println("boot gqlsa")
}

func setDirectives(c *generated.Config) {
	c.Directives.CanAdmin = gqldirectives.CanAdminDirective

	c.Directives.CanApp = gqldirectives.CanAppDirective
}

// Custom schema extension
func extendSchema(schema *graphql.ExecutableSchema) {
	// fields, err := fetchDynamicFields()
	// if err != nil {
	// 	log.Fatalf("failed to fetch dynamic fields: %v", err)
	// }

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
