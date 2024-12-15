// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"saas/gen/ent"
	"saas/gen/ent/adminuser"
	"saas/gen/ent/app"
	"saas/gen/ent/appperm"
	"saas/gen/ent/approle"
	"saas/gen/ent/approleperm"
	"saas/gen/ent/appuser"
	"saas/gen/ent/kache"
	"saas/gen/ent/keyvalue"
	"saas/gen/ent/mailconn"
	"saas/gen/ent/media"
	"saas/gen/ent/oauthconnection"
	"saas/gen/ent/plan"
	"saas/gen/ent/post"
	"saas/gen/ent/postcategory"
	"saas/gen/ent/poststatus"
	"saas/gen/ent/posttag"
	"saas/gen/ent/posttype"
	"saas/gen/ent/posttypeform"
	"saas/gen/ent/predicate"
	"saas/gen/ent/session"
	"saas/gen/ent/temp"
	"saas/gen/ent/templ"
	"saas/gen/ent/todo"
	"saas/gen/ent/user"
	"saas/gen/ent/workspace"
	"saas/gen/ent/workspaceinvite"
	"saas/gen/ent/workspaceuser"

	"entgo.io/ent/dialect/sql"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The AdminUserFunc type is an adapter to allow the use of ordinary function as a Querier.
type AdminUserFunc func(context.Context, *ent.AdminUserQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AdminUserFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AdminUserQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AdminUserQuery", q)
}

// The TraverseAdminUser type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAdminUser func(context.Context, *ent.AdminUserQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAdminUser) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAdminUser) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AdminUserQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AdminUserQuery", q)
}

// The AppFunc type is an adapter to allow the use of ordinary function as a Querier.
type AppFunc func(context.Context, *ent.AppQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AppFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AppQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AppQuery", q)
}

// The TraverseApp type is an adapter to allow the use of ordinary function as Traverser.
type TraverseApp func(context.Context, *ent.AppQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseApp) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseApp) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AppQuery", q)
}

// The AppPermFunc type is an adapter to allow the use of ordinary function as a Querier.
type AppPermFunc func(context.Context, *ent.AppPermQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AppPermFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AppPermQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AppPermQuery", q)
}

// The TraverseAppPerm type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAppPerm func(context.Context, *ent.AppPermQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAppPerm) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAppPerm) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppPermQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AppPermQuery", q)
}

// The AppRoleFunc type is an adapter to allow the use of ordinary function as a Querier.
type AppRoleFunc func(context.Context, *ent.AppRoleQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AppRoleFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AppRoleQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AppRoleQuery", q)
}

// The TraverseAppRole type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAppRole func(context.Context, *ent.AppRoleQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAppRole) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAppRole) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppRoleQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AppRoleQuery", q)
}

// The AppRolePermFunc type is an adapter to allow the use of ordinary function as a Querier.
type AppRolePermFunc func(context.Context, *ent.AppRolePermQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AppRolePermFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AppRolePermQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AppRolePermQuery", q)
}

// The TraverseAppRolePerm type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAppRolePerm func(context.Context, *ent.AppRolePermQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAppRolePerm) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAppRolePerm) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppRolePermQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AppRolePermQuery", q)
}

// The AppUserFunc type is an adapter to allow the use of ordinary function as a Querier.
type AppUserFunc func(context.Context, *ent.AppUserQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AppUserFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AppUserQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AppUserQuery", q)
}

// The TraverseAppUser type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAppUser func(context.Context, *ent.AppUserQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAppUser) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAppUser) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppUserQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AppUserQuery", q)
}

// The KacheFunc type is an adapter to allow the use of ordinary function as a Querier.
type KacheFunc func(context.Context, *ent.KacheQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f KacheFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.KacheQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.KacheQuery", q)
}

// The TraverseKache type is an adapter to allow the use of ordinary function as Traverser.
type TraverseKache func(context.Context, *ent.KacheQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseKache) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseKache) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.KacheQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.KacheQuery", q)
}

// The KeyvalueFunc type is an adapter to allow the use of ordinary function as a Querier.
type KeyvalueFunc func(context.Context, *ent.KeyvalueQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f KeyvalueFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.KeyvalueQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.KeyvalueQuery", q)
}

