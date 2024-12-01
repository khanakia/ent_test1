// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
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

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

var appImplementors = []string{"App", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*App) IsNode() {}

var mailconnImplementors = []string{"MailConn", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*MailConn) IsNode() {}

var mediaImplementors = []string{"Media", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Media) IsNode() {}

var oauthconnectionImplementors = []string{"OauthConnection", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*OauthConnection) IsNode() {}

var postImplementors = []string{"Post", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Post) IsNode() {}

var postcategoryImplementors = []string{"PostCategory", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*PostCategory) IsNode() {}

var poststatusImplementors = []string{"PostStatus", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*PostStatus) IsNode() {}

var posttagImplementors = []string{"PostTag", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*PostTag) IsNode() {}

var posttypeImplementors = []string{"PostType", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*PostType) IsNode() {}

var posttypeformImplementors = []string{"PostTypeForm", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*PostTypeForm) IsNode() {}

var templImplementors = []string{"Templ", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Templ) IsNode() {}

var todoImplementors = []string{"Todo", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Todo) IsNode() {}

var userImplementors = []string{"User", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*User) IsNode() {}

var workspaceImplementors = []string{"Workspace", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Workspace) IsNode() {}

var workspaceinviteImplementors = []string{"WorkspaceInvite", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*WorkspaceInvite) IsNode() {}

var workspaceuserImplementors = []string{"WorkspaceUser", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*WorkspaceUser) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, string) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, string) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, string) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id string) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id string, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id string) (Noder, error) {
	switch table {
	case app.Table:
		query := c.App.Query().
			Where(app.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, appImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case mailconn.Table:
		query := c.MailConn.Query().
			Where(mailconn.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, mailconnImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case media.Table:
		query := c.Media.Query().
			Where(media.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, mediaImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case oauthconnection.Table:
		query := c.OauthConnection.Query().
			Where(oauthconnection.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, oauthconnectionImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case post.Table:
		query := c.Post.Query().
			Where(post.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, postImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case postcategory.Table:
		query := c.PostCategory.Query().
			Where(postcategory.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, postcategoryImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case poststatus.Table:
		query := c.PostStatus.Query().
			Where(poststatus.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, poststatusImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case posttag.Table:
		query := c.PostTag.Query().
			Where(posttag.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, posttagImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case posttype.Table:
		query := c.PostType.Query().
			Where(posttype.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, posttypeImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case posttypeform.Table:
		query := c.PostTypeForm.Query().
			Where(posttypeform.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, posttypeformImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case templ.Table:
		query := c.Templ.Query().
			Where(templ.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, templImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case todo.Table:
		query := c.Todo.Query().
			Where(todo.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, todoImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, userImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case workspace.Table:
		query := c.Workspace.Query().
			Where(workspace.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, workspaceImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case workspaceinvite.Table:
		query := c.WorkspaceInvite.Query().
			Where(workspaceinvite.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, workspaceinviteImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case workspaceuser.Table:
		query := c.WorkspaceUser.Query().
			Where(workspaceuser.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, workspaceuserImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []string, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]string)
	id2idx := make(map[string][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []string) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[string][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case app.Table:
		query := c.App.Query().
			Where(app.IDIn(ids...))
		query, err := query.CollectFields(ctx, appImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case mailconn.Table:
		query := c.MailConn.Query().
			Where(mailconn.IDIn(ids...))
		query, err := query.CollectFields(ctx, mailconnImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case media.Table:
		query := c.Media.Query().
			Where(media.IDIn(ids...))
		query, err := query.CollectFields(ctx, mediaImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case oauthconnection.Table:
		query := c.OauthConnection.Query().
			Where(oauthconnection.IDIn(ids...))
		query, err := query.CollectFields(ctx, oauthconnectionImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case post.Table:
		query := c.Post.Query().
			Where(post.IDIn(ids...))
		query, err := query.CollectFields(ctx, postImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case postcategory.Table:
		query := c.PostCategory.Query().
			Where(postcategory.IDIn(ids...))
		query, err := query.CollectFields(ctx, postcategoryImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case poststatus.Table:
		query := c.PostStatus.Query().
			Where(poststatus.IDIn(ids...))
		query, err := query.CollectFields(ctx, poststatusImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case posttag.Table:
		query := c.PostTag.Query().
			Where(posttag.IDIn(ids...))
		query, err := query.CollectFields(ctx, posttagImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case posttype.Table:
		query := c.PostType.Query().
			Where(posttype.IDIn(ids...))
		query, err := query.CollectFields(ctx, posttypeImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case posttypeform.Table:
		query := c.PostTypeForm.Query().
			Where(posttypeform.IDIn(ids...))
		query, err := query.CollectFields(ctx, posttypeformImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case templ.Table:
		query := c.Templ.Query().
			Where(templ.IDIn(ids...))
		query, err := query.CollectFields(ctx, templImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case todo.Table:
		query := c.Todo.Query().
			Where(todo.IDIn(ids...))
		query, err := query.CollectFields(ctx, todoImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, userImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workspace.Table:
		query := c.Workspace.Query().
			Where(workspace.IDIn(ids...))
		query, err := query.CollectFields(ctx, workspaceImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workspaceinvite.Table:
		query := c.WorkspaceInvite.Query().
			Where(workspaceinvite.IDIn(ids...))
		query, err := query.CollectFields(ctx, workspaceinviteImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workspaceuser.Table:
		query := c.WorkspaceUser.Query().
			Where(workspaceuser.IDIn(ids...))
		query, err := query.CollectFields(ctx, workspaceuserImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
