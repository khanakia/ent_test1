package main

import (
	"fmt"
	"log"
	"saas/pkg/constants"

	_ "github.com/lib/pq"
	"github.com/vektah/gqlparser/v2/ast"

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
