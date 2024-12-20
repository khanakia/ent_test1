package main

import (
	"embed"
	"os"
	"saas/pkg/entsaasmedia"
	"strings"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc/gen"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
)

//go:embed template/*
var _templates embed.FS

func parseT(path string) *gen.Template {
	return gen.MustParse(gen.NewTemplate(path).
		Funcs(entgql.TemplateFuncs).
		ParseFS(_templates, path))
}

func outputWriter(schema *ast.Schema) error {
	// generatePublicGql(schema)
	os.WriteFile("../gqlsa/graph/ent.graphql", []byte(printSchema(schema)), 0644)

	// should be after gqlsa generated as schema is pointer so publicgql conditions will affect
	// gqlsa ent.graphql file

	return nil
}

func printSchema(schema *ast.Schema) string {
	// updateSchema(schema)

	sb := &strings.Builder{}
	formatter.
		NewFormatter(sb, formatter.WithIndent("  ")).
		FormatSchema(schema)
	return sb.String()
}

func annotation(ants gen.Annotations) (*entgql.Annotation, error) {
	ant := &entgql.Annotation{}
	if ants != nil && ants[ant.Name()] != nil {
		if err := ant.Decode(ants[ant.Name()]); err != nil {
			return nil, err
		}
	}
	return ant, nil
}

func gqlTypeFromNode(t *gen.Type) (gqlType string, ant *entgql.Annotation, err error) {
	if ant, err = annotation(t.Annotations); err != nil {
		return
	}
	gqlType = t.Name
	if ant.Type != "" {
		gqlType = ant.Type
	}
	return
}

func buildTypes(g *gen.Graph, s *ast.Schema) error {
	/*
	 * Loop through each node
	 * Check if it has the SaasMedia annotation
	 * extract Input names [{CreatePostInput true} {UpdatePostInput false}]
	 * Now Find the those inputs in ast.Schema and add new Fields to those inputs
	 */
	for _, node := range g.Nodes {
		if node.HasCompositeID() {
			continue
		}

		if node.Annotations["SaasMedia"] == nil {
			continue
		}

		saasMedia, err := entsaasmedia.ExtractAnnotation(node.Annotations)

		if err != nil {
			continue
		}

		// goutil.PrintToJSON(saasMedia)

		_, ant, err := gqlTypeFromNode(node)
		if err != nil {
			return err
		}

		inputs, err := extractMutationInputs(node, ant)
		if err != nil {
			return err
		}

		// now we compare the input in schema and add the missing inputs
		addNewFieldsToSchema(saasMedia, inputs, s)
	}
	return nil
}

type MutationInputs struct {
	IsCreate bool
}

func extractMutationInputs(t *gen.Type, ant *entgql.Annotation) (map[string]MutationInputs, error) {
	inputs := make(map[string]MutationInputs)

	for _, i := range ant.MutationInputs {
		if i.IsCreate && ant.Skip.Is(entgql.SkipMutationCreateInput) {
			continue
		}
		if !i.IsCreate && ant.Skip.Is(entgql.SkipMutationUpdateInput) {
			continue
		}

		desc := entgql.MutationDescriptor{Type: t, IsCreate: i.IsCreate}
		name, err := desc.Input()
		if err != nil {
			return nil, err
		}

		inputs[name] = MutationInputs{
			IsCreate: i.IsCreate,
		}
	}

	return inputs, nil
}
func namedType(name string, nullable bool) *ast.Type {
	if nullable {
		return ast.NamedType(name, nil)
	}
	return ast.NonNullNamedType(name, nil)
}

func addNewFieldsToSchema(saasMedia *entsaasmedia.Annotation, inputs map[string]MutationInputs, s *ast.Schema) {
	for _, def := range s.Types {

		i, ok := inputs[def.Name]

		if ok {
			for _, mediaField := range saasMedia.Fields {
				switch {
				case i.IsCreate:
					def.Fields = append(def.Fields, &ast.FieldDefinition{
						Name: mediaField.GqlCreate,
						Type: namedType("[ID!]", true),
					})
				default:
					def.Fields = append(def.Fields, &ast.FieldDefinition{
						Name: mediaField.GqlAdd,
						Type: namedType("[ID!]", true),
					}, &ast.FieldDefinition{
						Name: mediaField.GqlRemove,
						Type: namedType("[ID!]", true),
					})
				}
				if !i.IsCreate {
					def.Fields = append(def.Fields, &ast.FieldDefinition{
						Name: mediaField.GqlClear,
						Type: namedType("Boolean", true),
					})
				}
			}

		}
	}
}
