package main

import (
	"log"

	_ "github.com/lib/pq"

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
