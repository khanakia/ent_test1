// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"saas/gen/ent/posttype"
	"saas/gen/ent/posttypeform"
	"saas/gen/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostTypeFormQuery is the builder for querying PostTypeForm entities.
type PostTypeFormQuery struct {
	config
	ctx          *QueryContext
	order        []posttypeform.OrderOption
	inters       []Interceptor
	predicates   []predicate.PostTypeForm
	withPostType *PostTypeQuery
	loadTotal    []func(context.Context, []*PostTypeForm) error
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PostTypeFormQuery builder.
func (ptfq *PostTypeFormQuery) Where(ps ...predicate.PostTypeForm) *PostTypeFormQuery {
	ptfq.predicates = append(ptfq.predicates, ps...)
	return ptfq
}

// Limit the number of records to be returned by this query.
func (ptfq *PostTypeFormQuery) Limit(limit int) *PostTypeFormQuery {
	ptfq.ctx.Limit = &limit
	return ptfq
}

// Offset to start from.
func (ptfq *PostTypeFormQuery) Offset(offset int) *PostTypeFormQuery {
	ptfq.ctx.Offset = &offset
	return ptfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptfq *PostTypeFormQuery) Unique(unique bool) *PostTypeFormQuery {
	ptfq.ctx.Unique = &unique
	return ptfq
}

// Order specifies how the records should be ordered.
func (ptfq *PostTypeFormQuery) Order(o ...posttypeform.OrderOption) *PostTypeFormQuery {
	ptfq.order = append(ptfq.order, o...)
	return ptfq
}

// QueryPostType chains the current query on the "post_type" edge.
func (ptfq *PostTypeFormQuery) QueryPostType() *PostTypeQuery {
	query := (&PostTypeClient{config: ptfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(posttypeform.Table, posttypeform.FieldID, selector),
			sqlgraph.To(posttype.Table, posttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, posttypeform.PostTypeTable, posttypeform.PostTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PostTypeForm entity from the query.
// Returns a *NotFoundError when no PostTypeForm was found.
func (ptfq *PostTypeFormQuery) First(ctx context.Context) (*PostTypeForm, error) {
	nodes, err := ptfq.Limit(1).All(setContextOp(ctx, ptfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{posttypeform.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptfq *PostTypeFormQuery) FirstX(ctx context.Context) *PostTypeForm {
	node, err := ptfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PostTypeForm ID from the query.
// Returns a *NotFoundError when no PostTypeForm ID was found.
func (ptfq *PostTypeFormQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ptfq.Limit(1).IDs(setContextOp(ctx, ptfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{posttypeform.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptfq *PostTypeFormQuery) FirstIDX(ctx context.Context) string {
	id, err := ptfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PostTypeForm entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PostTypeForm entity is found.
// Returns a *NotFoundError when no PostTypeForm entities are found.
func (ptfq *PostTypeFormQuery) Only(ctx context.Context) (*PostTypeForm, error) {
	nodes, err := ptfq.Limit(2).All(setContextOp(ctx, ptfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{posttypeform.Label}
	default:
		return nil, &NotSingularError{posttypeform.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptfq *PostTypeFormQuery) OnlyX(ctx context.Context) *PostTypeForm {
	node, err := ptfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PostTypeForm ID in the query.
// Returns a *NotSingularError when more than one PostTypeForm ID is found.
// Returns a *NotFoundError when no entities are found.
func (ptfq *PostTypeFormQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ptfq.Limit(2).IDs(setContextOp(ctx, ptfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{posttypeform.Label}
	default:
		err = &NotSingularError{posttypeform.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptfq *PostTypeFormQuery) OnlyIDX(ctx context.Context) string {
	id, err := ptfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PostTypeForms.
func (ptfq *PostTypeFormQuery) All(ctx context.Context) ([]*PostTypeForm, error) {
	ctx = setContextOp(ctx, ptfq.ctx, "All")
	if err := ptfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PostTypeForm, *PostTypeFormQuery]()
	return withInterceptors[[]*PostTypeForm](ctx, ptfq, qr, ptfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ptfq *PostTypeFormQuery) AllX(ctx context.Context) []*PostTypeForm {
	nodes, err := ptfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PostTypeForm IDs.
func (ptfq *PostTypeFormQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ptfq.ctx.Unique == nil && ptfq.path != nil {
		ptfq.Unique(true)
	}
	ctx = setContextOp(ctx, ptfq.ctx, "IDs")
	if err = ptfq.Select(posttypeform.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptfq *PostTypeFormQuery) IDsX(ctx context.Context) []string {
	ids, err := ptfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptfq *PostTypeFormQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ptfq.ctx, "Count")
	if err := ptfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ptfq, querierCount[*PostTypeFormQuery](), ptfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ptfq *PostTypeFormQuery) CountX(ctx context.Context) int {
	count, err := ptfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptfq *PostTypeFormQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ptfq.ctx, "Exist")
	switch _, err := ptfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ptfq *PostTypeFormQuery) ExistX(ctx context.Context) bool {
	exist, err := ptfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PostTypeFormQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptfq *PostTypeFormQuery) Clone() *PostTypeFormQuery {
	if ptfq == nil {
		return nil
	}
	return &PostTypeFormQuery{
		config:       ptfq.config,
		ctx:          ptfq.ctx.Clone(),
		order:        append([]posttypeform.OrderOption{}, ptfq.order...),
		inters:       append([]Interceptor{}, ptfq.inters...),
		predicates:   append([]predicate.PostTypeForm{}, ptfq.predicates...),
		withPostType: ptfq.withPostType.Clone(),
		// clone intermediate query.
		sql:  ptfq.sql.Clone(),
		path: ptfq.path,
	}
}

// WithPostType tells the query-builder to eager-load the nodes that are connected to
// the "post_type" edge. The optional arguments are used to configure the query builder of the edge.
func (ptfq *PostTypeFormQuery) WithPostType(opts ...func(*PostTypeQuery)) *PostTypeFormQuery {
	query := (&PostTypeClient{config: ptfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptfq.withPostType = query
	return ptfq
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
//	client.PostTypeForm.Query().
//		GroupBy(posttypeform.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ptfq *PostTypeFormQuery) GroupBy(field string, fields ...string) *PostTypeFormGroupBy {
	ptfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PostTypeFormGroupBy{build: ptfq}
	grbuild.flds = &ptfq.ctx.Fields
	grbuild.label = posttypeform.Label
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
//	client.PostTypeForm.Query().
//		Select(posttypeform.FieldCreatedAt).
//		Scan(ctx, &v)
func (ptfq *PostTypeFormQuery) Select(fields ...string) *PostTypeFormSelect {
	ptfq.ctx.Fields = append(ptfq.ctx.Fields, fields...)
	sbuild := &PostTypeFormSelect{PostTypeFormQuery: ptfq}
	sbuild.label = posttypeform.Label
	sbuild.flds, sbuild.scan = &ptfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PostTypeFormSelect configured with the given aggregations.
func (ptfq *PostTypeFormQuery) Aggregate(fns ...AggregateFunc) *PostTypeFormSelect {
	return ptfq.Select().Aggregate(fns...)
}

func (ptfq *PostTypeFormQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ptfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ptfq); err != nil {
				return err
			}
		}
	}
	for _, f := range ptfq.ctx.Fields {
		if !posttypeform.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptfq.path != nil {
		prev, err := ptfq.path(ctx)
		if err != nil {
			return err
		}
		ptfq.sql = prev
	}
	return nil
}

func (ptfq *PostTypeFormQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PostTypeForm, error) {
	var (
		nodes       = []*PostTypeForm{}
		_spec       = ptfq.querySpec()
		loadedTypes = [1]bool{
			ptfq.withPostType != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PostTypeForm).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PostTypeForm{config: ptfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ptfq.modifiers) > 0 {
		_spec.Modifiers = ptfq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ptfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ptfq.withPostType; query != nil {
		if err := ptfq.loadPostType(ctx, query, nodes, nil,
			func(n *PostTypeForm, e *PostType) { n.Edges.PostType = e }); err != nil {
			return nil, err
		}
	}
	for i := range ptfq.loadTotal {
		if err := ptfq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ptfq *PostTypeFormQuery) loadPostType(ctx context.Context, query *PostTypeQuery, nodes []*PostTypeForm, init func(*PostTypeForm), assign func(*PostTypeForm, *PostType)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*PostTypeForm)
	for i := range nodes {
		fk := nodes[i].PostTypeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(posttype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "post_type_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ptfq *PostTypeFormQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptfq.querySpec()
	if len(ptfq.modifiers) > 0 {
		_spec.Modifiers = ptfq.modifiers
	}
	_spec.Node.Columns = ptfq.ctx.Fields
	if len(ptfq.ctx.Fields) > 0 {
		_spec.Unique = ptfq.ctx.Unique != nil && *ptfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ptfq.driver, _spec)
}

func (ptfq *PostTypeFormQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(posttypeform.Table, posttypeform.Columns, sqlgraph.NewFieldSpec(posttypeform.FieldID, field.TypeString))
	_spec.From = ptfq.sql
	if unique := ptfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ptfq.path != nil {
		_spec.Unique = true
	}
	if fields := ptfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, posttypeform.FieldID)
		for i := range fields {
			if fields[i] != posttypeform.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ptfq.withPostType != nil {
			_spec.Node.AddColumnOnce(posttypeform.FieldPostTypeID)
		}
	}
	if ps := ptfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptfq *PostTypeFormQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptfq.driver.Dialect())
	t1 := builder.Table(posttypeform.Table)
	columns := ptfq.ctx.Fields
	if len(columns) == 0 {
		columns = posttypeform.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptfq.sql != nil {
		selector = ptfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptfq.ctx.Unique != nil && *ptfq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ptfq.modifiers {
		m(selector)
	}
	for _, p := range ptfq.predicates {
		p(selector)
	}
	for _, p := range ptfq.order {
		p(selector)
	}
	if offset := ptfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ptfq *PostTypeFormQuery) Modify(modifiers ...func(s *sql.Selector)) *PostTypeFormSelect {
	ptfq.modifiers = append(ptfq.modifiers, modifiers...)
	return ptfq.Select()
}

// PostTypeFormGroupBy is the group-by builder for PostTypeForm entities.
type PostTypeFormGroupBy struct {
	selector
	build *PostTypeFormQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptfgb *PostTypeFormGroupBy) Aggregate(fns ...AggregateFunc) *PostTypeFormGroupBy {
	ptfgb.fns = append(ptfgb.fns, fns...)
	return ptfgb
}

// Scan applies the selector query and scans the result into the given value.
func (ptfgb *PostTypeFormGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptfgb.build.ctx, "GroupBy")
	if err := ptfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PostTypeFormQuery, *PostTypeFormGroupBy](ctx, ptfgb.build, ptfgb, ptfgb.build.inters, v)
}

func (ptfgb *PostTypeFormGroupBy) sqlScan(ctx context.Context, root *PostTypeFormQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ptfgb.fns))
	for _, fn := range ptfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ptfgb.flds)+len(ptfgb.fns))
		for _, f := range *ptfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ptfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PostTypeFormSelect is the builder for selecting fields of PostTypeForm entities.
type PostTypeFormSelect struct {
	*PostTypeFormQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ptfs *PostTypeFormSelect) Aggregate(fns ...AggregateFunc) *PostTypeFormSelect {
	ptfs.fns = append(ptfs.fns, fns...)
	return ptfs
}

// Scan applies the selector query and scans the result into the given value.
func (ptfs *PostTypeFormSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptfs.ctx, "Select")
	if err := ptfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PostTypeFormQuery, *PostTypeFormSelect](ctx, ptfs.PostTypeFormQuery, ptfs, ptfs.inters, v)
}

func (ptfs *PostTypeFormSelect) sqlScan(ctx context.Context, root *PostTypeFormQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ptfs.fns))
	for _, fn := range ptfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ptfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ptfs *PostTypeFormSelect) Modify(modifiers ...func(s *sql.Selector)) *PostTypeFormSelect {
	ptfs.modifiers = append(ptfs.modifiers, modifiers...)
	return ptfs
}