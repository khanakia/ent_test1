// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"saas/gen/ent/adminuser"
	"saas/gen/ent/app"
	"saas/gen/ent/appuser"
	"saas/gen/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminUserQuery is the builder for querying AdminUser entities.
type AdminUserQuery struct {
	config
	ctx               *QueryContext
	order             []adminuser.OrderOption
	inters            []Interceptor
	predicates        []predicate.AdminUser
	withApps          *AppQuery
	withAppUsers      *AppUserQuery
	loadTotal         []func(context.Context, []*AdminUser) error
	modifiers         []func(*sql.Selector)
	withNamedApps     map[string]*AppQuery
	withNamedAppUsers map[string]*AppUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdminUserQuery builder.
func (auq *AdminUserQuery) Where(ps ...predicate.AdminUser) *AdminUserQuery {
	auq.predicates = append(auq.predicates, ps...)
	return auq
}

// Limit the number of records to be returned by this query.
func (auq *AdminUserQuery) Limit(limit int) *AdminUserQuery {
	auq.ctx.Limit = &limit
	return auq
}

// Offset to start from.
func (auq *AdminUserQuery) Offset(offset int) *AdminUserQuery {
	auq.ctx.Offset = &offset
	return auq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (auq *AdminUserQuery) Unique(unique bool) *AdminUserQuery {
	auq.ctx.Unique = &unique
	return auq
}

// Order specifies how the records should be ordered.
func (auq *AdminUserQuery) Order(o ...adminuser.OrderOption) *AdminUserQuery {
	auq.order = append(auq.order, o...)
	return auq
}

// QueryApps chains the current query on the "apps" edge.
func (auq *AdminUserQuery) QueryApps() *AppQuery {
	query := (&AppClient{config: auq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := auq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := auq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adminuser.Table, adminuser.FieldID, selector),
			sqlgraph.To(app.Table, app.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, adminuser.AppsTable, adminuser.AppsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(auq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAppUsers chains the current query on the "app_users" edge.
func (auq *AdminUserQuery) QueryAppUsers() *AppUserQuery {
	query := (&AppUserClient{config: auq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := auq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := auq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adminuser.Table, adminuser.FieldID, selector),
			sqlgraph.To(appuser.Table, appuser.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, adminuser.AppUsersTable, adminuser.AppUsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(auq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AdminUser entity from the query.
// Returns a *NotFoundError when no AdminUser was found.
func (auq *AdminUserQuery) First(ctx context.Context) (*AdminUser, error) {
	nodes, err := auq.Limit(1).All(setContextOp(ctx, auq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adminuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (auq *AdminUserQuery) FirstX(ctx context.Context) *AdminUser {
	node, err := auq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AdminUser ID from the query.
// Returns a *NotFoundError when no AdminUser ID was found.
func (auq *AdminUserQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = auq.Limit(1).IDs(setContextOp(ctx, auq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adminuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (auq *AdminUserQuery) FirstIDX(ctx context.Context) string {
	id, err := auq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AdminUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AdminUser entity is found.
// Returns a *NotFoundError when no AdminUser entities are found.
func (auq *AdminUserQuery) Only(ctx context.Context) (*AdminUser, error) {
	nodes, err := auq.Limit(2).All(setContextOp(ctx, auq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adminuser.Label}
	default:
		return nil, &NotSingularError{adminuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (auq *AdminUserQuery) OnlyX(ctx context.Context) *AdminUser {
	node, err := auq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AdminUser ID in the query.
// Returns a *NotSingularError when more than one AdminUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (auq *AdminUserQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = auq.Limit(2).IDs(setContextOp(ctx, auq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adminuser.Label}
	default:
		err = &NotSingularError{adminuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (auq *AdminUserQuery) OnlyIDX(ctx context.Context) string {
	id, err := auq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AdminUsers.
func (auq *AdminUserQuery) All(ctx context.Context) ([]*AdminUser, error) {
	ctx = setContextOp(ctx, auq.ctx, ent.OpQueryAll)
	if err := auq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AdminUser, *AdminUserQuery]()
	return withInterceptors[[]*AdminUser](ctx, auq, qr, auq.inters)
}

// AllX is like All, but panics if an error occurs.
func (auq *AdminUserQuery) AllX(ctx context.Context) []*AdminUser {
	nodes, err := auq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AdminUser IDs.
func (auq *AdminUserQuery) IDs(ctx context.Context) (ids []string, err error) {
	if auq.ctx.Unique == nil && auq.path != nil {
		auq.Unique(true)
	}
	ctx = setContextOp(ctx, auq.ctx, ent.OpQueryIDs)
	if err = auq.Select(adminuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (auq *AdminUserQuery) IDsX(ctx context.Context) []string {
	ids, err := auq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (auq *AdminUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, auq.ctx, ent.OpQueryCount)
	if err := auq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, auq, querierCount[*AdminUserQuery](), auq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (auq *AdminUserQuery) CountX(ctx context.Context) int {
	count, err := auq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (auq *AdminUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, auq.ctx, ent.OpQueryExist)
	switch _, err := auq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (auq *AdminUserQuery) ExistX(ctx context.Context) bool {
	exist, err := auq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdminUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (auq *AdminUserQuery) Clone() *AdminUserQuery {
	if auq == nil {
		return nil
	}
	return &AdminUserQuery{
		config:       auq.config,
		ctx:          auq.ctx.Clone(),
		order:        append([]adminuser.OrderOption{}, auq.order...),
		inters:       append([]Interceptor{}, auq.inters...),
		predicates:   append([]predicate.AdminUser{}, auq.predicates...),
		withApps:     auq.withApps.Clone(),
		withAppUsers: auq.withAppUsers.Clone(),
		// clone intermediate query.
		sql:       auq.sql.Clone(),
		path:      auq.path,
		modifiers: append([]func(*sql.Selector){}, auq.modifiers...),
	}
}

// WithApps tells the query-builder to eager-load the nodes that are connected to
// the "apps" edge. The optional arguments are used to configure the query builder of the edge.
func (auq *AdminUserQuery) WithApps(opts ...func(*AppQuery)) *AdminUserQuery {
	query := (&AppClient{config: auq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	auq.withApps = query
	return auq
}

// WithAppUsers tells the query-builder to eager-load the nodes that are connected to
// the "app_users" edge. The optional arguments are used to configure the query builder of the edge.
func (auq *AdminUserQuery) WithAppUsers(opts ...func(*AppUserQuery)) *AdminUserQuery {
	query := (&AppUserClient{config: auq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	auq.withAppUsers = query
	return auq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AdminUser.Query().
//		GroupBy(adminuser.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (auq *AdminUserQuery) GroupBy(field string, fields ...string) *AdminUserGroupBy {
	auq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AdminUserGroupBy{build: auq}
	grbuild.flds = &auq.ctx.Fields
	grbuild.label = adminuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.AdminUser.Query().
//		Select(adminuser.FieldCreatedAt).
//		Scan(ctx, &v)
func (auq *AdminUserQuery) Select(fields ...string) *AdminUserSelect {
	auq.ctx.Fields = append(auq.ctx.Fields, fields...)
	sbuild := &AdminUserSelect{AdminUserQuery: auq}
	sbuild.label = adminuser.Label
	sbuild.flds, sbuild.scan = &auq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AdminUserSelect configured with the given aggregations.
func (auq *AdminUserQuery) Aggregate(fns ...AggregateFunc) *AdminUserSelect {
	return auq.Select().Aggregate(fns...)
}

func (auq *AdminUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range auq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, auq); err != nil {
				return err
			}
		}
	}
	for _, f := range auq.ctx.Fields {
		if !adminuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if auq.path != nil {
		prev, err := auq.path(ctx)
		if err != nil {
			return err
		}
		auq.sql = prev
	}
	return nil
}

func (auq *AdminUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AdminUser, error) {
	var (
		nodes       = []*AdminUser{}
		_spec       = auq.querySpec()
		loadedTypes = [2]bool{
			auq.withApps != nil,
			auq.withAppUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AdminUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AdminUser{config: auq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(auq.modifiers) > 0 {
		_spec.Modifiers = auq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, auq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := auq.withApps; query != nil {
		if err := auq.loadApps(ctx, query, nodes,
			func(n *AdminUser) { n.Edges.Apps = []*App{} },
			func(n *AdminUser, e *App) { n.Edges.Apps = append(n.Edges.Apps, e) }); err != nil {
			return nil, err
		}
	}
	if query := auq.withAppUsers; query != nil {
		if err := auq.loadAppUsers(ctx, query, nodes,
			func(n *AdminUser) { n.Edges.AppUsers = []*AppUser{} },
			func(n *AdminUser, e *AppUser) { n.Edges.AppUsers = append(n.Edges.AppUsers, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range auq.withNamedApps {
		if err := auq.loadApps(ctx, query, nodes,
			func(n *AdminUser) { n.appendNamedApps(name) },
			func(n *AdminUser, e *App) { n.appendNamedApps(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range auq.withNamedAppUsers {
		if err := auq.loadAppUsers(ctx, query, nodes,
			func(n *AdminUser) { n.appendNamedAppUsers(name) },
			func(n *AdminUser, e *AppUser) { n.appendNamedAppUsers(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range auq.loadTotal {
		if err := auq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (auq *AdminUserQuery) loadApps(ctx context.Context, query *AppQuery, nodes []*AdminUser, init func(*AdminUser), assign func(*AdminUser, *App)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*AdminUser)
	nids := make(map[string]map[*AdminUser]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(adminuser.AppsTable)
		s.Join(joinT).On(s.C(app.FieldID), joinT.C(adminuser.AppsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(adminuser.AppsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(adminuser.AppsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*AdminUser]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*App](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "apps" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (auq *AdminUserQuery) loadAppUsers(ctx context.Context, query *AppUserQuery, nodes []*AdminUser, init func(*AdminUser), assign func(*AdminUser, *AppUser)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*AdminUser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(appuser.FieldAdminUserID)
	}
	query.Where(predicate.AppUser(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(adminuser.AppUsersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AdminUserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "admin_user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (auq *AdminUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := auq.querySpec()
	if len(auq.modifiers) > 0 {
		_spec.Modifiers = auq.modifiers
	}
	_spec.Node.Columns = auq.ctx.Fields
	if len(auq.ctx.Fields) > 0 {
		_spec.Unique = auq.ctx.Unique != nil && *auq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, auq.driver, _spec)
}

func (auq *AdminUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(adminuser.Table, adminuser.Columns, sqlgraph.NewFieldSpec(adminuser.FieldID, field.TypeString))
	_spec.From = auq.sql
	if unique := auq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if auq.path != nil {
		_spec.Unique = true
	}
	if fields := auq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminuser.FieldID)
		for i := range fields {
			if fields[i] != adminuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := auq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := auq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := auq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := auq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (auq *AdminUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(auq.driver.Dialect())
	t1 := builder.Table(adminuser.Table)
	columns := auq.ctx.Fields
	if len(columns) == 0 {
		columns = adminuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if auq.sql != nil {
		selector = auq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if auq.ctx.Unique != nil && *auq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range auq.modifiers {
		m(selector)
	}
	for _, p := range auq.predicates {
		p(selector)
	}
	for _, p := range auq.order {
		p(selector)
	}
	if offset := auq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := auq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (auq *AdminUserQuery) Modify(modifiers ...func(s *sql.Selector)) *AdminUserSelect {
	auq.modifiers = append(auq.modifiers, modifiers...)
	return auq.Select()
}

// WithNamedApps tells the query-builder to eager-load the nodes that are connected to the "apps"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (auq *AdminUserQuery) WithNamedApps(name string, opts ...func(*AppQuery)) *AdminUserQuery {
	query := (&AppClient{config: auq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if auq.withNamedApps == nil {
		auq.withNamedApps = make(map[string]*AppQuery)
	}
	auq.withNamedApps[name] = query
	return auq
}

// WithNamedAppUsers tells the query-builder to eager-load the nodes that are connected to the "app_users"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (auq *AdminUserQuery) WithNamedAppUsers(name string, opts ...func(*AppUserQuery)) *AdminUserQuery {
	query := (&AppUserClient{config: auq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if auq.withNamedAppUsers == nil {
		auq.withNamedAppUsers = make(map[string]*AppUserQuery)
	}
	auq.withNamedAppUsers[name] = query
	return auq
}

// AdminUserGroupBy is the group-by builder for AdminUser entities.
type AdminUserGroupBy struct {
	selector
	build *AdminUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (augb *AdminUserGroupBy) Aggregate(fns ...AggregateFunc) *AdminUserGroupBy {
	augb.fns = append(augb.fns, fns...)
	return augb
}

// Scan applies the selector query and scans the result into the given value.
func (augb *AdminUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, augb.build.ctx, ent.OpQueryGroupBy)
	if err := augb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminUserQuery, *AdminUserGroupBy](ctx, augb.build, augb, augb.build.inters, v)
}

func (augb *AdminUserGroupBy) sqlScan(ctx context.Context, root *AdminUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(augb.fns))
	for _, fn := range augb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*augb.flds)+len(augb.fns))
		for _, f := range *augb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*augb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := augb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AdminUserSelect is the builder for selecting fields of AdminUser entities.
type AdminUserSelect struct {
	*AdminUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aus *AdminUserSelect) Aggregate(fns ...AggregateFunc) *AdminUserSelect {
	aus.fns = append(aus.fns, fns...)
	return aus
}

// Scan applies the selector query and scans the result into the given value.
func (aus *AdminUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aus.ctx, ent.OpQuerySelect)
	if err := aus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminUserQuery, *AdminUserSelect](ctx, aus.AdminUserQuery, aus, aus.inters, v)
}

func (aus *AdminUserSelect) sqlScan(ctx context.Context, root *AdminUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aus.fns))
	for _, fn := range aus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aus *AdminUserSelect) Modify(modifiers ...func(s *sql.Selector)) *AdminUserSelect {
	aus.modifiers = append(aus.modifiers, modifiers...)
	return aus
}
