package gqlsaresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"fmt"
	"gqlsa/graph/generated"
	"reflect"
	"saas/gen/ent"
	"saas/gen/ent/app"
	"saas/gen/ent/mailconn"
	"saas/gen/ent/media"
	"saas/gen/ent/oauthconnection"
	"saas/gen/ent/post"
	"saas/gen/ent/postcategory"
	"saas/gen/ent/poststatus"
	"saas/gen/ent/posttag"
	"saas/gen/ent/posttype"
	"saas/gen/ent/posttypeform"
	"saas/gen/ent/templ"
	"saas/gen/ent/todo"
	"saas/gen/ent/user"
	"saas/gen/ent/workspace"
	"saas/gen/ent/workspaceinvite"
	"saas/gen/ent/workspaceuser"
	"saas/pkg/middleware/adminauthmiddleware"
	"saas/pkg/middleware/appmiddleware"
	"strings"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
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
	case "Post":
		table = post.Table
	case "OauthConnection":
		table = oauthconnection.Table
	case "User":
		table = user.Table
	case "Workspace":
		table = workspace.Table
	case "WorkspaceInvite":
		table = workspaceinvite.Table
	case "WorkspaceUser":
		table = workspaceuser.Table
	case "MailConn":
		table = mailconn.Table
	case "Templ":
		table = templ.Table
	}

	// there is no way i could send the sql.Where query to Noder and also noder only use the asked field to select query selectFields
	// below so i had to find a hack and now i always inject appID as query field so Noder function sends the sql query and apend select id, app_id from table
	// and then we use reflect to extract the AppID from the Noder struct and validate that if user is authorized to access the record
	fc := graphql.GetFieldContext(ctx)
	selSet1 := fc.Field.SelectionSet[1]
	switch sel := selSet1.(type) {
	case *ast.InlineFragment:
		newSelections := append(sel.SelectionSet, &ast.Field{
			Alias:      "appID",
			Name:       "appID",
			Arguments:  nil,
			Directives: nil,
		})

		sel.SelectionSet = newSelections
	}

	noder, err := r.Plugin.EntDB.Client().Noder(ctx, idtail[1], ent.WithFixedNodeType(table))

	nodev := reflect.ValueOf(noder)

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)
	appID := reflect.Indirect(nodev).FieldByName("AppID").String()

	if appID != app.ID {
		return nil, fmt.Errorf("access denied app")
	}

	fmt.Println(reflect.Indirect(nodev).FieldByName("AppID"))

	return noder, err
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	return r.Plugin.EntDB.Client().Noders(ctx, ids)
}

// Apps is the resolver for the apps field.
func (r *queryResolver) Apps(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.AppOrder, where *ent.AppWhereInput) (*ent.AppConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	return r.Plugin.EntDB.Client().App.Query().
		Where(app.AdminUserID(cuser.ID)).
		Paginate(ctx, after, first, before, last,
			// ent.WithAppOrder(orderBy),
			ent.WithAppFilter(where.Filter),
		)
}

// MailConns is the resolver for the mailConns field.
func (r *queryResolver) MailConns(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.MailConnOrder, where *ent.MailConnWhereInput) (*ent.MailConnConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}
	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().MailConn.Query().
		Where(mailconn.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithMailConnOrder(orderBy),
			ent.WithMailConnFilter(where.Filter),
		)
}

// Medias is the resolver for the medias field.
func (r *queryResolver) Medias(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.MediaOrder, where *ent.MediaWhereInput) (*ent.MediaConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().Media.Query().
		Where(media.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithMediaOrder(orderBy),
			ent.WithMediaFilter(where.Filter),
		)
}

// OauthConnections is the resolver for the oauthConnections field.
func (r *queryResolver) OauthConnections(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.OauthConnectionOrder, where *ent.OauthConnectionWhereInput) (*ent.OauthConnectionConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}
	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().OauthConnection.Query().
		Where(oauthconnection.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithOauthConnectionOrder(orderBy),
			ent.WithOauthConnectionFilter(where.Filter),
		)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().Post.Query().
		Where(post.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithPostOrder(orderBy),
			ent.WithPostFilter(where.Filter),
		)
}

