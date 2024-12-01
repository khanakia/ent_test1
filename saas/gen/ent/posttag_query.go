// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"saas/gen/ent/post"
	"saas/gen/ent/posttag"
	"saas/gen/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostTagQuery is the builder for querying PostTag entities.
type PostTagQuery struct {
	config
	ctx            *QueryContext
	order          []posttag.OrderOption
	inters         []Interceptor
	predicates     []predicate.PostTag
	withPosts      *PostQuery
	loadTotal      []func(context.Context, []*PostTag) error
	modifiers      []func(*sql.Selector)
	withNamedPosts map[string]*PostQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PostTagQuery builder.
func (ptq *PostTagQuery) Where(ps ...predicate.PostTag) *PostTagQuery {
	ptq.predicates = append(ptq.predicates, ps...)
	return ptq
}

// Limit the number of records to be returned by this query.
func (ptq *PostTagQuery) Limit(limit int) *PostTagQuery {
	ptq.ctx.Limit = &limit
	return ptq
}

// Offset to start from.
func (ptq *PostTagQuery) Offset(offset int) *PostTagQuery {
	ptq.ctx.Offset = &offset
	return ptq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptq *PostTagQuery) Unique(unique bool) *PostTagQuery {
	ptq.ctx.Unique = &unique
	return ptq
}

// Order specifies how the records should be ordered.
func (ptq *PostTagQuery) Order(o ...posttag.OrderOption) *PostTagQuery {
	ptq.order = append(ptq.order, o...)
	return ptq
}

// QueryPosts chains the current query on the "posts" edge.
func (ptq *PostTagQuery) QueryPosts() *PostQuery {
	query := (&PostClient{config: ptq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(posttag.Table, posttag.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, posttag.PostsTable, posttag.PostsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PostTag entity from the query.
// Returns a *NotFoundError when no PostTag was found.
func (ptq *PostTagQuery) First(ctx context.Context) (*PostTag, error) {
	nodes, err := ptq.Limit(1).All(setContextOp(ctx, ptq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{posttag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptq *PostTagQuery) FirstX(ctx context.Context) *PostTag {
	node, err := ptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PostTag ID from the query.
// Returns a *NotFoundError when no PostTag ID was found.
func (ptq *PostTagQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ptq.Limit(1).IDs(setContextOp(ctx, ptq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{posttag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptq *PostTagQuery) FirstIDX(ctx context.Context) string {
	id, err := ptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PostTag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PostTag entity is found.
// Returns a *NotFoundError when no PostTag entities are found.
func (ptq *PostTagQuery) Only(ctx context.Context) (*PostTag, error) {
	nodes, err := ptq.Limit(2).All(setContextOp(ctx, ptq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{posttag.Label}
	default:
		return nil, &NotSingularError{posttag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptq *PostTagQuery) OnlyX(ctx context.Context) *PostTag {
	node, err := ptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PostTag ID in the query.
// Returns a *NotSingularError when more than one PostTag ID is found.
// Returns a *NotFoundError when no entities are found.
func (ptq *PostTagQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ptq.Limit(2).IDs(setContextOp(ctx, ptq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{posttag.Label}
	default:
		err = &NotSingularError{posttag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptq *PostTagQuery) OnlyIDX(ctx context.Context) string {
	id, err := ptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PostTags.
func (ptq *PostTagQuery) All(ctx context.Context) ([]*PostTag, error) {
	ctx = setContextOp(ctx, ptq.ctx, "All")
	if err := ptq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PostTag, *PostTagQuery]()
	return withInterceptors[[]*PostTag](ctx, ptq, qr, ptq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ptq *PostTagQuery) AllX(ctx context.Context) []*PostTag {
	nodes, err := ptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PostTag IDs.
func (ptq *PostTagQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ptq.ctx.Unique == nil && ptq.path != nil {
		ptq.Unique(true)
	}
	ctx = setContextOp(ctx, ptq.ctx, "IDs")
	if err = ptq.Select(posttag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptq *PostTagQuery) IDsX(ctx context.Context) []string {
	ids, err := ptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptq *PostTagQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ptq.ctx, "Count")
	if err := ptq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ptq, querierCount[*PostTagQuery](), ptq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ptq *PostTagQuery) CountX(ctx context.Context) int {
	count, err := ptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptq *PostTagQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ptq.ctx, "Exist")
	switch _, err := ptq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ptq *PostTagQuery) ExistX(ctx context.Context) bool {
	exist, err := ptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PostTagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptq *PostTagQuery) Clone() *PostTagQuery {
	if ptq == nil {
		return nil
	}
	return &PostTagQuery{
		config:     ptq.config,
		ctx:        ptq.ctx.Clone(),
		order:      append([]posttag.OrderOption{}, ptq.order...),
		inters:     append([]Interceptor{}, ptq.inters...),
		predicates: append([]predicate.PostTag{}, ptq.predicates...),
		withPosts:  ptq.withPosts.Clone(),
		// clone intermediate query.
		sql:  ptq.sql.Clone(),
		path: ptq.path,
	}
}

// WithPosts tells the query-builder to eager-load the nodes that are connected to
// the "posts" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *PostTagQuery) WithPosts(opts ...func(*PostQuery)) *PostTagQuery {
	query := (&PostClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptq.withPosts = query
	return ptq
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
//	client.PostTag.Query().
//		GroupBy(posttag.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ptq *PostTagQuery) GroupBy(field string, fields ...string) *PostTagGroupBy {
	ptq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PostTagGroupBy{build: ptq}
	grbuild.flds = &ptq.ctx.Fields
	grbuild.label = posttag.Label
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
//	client.PostTag.Query().
//		Select(posttag.FieldCreatedAt).
//		Scan(ctx, &v)
func (ptq *PostTagQuery) Select(fields ...string) *PostTagSelect {
	ptq.ctx.Fields = append(ptq.ctx.Fields, fields...)
	sbuild := &PostTagSelect{PostTagQuery: ptq}
	sbuild.label = posttag.Label
	sbuild.flds, sbuild.scan = &ptq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PostTagSelect configured with the given aggregations.
func (ptq *PostTagQuery) Aggregate(fns ...AggregateFunc) *PostTagSelect {
	return ptq.Select().Aggregate(fns...)
}

func (ptq *PostTagQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ptq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ptq); err != nil {
				return err
			}
		}
	}
	for _, f := range ptq.ctx.Fields {
		if !posttag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptq.path != nil {
		prev, err := ptq.path(ctx)
		if err != nil {
			return err
		}
		ptq.sql = prev
	}
	return nil
}

func (ptq *PostTagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PostTag, error) {
	var (
		nodes       = []*PostTag{}
		_spec       = ptq.querySpec()
		loadedTypes = [1]bool{
			ptq.withPosts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PostTag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PostTag{config: ptq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ptq.modifiers) > 0 {
		_spec.Modifiers = ptq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ptq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ptq.withPosts; query != nil {
		if err := ptq.loadPosts(ctx, query, nodes,
			func(n *PostTag) { n.Edges.Posts = []*Post{} },
			func(n *PostTag, e *Post) { n.Edges.Posts = append(n.Edges.Posts, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ptq.withNamedPosts {
		if err := ptq.loadPosts(ctx, query, nodes,
			func(n *PostTag) { n.appendNamedPosts(name) },
			func(n *PostTag, e *Post) { n.appendNamedPosts(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range ptq.loadTotal {
		if err := ptq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ptq *PostTagQuery) loadPosts(ctx context.Context, query *PostQuery, nodes []*PostTag, init func(*PostTag), assign func(*PostTag, *Post)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*PostTag)
	nids := make(map[string]map[*PostTag]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(posttag.PostsTable)
		s.Join(joinT).On(s.C(post.FieldID), joinT.C(posttag.PostsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(posttag.PostsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(posttag.PostsPrimaryKey[0]))
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
					nids[inValue] = map[*PostTag]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Post](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "posts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (ptq *PostTagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptq.querySpec()
	if len(ptq.modifiers) > 0 {
		_spec.Modifiers = ptq.modifiers
	}
	_spec.Node.Columns = ptq.ctx.Fields
	if len(ptq.ctx.Fields) > 0 {
		_spec.Unique = ptq.ctx.Unique != nil && *ptq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ptq.driver, _spec)
}

func (ptq *PostTagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(posttag.Table, posttag.Columns, sqlgraph.NewFieldSpec(posttag.FieldID, field.TypeString))
	_spec.From = ptq.sql
	if unique := ptq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ptq.path != nil {
		_spec.Unique = true
	}
	if fields := ptq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, posttag.FieldID)
		for i := range fields {
			if fields[i] != posttag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptq *PostTagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptq.driver.Dialect())
	t1 := builder.Table(posttag.Table)
	columns := ptq.ctx.Fields
	if len(columns) == 0 {
		columns = posttag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptq.sql != nil {
		selector = ptq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptq.ctx.Unique != nil && *ptq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ptq.modifiers {
		m(selector)
	}
	for _, p := range ptq.predicates {
		p(selector)
	}
	for _, p := range ptq.order {
		p(selector)
	}
	if offset := ptq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ptq *PostTagQuery) Modify(modifiers ...func(s *sql.Selector)) *PostTagSelect {
	ptq.modifiers = append(ptq.modifiers, modifiers...)
	return ptq.Select()
}

// WithNamedPosts tells the query-builder to eager-load the nodes that are connected to the "posts"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ptq *PostTagQuery) WithNamedPosts(name string, opts ...func(*PostQuery)) *PostTagQuery {
	query := (&PostClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if ptq.withNamedPosts == nil {
		ptq.withNamedPosts = make(map[string]*PostQuery)
	}
	ptq.withNamedPosts[name] = query
	return ptq
}

// PostTagGroupBy is the group-by builder for PostTag entities.
type PostTagGroupBy struct {
	selector
	build *PostTagQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptgb *PostTagGroupBy) Aggregate(fns ...AggregateFunc) *PostTagGroupBy {
	ptgb.fns = append(ptgb.fns, fns...)
	return ptgb
}

// Scan applies the selector query and scans the result into the given value.
func (ptgb *PostTagGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptgb.build.ctx, "GroupBy")
	if err := ptgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PostTagQuery, *PostTagGroupBy](ctx, ptgb.build, ptgb, ptgb.build.inters, v)
}

func (ptgb *PostTagGroupBy) sqlScan(ctx context.Context, root *PostTagQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ptgb.fns))
	for _, fn := range ptgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ptgb.flds)+len(ptgb.fns))
		for _, f := range *ptgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ptgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PostTagSelect is the builder for selecting fields of PostTag entities.
type PostTagSelect struct {
	*PostTagQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pts *PostTagSelect) Aggregate(fns ...AggregateFunc) *PostTagSelect {
	pts.fns = append(pts.fns, fns...)
	return pts
}

// Scan applies the selector query and scans the result into the given value.
func (pts *PostTagSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pts.ctx, "Select")
	if err := pts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PostTagQuery, *PostTagSelect](ctx, pts.PostTagQuery, pts, pts.inters, v)
}

func (pts *PostTagSelect) sqlScan(ctx context.Context, root *PostTagQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pts.fns))
	for _, fn := range pts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pts *PostTagSelect) Modify(modifiers ...func(s *sql.Selector)) *PostTagSelect {
	pts.modifiers = append(pts.modifiers, modifiers...)
	return pts
}