// The TraverseKeyvalue type is an adapter to allow the use of ordinary function as Traverser.
type TraverseKeyvalue func(context.Context, *ent.KeyvalueQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseKeyvalue) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseKeyvalue) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.KeyvalueQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.KeyvalueQuery", q)
}

// The MailConnFunc type is an adapter to allow the use of ordinary function as a Querier.
type MailConnFunc func(context.Context, *ent.MailConnQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MailConnFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MailConnQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MailConnQuery", q)
}

// The TraverseMailConn type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMailConn func(context.Context, *ent.MailConnQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMailConn) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMailConn) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MailConnQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MailConnQuery", q)
}

// The MediaFunc type is an adapter to allow the use of ordinary function as a Querier.
type MediaFunc func(context.Context, *ent.MediaQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MediaFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MediaQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MediaQuery", q)
}

// The TraverseMedia type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMedia func(context.Context, *ent.MediaQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMedia) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMedia) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MediaQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MediaQuery", q)
}

// The OauthConnectionFunc type is an adapter to allow the use of ordinary function as a Querier.
type OauthConnectionFunc func(context.Context, *ent.OauthConnectionQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f OauthConnectionFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.OauthConnectionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.OauthConnectionQuery", q)
}

// The TraverseOauthConnection type is an adapter to allow the use of ordinary function as Traverser.
type TraverseOauthConnection func(context.Context, *ent.OauthConnectionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseOauthConnection) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseOauthConnection) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OauthConnectionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.OauthConnectionQuery", q)
}

// The PlanFunc type is an adapter to allow the use of ordinary function as a Querier.
type PlanFunc func(context.Context, *ent.PlanQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PlanFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PlanQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PlanQuery", q)
}

// The TraversePlan type is an adapter to allow the use of ordinary function as Traverser.
type TraversePlan func(context.Context, *ent.PlanQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePlan) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePlan) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PlanQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PlanQuery", q)
}

// The PostFunc type is an adapter to allow the use of ordinary function as a Querier.
type PostFunc func(context.Context, *ent.PostQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PostFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PostQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PostQuery", q)
}

// The TraversePost type is an adapter to allow the use of ordinary function as Traverser.
type TraversePost func(context.Context, *ent.PostQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePost) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePost) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PostQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PostQuery", q)
}

// The PostCategoryFunc type is an adapter to allow the use of ordinary function as a Querier.
type PostCategoryFunc func(context.Context, *ent.PostCategoryQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PostCategoryFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PostCategoryQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PostCategoryQuery", q)
}

// The TraversePostCategory type is an adapter to allow the use of ordinary function as Traverser.
type TraversePostCategory func(context.Context, *ent.PostCategoryQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePostCategory) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePostCategory) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PostCategoryQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PostCategoryQuery", q)
}

// The PostStatusFunc type is an adapter to allow the use of ordinary function as a Querier.
type PostStatusFunc func(context.Context, *ent.PostStatusQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PostStatusFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PostStatusQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PostStatusQuery", q)
}

// The TraversePostStatus type is an adapter to allow the use of ordinary function as Traverser.
type TraversePostStatus func(context.Context, *ent.PostStatusQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePostStatus) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePostStatus) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PostStatusQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PostStatusQuery", q)
}

// The PostTagFunc type is an adapter to allow the use of ordinary function as a Querier.
type PostTagFunc func(context.Context, *ent.PostTagQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PostTagFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PostTagQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PostTagQuery", q)
}

// The TraversePostTag type is an adapter to allow the use of ordinary function as Traverser.
type TraversePostTag func(context.Context, *ent.PostTagQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePostTag) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePostTag) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PostTagQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PostTagQuery", q)
}

// The PostTypeFunc type is an adapter to allow the use of ordinary function as a Querier.
type PostTypeFunc func(context.Context, *ent.PostTypeQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PostTypeFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PostTypeQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PostTypeQuery", q)
}

// The TraversePostType type is an adapter to allow the use of ordinary function as Traverser.
type TraversePostType func(context.Context, *ent.PostTypeQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePostType) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePostType) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PostTypeQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PostTypeQuery", q)
}

// The PostTypeFormFunc type is an adapter to allow the use of ordinary function as a Querier.
type PostTypeFormFunc func(context.Context, *ent.PostTypeFormQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PostTypeFormFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PostTypeFormQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PostTypeFormQuery", q)
}

