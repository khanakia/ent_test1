package schema

import (
	"context"
	"fmt"
	"saas/pkg/entsaasmedia"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("slug").Optional(),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Post) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("app_id", "slug").Unique(),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		// BaseApp{},
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// MediaMutations(MediaInputOption{Name: "featured", BuilderField: "featuredmedias"}),
		entsaasmedia.MediaFields(entsaasmedia.FieldOption{Name: "featured"}, entsaasmedia.FieldOption{Name: "icon"}),
		entgql.Skip(entgql.SkipWhereInput),
		// entgql.RelayConnection(),
		entgql.QueryField("posts").Directives(),
		// entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Post) Hooks() []ent.Hook {
	return []ent.Hook{
		// slugMutateHook(),
	}
}

func slugMutateHook() func(next ent.Mutator) ent.Mutator {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			fmt.Printf("Type: %s, Operation: %s, ConcreteType: %T\n", m.Type(), m.Op(), m)

			if s, ok := m.(interface {
				FeaturedMediasIDs() (ids []string)
			}); ok {
				fmt.Println(s.FeaturedMediasIDs())
			}

			// v, err := next.Mutate(ctx, m)
			// Post mutation action.
			// fmt.Println("new value:", v)
			// return v, err

			// for _, s := range m.Fields() {
			// 	f, _ := m.Field(s)
			// 	fmt.Println(s, f)
			// 	// if f == nil {
			// 	// 	continue
			// 	// }

			// 	// TODO: Check if the field is annotated with ExpirableCookie
			// }
			return next.Mutate(ctx, m)

		})
	}
}

// type Annotation struct {
// 	// OrderField is the ordering field as defined in graphql schema.
// 	Inputs []MediaInputOption
// }

// // Name implements ent.Annotation interface.
// func (Annotation) Name() string {
// 	return "SaasMediaField"
// }

// type MediaInputOption struct {
// 	Name         string
// 	BuilderField string
// 	Unique       bool
// }

// func MediaMutations(inputs ...MediaInputOption) Annotation {
// 	if len(inputs) == 0 {
// 		inputs = []MediaInputOption{}
// 	}

// 	a := []MediaInputOption{}
// 	for _, f := range inputs {
// 		a = append(a, f)
// 	}
// 	return Annotation{Inputs: a}
// }
