package gqlsaresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"fmt"
	"gqlsa/graph/generated"
	"saas/gen/ent"
	"saas/gen/ent/postcategory"
	"saas/gen/ent/poststatus"
	"saas/gen/ent/posttag"
	"saas/gen/ent/posttype"
	"saas/gen/ent/todo"
	"saas/pkg/appfn"
	"saas/pkg/middleware"
	"strings"

	"entgo.io/contrib/entgql"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	idtail := strings.Split(id, "/")
	if len(idtail) != 2 {
		return nil, fmt.Errorf("invalid id format: %s", id)
	}

	table := ""
	// switch util.GetNodeType(id) {
	switch idtail[0] {
	case "todo":
		table = todo.Table
	case "PostType":
		table = posttype.Table
	case "PostStatus":
		table = poststatus.Table
	case "PostTag":
		table = posttag.Table
	case "PostCategory":
		table = postcategory.Table
	}

	return r.Plugin.EntDB.Client().Noder(ctx, idtail[1], ent.WithFixedNodeType(table))
	// return r.Plugin.EntDB.Client().Noder(ctx, id, ent.WithFixedNodeType(todo.Table))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	return r.Plugin.EntDB.Client().Noders(ctx, ids)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	cuser, err := middleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	if !appfn.IsUserSA(cuser) {
		return nil, fmt.Errorf("unauthorized access")
	}

	return r.Plugin.EntDB.Client().Post.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithPostOrder(orderBy),
			ent.WithPostFilter(where.Filter),
		)
}

// PostCategories is the resolver for the postCategories field.
func (r *queryResolver) PostCategories(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.PostCategoryOrder, where *ent.PostCategoryWhereInput) (*ent.PostCategoryConnection, error) {
	cuser, err := middleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	if !appfn.IsUserSA(cuser) {
		return nil, fmt.Errorf("unauthorized access")
	}

	return r.Plugin.EntDB.Client().PostCategory.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithPostCategoryOrder(orderBy),
			ent.WithPostCategoryFilter(where.Filter),
		)
}

// PostStatuses is the resolver for the postStatuses field.
func (r *queryResolver) PostStatuses(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.PostStatusOrder, where *ent.PostStatusWhereInput) (*ent.PostStatusConnection, error) {
	cuser, err := middleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	if !appfn.IsUserSA(cuser) {
		return nil, fmt.Errorf("unauthorized access")
	}

	return r.Plugin.EntDB.Client().PostStatus.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithPostStatusOrder(orderBy),
			ent.WithPostStatusFilter(where.Filter),
		)
}

// PostTypes is the resolver for the postTypes field.
func (r *queryResolver) PostTypes(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.PostTypeOrder, where *ent.PostTypeWhereInput) (*ent.PostTypeConnection, error) {
	cuser, err := middleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	if !appfn.IsUserSA(cuser) {
		return nil, fmt.Errorf("unauthorized access")
	}

	return r.Plugin.EntDB.Client().PostType.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithPostTypeOrder(orderBy),
			ent.WithPostTypeFilter(where.Filter),
		)
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy *ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	cuser, err := middleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	if !appfn.IsUserSA(cuser) {
		return nil, fmt.Errorf("unauthorized access")
	}
	return r.Plugin.EntDB.Client().Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
			ent.WithTodoFilter(where.Filter),
		)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