// The TraversePostTypeForm type is an adapter to allow the use of ordinary function as Traverser.
type TraversePostTypeForm func(context.Context, *ent.PostTypeFormQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePostTypeForm) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePostTypeForm) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PostTypeFormQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PostTypeFormQuery", q)
}

// The SessionFunc type is an adapter to allow the use of ordinary function as a Querier.
type SessionFunc func(context.Context, *ent.SessionQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f SessionFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.SessionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.SessionQuery", q)
}

// The TraverseSession type is an adapter to allow the use of ordinary function as Traverser.
type TraverseSession func(context.Context, *ent.SessionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseSession) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseSession) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SessionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.SessionQuery", q)
}

// The TempFunc type is an adapter to allow the use of ordinary function as a Querier.
type TempFunc func(context.Context, *ent.TempQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TempFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TempQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TempQuery", q)
}

// The TraverseTemp type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTemp func(context.Context, *ent.TempQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTemp) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTemp) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TempQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TempQuery", q)
}

// The TemplFunc type is an adapter to allow the use of ordinary function as a Querier.
type TemplFunc func(context.Context, *ent.TemplQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TemplFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TemplQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TemplQuery", q)
}

// The TraverseTempl type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTempl func(context.Context, *ent.TemplQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTempl) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTempl) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TemplQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TemplQuery", q)
}

// The TodoFunc type is an adapter to allow the use of ordinary function as a Querier.
type TodoFunc func(context.Context, *ent.TodoQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TodoFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TodoQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TodoQuery", q)
}

// The TraverseTodo type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTodo func(context.Context, *ent.TodoQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTodo) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTodo) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TodoQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TodoQuery", q)
}

// The UserFunc type is an adapter to allow the use of ordinary function as a Querier.
type UserFunc func(context.Context, *ent.UserQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f UserFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// The TraverseUser type is an adapter to allow the use of ordinary function as Traverser.
type TraverseUser func(context.Context, *ent.UserQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseUser) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseUser) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// The WorkspaceFunc type is an adapter to allow the use of ordinary function as a Querier.
type WorkspaceFunc func(context.Context, *ent.WorkspaceQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f WorkspaceFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.WorkspaceQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.WorkspaceQuery", q)
}

// The TraverseWorkspace type is an adapter to allow the use of ordinary function as Traverser.
type TraverseWorkspace func(context.Context, *ent.WorkspaceQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseWorkspace) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseWorkspace) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.WorkspaceQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.WorkspaceQuery", q)
}

// The WorkspaceInviteFunc type is an adapter to allow the use of ordinary function as a Querier.
type WorkspaceInviteFunc func(context.Context, *ent.WorkspaceInviteQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f WorkspaceInviteFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.WorkspaceInviteQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.WorkspaceInviteQuery", q)
}

// The TraverseWorkspaceInvite type is an adapter to allow the use of ordinary function as Traverser.
type TraverseWorkspaceInvite func(context.Context, *ent.WorkspaceInviteQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseWorkspaceInvite) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseWorkspaceInvite) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.WorkspaceInviteQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.WorkspaceInviteQuery", q)
}

// The WorkspaceUserFunc type is an adapter to allow the use of ordinary function as a Querier.
type WorkspaceUserFunc func(context.Context, *ent.WorkspaceUserQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f WorkspaceUserFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.WorkspaceUserQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.WorkspaceUserQuery", q)
}

