// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/post"
	"saas/gen/ent/poststatus"
	"saas/gen/ent/posttype"
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostStatusUpdate is the builder for updating PostStatus entities.
type PostStatusUpdate struct {
	config
	hooks     []Hook
	mutation  *PostStatusMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PostStatusUpdate builder.
func (psu *PostStatusUpdate) Where(ps ...predicate.PostStatus) *PostStatusUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// SetUpdatedAt sets the "updated_at" field.
func (psu *PostStatusUpdate) SetUpdatedAt(t time.Time) *PostStatusUpdate {
	psu.mutation.SetUpdatedAt(t)
	return psu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (psu *PostStatusUpdate) ClearUpdatedAt() *PostStatusUpdate {
	psu.mutation.ClearUpdatedAt()
	return psu
}

// SetName sets the "name" field.
func (psu *PostStatusUpdate) SetName(s string) *PostStatusUpdate {
	psu.mutation.SetName(s)
	return psu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (psu *PostStatusUpdate) SetNillableName(s *string) *PostStatusUpdate {
	if s != nil {
		psu.SetName(*s)
	}
	return psu
}

// ClearName clears the value of the "name" field.
func (psu *PostStatusUpdate) ClearName() *PostStatusUpdate {
	psu.mutation.ClearName()
	return psu
}

// SetSlug sets the "slug" field.
func (psu *PostStatusUpdate) SetSlug(s string) *PostStatusUpdate {
	psu.mutation.SetSlug(s)
	return psu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (psu *PostStatusUpdate) SetNillableSlug(s *string) *PostStatusUpdate {
	if s != nil {
		psu.SetSlug(*s)
	}
	return psu
}

// ClearSlug clears the value of the "slug" field.
func (psu *PostStatusUpdate) ClearSlug() *PostStatusUpdate {
	psu.mutation.ClearSlug()
	return psu
}

// SetStatus sets the "status" field.
func (psu *PostStatusUpdate) SetStatus(b bool) *PostStatusUpdate {
	psu.mutation.SetStatus(b)
	return psu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (psu *PostStatusUpdate) SetNillableStatus(b *bool) *PostStatusUpdate {
	if b != nil {
		psu.SetStatus(*b)
	}
	return psu
}

// ClearStatus clears the value of the "status" field.
func (psu *PostStatusUpdate) ClearStatus() *PostStatusUpdate {
	psu.mutation.ClearStatus()
	return psu
}

// SetPostTypeID sets the "post_type_id" field.
func (psu *PostStatusUpdate) SetPostTypeID(s string) *PostStatusUpdate {
	psu.mutation.SetPostTypeID(s)
	return psu
}

// SetNillablePostTypeID sets the "post_type_id" field if the given value is not nil.
func (psu *PostStatusUpdate) SetNillablePostTypeID(s *string) *PostStatusUpdate {
	if s != nil {
		psu.SetPostTypeID(*s)
	}
	return psu
}

// ClearPostTypeID clears the value of the "post_type_id" field.
func (psu *PostStatusUpdate) ClearPostTypeID() *PostStatusUpdate {
	psu.mutation.ClearPostTypeID()
	return psu
}

// SetPostType sets the "post_type" edge to the PostType entity.
func (psu *PostStatusUpdate) SetPostType(p *PostType) *PostStatusUpdate {
	return psu.SetPostTypeID(p.ID)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (psu *PostStatusUpdate) AddPostIDs(ids ...string) *PostStatusUpdate {
	psu.mutation.AddPostIDs(ids...)
	return psu
}

// AddPosts adds the "posts" edges to the Post entity.
func (psu *PostStatusUpdate) AddPosts(p ...*Post) *PostStatusUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psu.AddPostIDs(ids...)
}

// Mutation returns the PostStatusMutation object of the builder.
func (psu *PostStatusUpdate) Mutation() *PostStatusMutation {
	return psu.mutation
}

// ClearPostType clears the "post_type" edge to the PostType entity.
func (psu *PostStatusUpdate) ClearPostType() *PostStatusUpdate {
	psu.mutation.ClearPostType()
	return psu
}

// ClearPosts clears all "posts" edges to the Post entity.
func (psu *PostStatusUpdate) ClearPosts() *PostStatusUpdate {
	psu.mutation.ClearPosts()
	return psu
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (psu *PostStatusUpdate) RemovePostIDs(ids ...string) *PostStatusUpdate {
	psu.mutation.RemovePostIDs(ids...)
	return psu
}

// RemovePosts removes "posts" edges to Post entities.
func (psu *PostStatusUpdate) RemovePosts(p ...*Post) *PostStatusUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psu.RemovePostIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *PostStatusUpdate) Save(ctx context.Context) (int, error) {
	psu.defaults()
	return withHooks(ctx, psu.sqlSave, psu.mutation, psu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psu *PostStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *PostStatusUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *PostStatusUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psu *PostStatusUpdate) defaults() {
	if _, ok := psu.mutation.UpdatedAt(); !ok && !psu.mutation.UpdatedAtCleared() {
		v := poststatus.UpdateDefaultUpdatedAt()
		psu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (psu *PostStatusUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PostStatusUpdate {
	psu.modifiers = append(psu.modifiers, modifiers...)
	return psu
}

func (psu *PostStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(poststatus.Table, poststatus.Columns, sqlgraph.NewFieldSpec(poststatus.FieldID, field.TypeString))
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if psu.mutation.CreatedAtCleared() {
		_spec.ClearField(poststatus.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := psu.mutation.UpdatedAt(); ok {
		_spec.SetField(poststatus.FieldUpdatedAt, field.TypeTime, value)
	}
	if psu.mutation.UpdatedAtCleared() {
		_spec.ClearField(poststatus.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := psu.mutation.Name(); ok {
		_spec.SetField(poststatus.FieldName, field.TypeString, value)
	}
	if psu.mutation.NameCleared() {
		_spec.ClearField(poststatus.FieldName, field.TypeString)
	}
	if value, ok := psu.mutation.Slug(); ok {
		_spec.SetField(poststatus.FieldSlug, field.TypeString, value)
	}
	if psu.mutation.SlugCleared() {
		_spec.ClearField(poststatus.FieldSlug, field.TypeString)
	}
	if value, ok := psu.mutation.Status(); ok {
		_spec.SetField(poststatus.FieldStatus, field.TypeBool, value)
	}
	if psu.mutation.StatusCleared() {
		_spec.ClearField(poststatus.FieldStatus, field.TypeBool)
	}
	if psu.mutation.PostTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   poststatus.PostTypeTable,
			Columns: []string{poststatus.PostTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttype.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.PostTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   poststatus.PostTypeTable,
			Columns: []string{poststatus.PostTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttype.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poststatus.PostsTable,
			Columns: []string{poststatus.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !psu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poststatus.PostsTable,
			Columns: []string{poststatus.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poststatus.PostsTable,
			Columns: []string{poststatus.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(psu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{poststatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	psu.mutation.done = true
	return n, nil
}

// PostStatusUpdateOne is the builder for updating a single PostStatus entity.
type PostStatusUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PostStatusMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (psuo *PostStatusUpdateOne) SetUpdatedAt(t time.Time) *PostStatusUpdateOne {
	psuo.mutation.SetUpdatedAt(t)
	return psuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (psuo *PostStatusUpdateOne) ClearUpdatedAt() *PostStatusUpdateOne {
	psuo.mutation.ClearUpdatedAt()
	return psuo
}

// SetName sets the "name" field.
func (psuo *PostStatusUpdateOne) SetName(s string) *PostStatusUpdateOne {
	psuo.mutation.SetName(s)
	return psuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (psuo *PostStatusUpdateOne) SetNillableName(s *string) *PostStatusUpdateOne {
	if s != nil {
		psuo.SetName(*s)
	}
	return psuo
}

// ClearName clears the value of the "name" field.
func (psuo *PostStatusUpdateOne) ClearName() *PostStatusUpdateOne {
	psuo.mutation.ClearName()
	return psuo
}

// SetSlug sets the "slug" field.
func (psuo *PostStatusUpdateOne) SetSlug(s string) *PostStatusUpdateOne {
	psuo.mutation.SetSlug(s)
	return psuo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (psuo *PostStatusUpdateOne) SetNillableSlug(s *string) *PostStatusUpdateOne {
	if s != nil {
		psuo.SetSlug(*s)
	}
	return psuo
}

// ClearSlug clears the value of the "slug" field.
func (psuo *PostStatusUpdateOne) ClearSlug() *PostStatusUpdateOne {
	psuo.mutation.ClearSlug()
	return psuo
}

// SetStatus sets the "status" field.
func (psuo *PostStatusUpdateOne) SetStatus(b bool) *PostStatusUpdateOne {
	psuo.mutation.SetStatus(b)
	return psuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (psuo *PostStatusUpdateOne) SetNillableStatus(b *bool) *PostStatusUpdateOne {
	if b != nil {
		psuo.SetStatus(*b)
	}
	return psuo
}

// ClearStatus clears the value of the "status" field.
func (psuo *PostStatusUpdateOne) ClearStatus() *PostStatusUpdateOne {
	psuo.mutation.ClearStatus()
	return psuo
}

// SetPostTypeID sets the "post_type_id" field.
func (psuo *PostStatusUpdateOne) SetPostTypeID(s string) *PostStatusUpdateOne {
	psuo.mutation.SetPostTypeID(s)
	return psuo
}

// SetNillablePostTypeID sets the "post_type_id" field if the given value is not nil.
func (psuo *PostStatusUpdateOne) SetNillablePostTypeID(s *string) *PostStatusUpdateOne {
	if s != nil {
		psuo.SetPostTypeID(*s)
	}
	return psuo
}

// ClearPostTypeID clears the value of the "post_type_id" field.
func (psuo *PostStatusUpdateOne) ClearPostTypeID() *PostStatusUpdateOne {
	psuo.mutation.ClearPostTypeID()
	return psuo
}

// SetPostType sets the "post_type" edge to the PostType entity.
func (psuo *PostStatusUpdateOne) SetPostType(p *PostType) *PostStatusUpdateOne {
	return psuo.SetPostTypeID(p.ID)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (psuo *PostStatusUpdateOne) AddPostIDs(ids ...string) *PostStatusUpdateOne {
	psuo.mutation.AddPostIDs(ids...)
	return psuo
}

// AddPosts adds the "posts" edges to the Post entity.
func (psuo *PostStatusUpdateOne) AddPosts(p ...*Post) *PostStatusUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psuo.AddPostIDs(ids...)
}

// Mutation returns the PostStatusMutation object of the builder.
func (psuo *PostStatusUpdateOne) Mutation() *PostStatusMutation {
	return psuo.mutation
}

// ClearPostType clears the "post_type" edge to the PostType entity.
func (psuo *PostStatusUpdateOne) ClearPostType() *PostStatusUpdateOne {
	psuo.mutation.ClearPostType()
	return psuo
}

// ClearPosts clears all "posts" edges to the Post entity.
func (psuo *PostStatusUpdateOne) ClearPosts() *PostStatusUpdateOne {
	psuo.mutation.ClearPosts()
	return psuo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (psuo *PostStatusUpdateOne) RemovePostIDs(ids ...string) *PostStatusUpdateOne {
	psuo.mutation.RemovePostIDs(ids...)
	return psuo
}

// RemovePosts removes "posts" edges to Post entities.
func (psuo *PostStatusUpdateOne) RemovePosts(p ...*Post) *PostStatusUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psuo.RemovePostIDs(ids...)
}

// Where appends a list predicates to the PostStatusUpdate builder.
func (psuo *PostStatusUpdateOne) Where(ps ...predicate.PostStatus) *PostStatusUpdateOne {
	psuo.mutation.Where(ps...)
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *PostStatusUpdateOne) Select(field string, fields ...string) *PostStatusUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated PostStatus entity.
func (psuo *PostStatusUpdateOne) Save(ctx context.Context) (*PostStatus, error) {
	psuo.defaults()
	return withHooks(ctx, psuo.sqlSave, psuo.mutation, psuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *PostStatusUpdateOne) SaveX(ctx context.Context) *PostStatus {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *PostStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *PostStatusUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psuo *PostStatusUpdateOne) defaults() {
	if _, ok := psuo.mutation.UpdatedAt(); !ok && !psuo.mutation.UpdatedAtCleared() {
		v := poststatus.UpdateDefaultUpdatedAt()
		psuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (psuo *PostStatusUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PostStatusUpdateOne {
	psuo.modifiers = append(psuo.modifiers, modifiers...)
	return psuo
}

func (psuo *PostStatusUpdateOne) sqlSave(ctx context.Context) (_node *PostStatus, err error) {
	_spec := sqlgraph.NewUpdateSpec(poststatus.Table, poststatus.Columns, sqlgraph.NewFieldSpec(poststatus.FieldID, field.TypeString))
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PostStatus.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, poststatus.FieldID)
		for _, f := range fields {
			if !poststatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != poststatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if psuo.mutation.CreatedAtCleared() {
		_spec.ClearField(poststatus.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := psuo.mutation.UpdatedAt(); ok {
		_spec.SetField(poststatus.FieldUpdatedAt, field.TypeTime, value)
	}
	if psuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(poststatus.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := psuo.mutation.Name(); ok {
		_spec.SetField(poststatus.FieldName, field.TypeString, value)
	}
	if psuo.mutation.NameCleared() {
		_spec.ClearField(poststatus.FieldName, field.TypeString)
	}
	if value, ok := psuo.mutation.Slug(); ok {
		_spec.SetField(poststatus.FieldSlug, field.TypeString, value)
	}
	if psuo.mutation.SlugCleared() {
		_spec.ClearField(poststatus.FieldSlug, field.TypeString)
	}
	if value, ok := psuo.mutation.Status(); ok {
		_spec.SetField(poststatus.FieldStatus, field.TypeBool, value)
	}
	if psuo.mutation.StatusCleared() {
		_spec.ClearField(poststatus.FieldStatus, field.TypeBool)
	}
	if psuo.mutation.PostTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   poststatus.PostTypeTable,
			Columns: []string{poststatus.PostTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttype.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.PostTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   poststatus.PostTypeTable,
			Columns: []string{poststatus.PostTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttype.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poststatus.PostsTable,
			Columns: []string{poststatus.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !psuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poststatus.PostsTable,
			Columns: []string{poststatus.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poststatus.PostsTable,
			Columns: []string{poststatus.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(psuo.modifiers...)
	_node = &PostStatus{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{poststatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	psuo.mutation.done = true
	return _node, nil
}