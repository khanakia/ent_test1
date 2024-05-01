// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"
	"saas/gen/ent"
	"saas/gen/ent/admin"
	"saas/gen/ent/kache"
	"saas/gen/ent/keyvalue"
	"saas/gen/ent/mailconnection"
	"saas/gen/ent/plan"
	"saas/gen/ent/predicate"
	"saas/gen/ent/project"
	"saas/gen/ent/session"
	"saas/gen/ent/temp"
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

// The AdminFunc type is an adapter to allow the use of ordinary function as a Querier.
type AdminFunc func(context.Context, *ent.AdminQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AdminFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AdminQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AdminQuery", q)
}

// The TraverseAdmin type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAdmin func(context.Context, *ent.AdminQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAdmin) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAdmin) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AdminQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AdminQuery", q)
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

// The MailConnectionFunc type is an adapter to allow the use of ordinary function as a Querier.
type MailConnectionFunc func(context.Context, *ent.MailConnectionQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MailConnectionFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MailConnectionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MailConnectionQuery", q)
}

// The TraverseMailConnection type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMailConnection func(context.Context, *ent.MailConnectionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMailConnection) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMailConnection) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MailConnectionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MailConnectionQuery", q)
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

// The ProjectFunc type is an adapter to allow the use of ordinary function as a Querier.
type ProjectFunc func(context.Context, *ent.ProjectQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f ProjectFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.ProjectQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.ProjectQuery", q)
}

// The TraverseProject type is an adapter to allow the use of ordinary function as Traverser.
type TraverseProject func(context.Context, *ent.ProjectQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseProject) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseProject) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProjectQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.ProjectQuery", q)
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
	case *ent.AdminQuery:
		return &query[*ent.AdminQuery, predicate.Admin, admin.OrderOption]{typ: ent.TypeAdmin, tq: q}, nil
	case *ent.KacheQuery:
		return &query[*ent.KacheQuery, predicate.Kache, kache.OrderOption]{typ: ent.TypeKache, tq: q}, nil
	case *ent.KeyvalueQuery:
		return &query[*ent.KeyvalueQuery, predicate.Keyvalue, keyvalue.OrderOption]{typ: ent.TypeKeyvalue, tq: q}, nil
	case *ent.MailConnectionQuery:
		return &query[*ent.MailConnectionQuery, predicate.MailConnection, mailconnection.OrderOption]{typ: ent.TypeMailConnection, tq: q}, nil
	case *ent.PlanQuery:
		return &query[*ent.PlanQuery, predicate.Plan, plan.OrderOption]{typ: ent.TypePlan, tq: q}, nil
	case *ent.ProjectQuery:
		return &query[*ent.ProjectQuery, predicate.Project, project.OrderOption]{typ: ent.TypeProject, tq: q}, nil
	case *ent.SessionQuery:
		return &query[*ent.SessionQuery, predicate.Session, session.OrderOption]{typ: ent.TypeSession, tq: q}, nil
	case *ent.TempQuery:
		return &query[*ent.TempQuery, predicate.Temp, temp.OrderOption]{typ: ent.TypeTemp, tq: q}, nil
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