// The TraverseWorkspaceUser type is an adapter to allow the use of ordinary function as Traverser.
type TraverseWorkspaceUser func(context.Context, *ent.WorkspaceUserQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseWorkspaceUser) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseWorkspaceUser) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.WorkspaceUserQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.WorkspaceUserQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.AdminUserQuery:
		return &query[*ent.AdminUserQuery, predicate.AdminUser, adminuser.OrderOption]{typ: ent.TypeAdminUser, tq: q}, nil
	case *ent.AppQuery:
		return &query[*ent.AppQuery, predicate.App, app.OrderOption]{typ: ent.TypeApp, tq: q}, nil
	case *ent.AppPermQuery:
		return &query[*ent.AppPermQuery, predicate.AppPerm, appperm.OrderOption]{typ: ent.TypeAppPerm, tq: q}, nil
	case *ent.AppRoleQuery:
		return &query[*ent.AppRoleQuery, predicate.AppRole, approle.OrderOption]{typ: ent.TypeAppRole, tq: q}, nil
	case *ent.AppRolePermQuery:
		return &query[*ent.AppRolePermQuery, predicate.AppRolePerm, approleperm.OrderOption]{typ: ent.TypeAppRolePerm, tq: q}, nil
	case *ent.AppUserQuery:
		return &query[*ent.AppUserQuery, predicate.AppUser, appuser.OrderOption]{typ: ent.TypeAppUser, tq: q}, nil
	case *ent.KacheQuery:
		return &query[*ent.KacheQuery, predicate.Kache, kache.OrderOption]{typ: ent.TypeKache, tq: q}, nil
	case *ent.KeyvalueQuery:
		return &query[*ent.KeyvalueQuery, predicate.Keyvalue, keyvalue.OrderOption]{typ: ent.TypeKeyvalue, tq: q}, nil
	case *ent.MailConnQuery:
		return &query[*ent.MailConnQuery, predicate.MailConn, mailconn.OrderOption]{typ: ent.TypeMailConn, tq: q}, nil
	case *ent.MediaQuery:
		return &query[*ent.MediaQuery, predicate.Media, media.OrderOption]{typ: ent.TypeMedia, tq: q}, nil
	case *ent.OauthConnectionQuery:
		return &query[*ent.OauthConnectionQuery, predicate.OauthConnection, oauthconnection.OrderOption]{typ: ent.TypeOauthConnection, tq: q}, nil
	case *ent.PlanQuery:
		return &query[*ent.PlanQuery, predicate.Plan, plan.OrderOption]{typ: ent.TypePlan, tq: q}, nil
	case *ent.PostQuery:
		return &query[*ent.PostQuery, predicate.Post, post.OrderOption]{typ: ent.TypePost, tq: q}, nil
	case *ent.PostCategoryQuery:
		return &query[*ent.PostCategoryQuery, predicate.PostCategory, postcategory.OrderOption]{typ: ent.TypePostCategory, tq: q}, nil
	case *ent.PostStatusQuery:
		return &query[*ent.PostStatusQuery, predicate.PostStatus, poststatus.OrderOption]{typ: ent.TypePostStatus, tq: q}, nil
	case *ent.PostTagQuery:
		return &query[*ent.PostTagQuery, predicate.PostTag, posttag.OrderOption]{typ: ent.TypePostTag, tq: q}, nil
	case *ent.PostTypeQuery:
		return &query[*ent.PostTypeQuery, predicate.PostType, posttype.OrderOption]{typ: ent.TypePostType, tq: q}, nil
	case *ent.PostTypeFormQuery:
		return &query[*ent.PostTypeFormQuery, predicate.PostTypeForm, posttypeform.OrderOption]{typ: ent.TypePostTypeForm, tq: q}, nil
	case *ent.SessionQuery:
		return &query[*ent.SessionQuery, predicate.Session, session.OrderOption]{typ: ent.TypeSession, tq: q}, nil
	case *ent.TempQuery:
		return &query[*ent.TempQuery, predicate.Temp, temp.OrderOption]{typ: ent.TypeTemp, tq: q}, nil
	case *ent.TemplQuery:
		return &query[*ent.TemplQuery, predicate.Templ, templ.OrderOption]{typ: ent.TypeTempl, tq: q}, nil
	case *ent.TodoQuery:
		return &query[*ent.TodoQuery, predicate.Todo, todo.OrderOption]{typ: ent.TypeTodo, tq: q}, nil
	case *ent.UserQuery:
		return &query[*ent.UserQuery, predicate.User, user.OrderOption]{typ: ent.TypeUser, tq: q}, nil
	case *ent.WorkspaceQuery:
		return &query[*ent.WorkspaceQuery, predicate.Workspace, workspace.OrderOption]{typ: ent.TypeWorkspace, tq: q}, nil
	case *ent.WorkspaceInviteQuery:
		return &query[*ent.WorkspaceInviteQuery, predicate.WorkspaceInvite, workspaceinvite.OrderOption]{typ: ent.TypeWorkspaceInvite, tq: q}, nil
	case *ent.WorkspaceUserQuery:
		return &query[*ent.WorkspaceUserQuery, predicate.WorkspaceUser, workspaceuser.OrderOption]{typ: ent.TypeWorkspaceUser, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
