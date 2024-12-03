package schema

import (
	"context"
	"saas/gen/ent"

	"github.com/gosimple/slug"
)

// auto slugify before save
func slugMutateHook() func(next ent.Mutator) ent.Mutator {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if s, ok := m.(interface {
				SetSlug(string)
				Slug() (r string, exists bool)
			}); ok {
				slugv, _ := s.Slug()
				s.SetSlug(slug.Make(slugv))
			}
			v, err := next.Mutate(ctx, m)
			// Post mutation action.
			// fmt.Println("new value:", v)
			return v, err
		})
	}
}
