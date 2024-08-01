// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"saas/gen/ent/predicate"
	"saas/gen/ent/workspace"
	"saas/gen/ent/workspaceinvite"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceInviteQuery is the builder for querying WorkspaceInvite entities.
type WorkspaceInviteQuery struct {
	config
	ctx           *QueryContext
	order         []workspaceinvite.OrderOption
	inters        []Interceptor
	predicates    []predicate.WorkspaceInvite
	withWorkspace *WorkspaceQuery
	loadTotal     []func(context.Context, []*WorkspaceInvite) error
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkspaceInviteQuery builder.
func (wiq *WorkspaceInviteQuery) Where(ps ...predicate.WorkspaceInvite) *WorkspaceInviteQuery {
	wiq.predicates = append(wiq.predicates, ps...)
	return wiq
}

// Limit the number of records to be returned by this query.
func (wiq *WorkspaceInviteQuery) Limit(limit int) *WorkspaceInviteQuery {
	wiq.ctx.Limit = &limit
	return wiq
}

// Offset to start from.
func (wiq *WorkspaceInviteQuery) Offset(offset int) *WorkspaceInviteQuery {
	wiq.ctx.Offset = &offset
	return wiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wiq *WorkspaceInviteQuery) Unique(unique bool) *WorkspaceInviteQuery {
	wiq.ctx.Unique = &unique
	return wiq
}

// Order specifies how the records should be ordered.
func (wiq *WorkspaceInviteQuery) Order(o ...workspaceinvite.OrderOption) *WorkspaceInviteQuery {
	wiq.order = append(wiq.order, o...)
	return wiq
}

// QueryWorkspace chains the current query on the "workspace" edge.
func (wiq *WorkspaceInviteQuery) QueryWorkspace() *WorkspaceQuery {
	query := (&WorkspaceClient{config: wiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workspaceinvite.Table, workspaceinvite.FieldID, selector),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workspaceinvite.WorkspaceTable, workspaceinvite.WorkspaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(wiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkspaceInvite entity from the query.
// Returns a *NotFoundError when no WorkspaceInvite was found.
func (wiq *WorkspaceInviteQuery) First(ctx context.Context) (*WorkspaceInvite, error) {
	nodes, err := wiq.Limit(1).All(setContextOp(ctx, wiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workspaceinvite.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wiq *WorkspaceInviteQuery) FirstX(ctx context.Context) *WorkspaceInvite {
	node, err := wiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkspaceInvite ID from the query.
// Returns a *NotFoundError when no WorkspaceInvite ID was found.
func (wiq *WorkspaceInviteQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = wiq.Limit(1).IDs(setContextOp(ctx, wiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workspaceinvite.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wiq *WorkspaceInviteQuery) FirstIDX(ctx context.Context) string {
	id, err := wiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkspaceInvite entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WorkspaceInvite entity is found.
// Returns a *NotFoundError when no WorkspaceInvite entities are found.
func (wiq *WorkspaceInviteQuery) Only(ctx context.Context) (*WorkspaceInvite, error) {
	nodes, err := wiq.Limit(2).All(setContextOp(ctx, wiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workspaceinvite.Label}
	default:
		return nil, &NotSingularError{workspaceinvite.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wiq *WorkspaceInviteQuery) OnlyX(ctx context.Context) *WorkspaceInvite {
	node, err := wiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkspaceInvite ID in the query.
// Returns a *NotSingularError when more than one WorkspaceInvite ID is found.
// Returns a *NotFoundError when no entities are found.
func (wiq *WorkspaceInviteQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = wiq.Limit(2).IDs(setContextOp(ctx, wiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workspaceinvite.Label}
	default:
		err = &NotSingularError{workspaceinvite.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wiq *WorkspaceInviteQuery) OnlyIDX(ctx context.Context) string {
	id, err := wiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkspaceInvites.
func (wiq *WorkspaceInviteQuery) All(ctx context.Context) ([]*WorkspaceInvite, error) {
	ctx = setContextOp(ctx, wiq.ctx, "All")
	if err := wiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WorkspaceInvite, *WorkspaceInviteQuery]()
	return withInterceptors[[]*WorkspaceInvite](ctx, wiq, qr, wiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wiq *WorkspaceInviteQuery) AllX(ctx context.Context) []*WorkspaceInvite {
	nodes, err := wiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkspaceInvite IDs.
func (wiq *WorkspaceInviteQuery) IDs(ctx context.Context) (ids []string, err error) {
	if wiq.ctx.Unique == nil && wiq.path != nil {
		wiq.Unique(true)
	}
	ctx = setContextOp(ctx, wiq.ctx, "IDs")
	if err = wiq.Select(workspaceinvite.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wiq *WorkspaceInviteQuery) IDsX(ctx context.Context) []string {
	ids, err := wiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wiq *WorkspaceInviteQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wiq.ctx, "Count")
	if err := wiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wiq, querierCount[*WorkspaceInviteQuery](), wiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wiq *WorkspaceInviteQuery) CountX(ctx context.Context) int {
	count, err := wiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wiq *WorkspaceInviteQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wiq.ctx, "Exist")
	switch _, err := wiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wiq *WorkspaceInviteQuery) ExistX(ctx context.Context) bool {
	exist, err := wiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkspaceInviteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wiq *WorkspaceInviteQuery) Clone() *WorkspaceInviteQuery {
	if wiq == nil {
		return nil
	}
	return &WorkspaceInviteQuery{
		config:        wiq.config,
		ctx:           wiq.ctx.Clone(),
		order:         append([]workspaceinvite.OrderOption{}, wiq.order...),
		inters:        append([]Interceptor{}, wiq.inters...),
		predicates:    append([]predicate.WorkspaceInvite{}, wiq.predicates...),
		withWorkspace: wiq.withWorkspace.Clone(),
		// clone intermediate query.
		sql:  wiq.sql.Clone(),
		path: wiq.path,
	}
}

// WithWorkspace tells the query-builder to eager-load the nodes that are connected to
// the "workspace" edge. The optional arguments are used to configure the query builder of the edge.
func (wiq *WorkspaceInviteQuery) WithWorkspace(opts ...func(*WorkspaceQuery)) *WorkspaceInviteQuery {
	query := (&WorkspaceClient{config: wiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wiq.withWorkspace = query
	return wiq
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
//	client.WorkspaceInvite.Query().
//		GroupBy(workspaceinvite.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wiq *WorkspaceInviteQuery) GroupBy(field string, fields ...string) *WorkspaceInviteGroupBy {
	wiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkspaceInviteGroupBy{build: wiq}
	grbuild.flds = &wiq.ctx.Fields
	grbuild.label = workspaceinvite.Label
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
//	client.WorkspaceInvite.Query().
//		Select(workspaceinvite.FieldCreatedAt).
//		Scan(ctx, &v)
func (wiq *WorkspaceInviteQuery) Select(fields ...string) *WorkspaceInviteSelect {
	wiq.ctx.Fields = append(wiq.ctx.Fields, fields...)
	sbuild := &WorkspaceInviteSelect{WorkspaceInviteQuery: wiq}
	sbuild.label = workspaceinvite.Label
	sbuild.flds, sbuild.scan = &wiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkspaceInviteSelect configured with the given aggregations.
func (wiq *WorkspaceInviteQuery) Aggregate(fns ...AggregateFunc) *WorkspaceInviteSelect {
	return wiq.Select().Aggregate(fns...)
}

func (wiq *WorkspaceInviteQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wiq); err != nil {
				return err
			}
		}
	}
	for _, f := range wiq.ctx.Fields {
		if !workspaceinvite.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wiq.path != nil {
		prev, err := wiq.path(ctx)
		if err != nil {
			return err
		}
		wiq.sql = prev
	}
	return nil
}

func (wiq *WorkspaceInviteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WorkspaceInvite, error) {
	var (
		nodes       = []*WorkspaceInvite{}
		_spec       = wiq.querySpec()
		loadedTypes = [1]bool{
			wiq.withWorkspace != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WorkspaceInvite).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WorkspaceInvite{config: wiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wiq.modifiers) > 0 {
		_spec.Modifiers = wiq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wiq.withWorkspace; query != nil {
		if err := wiq.loadWorkspace(ctx, query, nodes, nil,
			func(n *WorkspaceInvite, e *Workspace) { n.Edges.Workspace = e }); err != nil {
			return nil, err
		}
	}
	for i := range wiq.loadTotal {
		if err := wiq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wiq *WorkspaceInviteQuery) loadWorkspace(ctx context.Context, query *WorkspaceQuery, nodes []*WorkspaceInvite, init func(*WorkspaceInvite), assign func(*WorkspaceInvite, *Workspace)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*WorkspaceInvite)
	for i := range nodes {
		fk := nodes[i].WorkspaceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workspace.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "workspace_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (wiq *WorkspaceInviteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wiq.querySpec()
	if len(wiq.modifiers) > 0 {
		_spec.Modifiers = wiq.modifiers
	}
	_spec.Node.Columns = wiq.ctx.Fields
	if len(wiq.ctx.Fields) > 0 {
		_spec.Unique = wiq.ctx.Unique != nil && *wiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wiq.driver, _spec)
}

func (wiq *WorkspaceInviteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workspaceinvite.Table, workspaceinvite.Columns, sqlgraph.NewFieldSpec(workspaceinvite.FieldID, field.TypeString))
	_spec.From = wiq.sql
	if unique := wiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wiq.path != nil {
		_spec.Unique = true
	}
	if fields := wiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workspaceinvite.FieldID)
		for i := range fields {
			if fields[i] != workspaceinvite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if wiq.withWorkspace != nil {
			_spec.Node.AddColumnOnce(workspaceinvite.FieldWorkspaceID)
		}
	}
	if ps := wiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wiq *WorkspaceInviteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wiq.driver.Dialect())
	t1 := builder.Table(workspaceinvite.Table)
	columns := wiq.ctx.Fields
	if len(columns) == 0 {
		columns = workspaceinvite.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wiq.sql != nil {
		selector = wiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wiq.ctx.Unique != nil && *wiq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range wiq.modifiers {
		m(selector)
	}
	for _, p := range wiq.predicates {
		p(selector)
	}
	for _, p := range wiq.order {
		p(selector)
	}
	if offset := wiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wiq *WorkspaceInviteQuery) Modify(modifiers ...func(s *sql.Selector)) *WorkspaceInviteSelect {
	wiq.modifiers = append(wiq.modifiers, modifiers...)
	return wiq.Select()
}

// WorkspaceInviteGroupBy is the group-by builder for WorkspaceInvite entities.
type WorkspaceInviteGroupBy struct {
	selector
	build *WorkspaceInviteQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wigb *WorkspaceInviteGroupBy) Aggregate(fns ...AggregateFunc) *WorkspaceInviteGroupBy {
	wigb.fns = append(wigb.fns, fns...)
	return wigb
}

// Scan applies the selector query and scans the result into the given value.
func (wigb *WorkspaceInviteGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wigb.build.ctx, "GroupBy")
	if err := wigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkspaceInviteQuery, *WorkspaceInviteGroupBy](ctx, wigb.build, wigb, wigb.build.inters, v)
}

func (wigb *WorkspaceInviteGroupBy) sqlScan(ctx context.Context, root *WorkspaceInviteQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wigb.fns))
	for _, fn := range wigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wigb.flds)+len(wigb.fns))
		for _, f := range *wigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkspaceInviteSelect is the builder for selecting fields of WorkspaceInvite entities.
type WorkspaceInviteSelect struct {
	*WorkspaceInviteQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wis *WorkspaceInviteSelect) Aggregate(fns ...AggregateFunc) *WorkspaceInviteSelect {
	wis.fns = append(wis.fns, fns...)
	return wis
}

// Scan applies the selector query and scans the result into the given value.
func (wis *WorkspaceInviteSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wis.ctx, "Select")
	if err := wis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkspaceInviteQuery, *WorkspaceInviteSelect](ctx, wis.WorkspaceInviteQuery, wis, wis.inters, v)
}

func (wis *WorkspaceInviteSelect) sqlScan(ctx context.Context, root *WorkspaceInviteQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wis.fns))
	for _, fn := range wis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wis *WorkspaceInviteSelect) Modify(modifiers ...func(s *sql.Selector)) *WorkspaceInviteSelect {
	wis.modifiers = append(wis.modifiers, modifiers...)
	return wis
}
