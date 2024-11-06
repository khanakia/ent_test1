package gqlsa

import (
	"context"
	"fmt"
	"gqlsa/graph/generated"
	"gqlsa/graph/gqlsaresolver"
	"lace/gqlgin"
	"reflect"
	"saas/pkg/middleware/adminauthmiddleware"
	"saas/pkg/middleware/appmiddleware"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/ubgo/goutil"
	"github.com/vektah/gqlparser/v2/ast"
)

func setDirectives(c *generated.Config) {
	c.Directives.CanAdmin = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		cuser, _ := adminauthmiddleware.GetUserFromGqlCtx(ctx)
		if cuser == nil {
			return nil, fmt.Errorf("user not found")
		}

		if !cuser.Status {
			return nil, fmt.Errorf("access denied")
		}

		return next(ctx)
	}

	c.Directives.CanApp = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		cuser, _ := adminauthmiddleware.GetUserFromGqlCtx(ctx)
		if cuser == nil {
			return nil, fmt.Errorf("user not found")
		}

		if !cuser.Status {
			return nil, fmt.Errorf("access denied")
		}

		app, _ := appmiddleware.GetAppFromGqlCtx(ctx)
		if app == nil {
			return nil, fmt.Errorf("app not found")
		}

		if app.AdminUserID != cuser.ID {
			return nil, fmt.Errorf("app access denied")
		}

		// or let it pass through
		return next(ctx)
	}
}

func Boot(ginEngine *gin.Engine, resolver *gqlsaresolver.Resolver) {
	prefix := "/sa"
	rg := ginEngine.Group(prefix)

	c := generated.Config{Resolvers: resolver}
	setDirectives(&c)

	schm := generated.NewExecutableSchema(c)
	// extendSchema(&schm)
	gserver := handler.NewDefaultServer(schm)

	gqlgin.New(gqlgin.Config{
		RouterGroup:      rg,
		GqlServer:        gserver,
		PlaygroundKey:    resolver.AppConfig.Graphql.Key,
		RouteGroupPrefix: prefix,
		Middleware: []gin.HandlerFunc{
			adminauthmiddleware.MiddlewareSilent(resolver.Plugin.EntDB.Client()),
			appmiddleware.MiddlewareSilent(resolver.Plugin.EntDB.Client()),
		},
	})
	fmt.Println("boot gqlsa")
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