// PostCategories is the resolver for the postCategories field.
func (r *queryResolver) PostCategories(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.PostCategoryOrder, where *ent.PostCategoryWhereInput) (*ent.PostCategoryConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().PostCategory.Query().
		Where(postcategory.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithPostCategoryOrder(orderBy),
			ent.WithPostCategoryFilter(where.Filter),
		)
}

// PostStatuses is the resolver for the postStatuses field.
func (r *queryResolver) PostStatuses(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.PostStatusOrder, where *ent.PostStatusWhereInput) (*ent.PostStatusConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().PostStatus.Query().
		Where(poststatus.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithPostStatusOrder(orderBy),
			ent.WithPostStatusFilter(where.Filter),
		)
}

// PostTags is the resolver for the postTags field.
func (r *queryResolver) PostTags(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.PostTagOrder, where *ent.PostTagWhereInput) (*ent.PostTagConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().PostTag.Query().
		Where(posttag.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithPostTagOrder(orderBy),
			ent.WithPostTagFilter(where.Filter),
		)
}

// PostTypes is the resolver for the postTypes field.
func (r *queryResolver) PostTypes(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.PostTypeOrder, where *ent.PostTypeWhereInput) (*ent.PostTypeConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	// time.Sleep(4 * time.Second)

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().PostType.Query().
		Where(posttype.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithPostTypeOrder(orderBy),
			ent.WithPostTypeFilter(where.Filter),
		)
}

// PostTypeForms is the resolver for the postTypeForms field.
func (r *queryResolver) PostTypeForms(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.PostTypeFormOrder, where *ent.PostTypeFormWhereInput) (*ent.PostTypeFormConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().PostTypeForm.Query().
		Where(posttypeform.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithPostTypeFormOrder(orderBy),
			ent.WithPostTypeFormFilter(where.Filter),
		)
}

// Templs is the resolver for the templs field.
func (r *queryResolver) Templs(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.TemplOrder, where *ent.TemplWhereInput) (*ent.TemplConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}
	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().Templ.Query().
		Where(templ.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithTemplOrder(orderBy),
			ent.WithTemplFilter(where.Filter),
		)
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy *ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().Todo.Query().
		Where(todo.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
			ent.WithTodoFilter(where.Filter),
		)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().User.Query().
		Where(user.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
			ent.WithUserFilter(where.Filter),
		)
}

// Workspaces is the resolver for the workspaces field.
func (r *queryResolver) Workspaces(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.WorkspaceOrder, where *ent.WorkspaceWhereInput) (*ent.WorkspaceConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().Workspace.Query().
		Where(workspace.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithWorkspaceOrder(orderBy),
			ent.WithWorkspaceFilter(where.Filter),
		)
}

// WorkspaceInvites is the resolver for the workspaceInvites field.
func (r *queryResolver) WorkspaceInvites(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.WorkspaceInviteOrder, where *ent.WorkspaceInviteWhereInput) (*ent.WorkspaceInviteConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().WorkspaceInvite.Query().
		Where(workspaceinvite.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithWorkspaceInviteOrder(orderBy),
			ent.WithWorkspaceInviteFilter(where.Filter),
		)
}

// WorkspaceUsers is the resolver for the workspaceUsers field.
func (r *queryResolver) WorkspaceUsers(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy []*ent.WorkspaceUserOrder, where *ent.WorkspaceUserWhereInput) (*ent.WorkspaceUserConnection, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	app := appmiddleware.MustGetAppFromGqlCtx(ctx)

	return r.Plugin.EntDB.Client().WorkspaceUser.Query().
		Where(workspaceuser.AppID(app.ID)).
		Paginate(ctx, after, first, before, last,
			ent.WithWorkspaceUserOrder(orderBy),
			ent.WithWorkspaceUserFilter(where.Filter),
		)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
